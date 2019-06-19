package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
	"github.com/owulveryck/onnx-go/internal/x/images"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"gorgonia.org/tensor"
	"gorgonia.org/tensor/native"
)

// The 416x416 image is divided into a 13x13 grid. Each of these grid cells
// will predict 5 bounding boxes (boxesPerCell). A bounding box consists of
// five data items: x, y, width, height, and a confidence score. Each grid
// cell also predicts which class each bounding box belongs to.
//
const (
	hSize, wSize     = 416, 416
	blockSize        = 32
	gridHeight       = 13
	gridWidth        = 13
	boxesPerCell     = 5
	numClasses       = 20
	threshold        = 0.30
	drawingThreshold = 0.3
)

var (
	model   = flag.String("model", "model.onnx", "path to the model file")
	imgF    = flag.String("img", "", "path of an input tensor for testing")
	inputT  = flag.String("input", "", "tensor")
	img     image.Image
	classes = []string{"aeroplane", "bicycle", "bird", "boat", "bottle",
		"bus", "car", "cat", "chair", "cow",
		"diningtable", "dog", "horse", "motorbike", "person",
		"pottedplant", "sheep", "sofa", "train", "tv/monitor"}
	anchors = []float64{1.08, 1.19, 3.42, 4.41, 6.63, 11.38, 9.42, 5.11, 16.62, 10.52}
)

func main() {
	h := flag.Bool("h", false, "help")
	flag.Parse()
	if *h {
		flag.Usage()
		os.Exit(0)
	}
	if _, err := os.Stat(*model); err != nil && os.IsNotExist(err) {
		log.Fatalf("%v does not exist", *model)
	}
	// Create a backend receiver
	backend := gorgonnx.NewGraph()
	// Create a model and set the execution backend
	m := onnx.NewModel(backend)

	// read the onnx model
	b, err := ioutil.ReadFile(*model)
	if err != nil {
		log.Fatal(err)
	}
	// Decode it into the model
	must(m.UnmarshalBinary(b))

	m.SetInput(0, getInput())
	must(backend.Run())
	processOutput(m.GetOutputTensors())

}

func getInput() tensor.Tensor {
	if *inputT != "" {
		b, err := ioutil.ReadFile(*inputT)
		if err != nil {
			log.Fatal(err)
		}
		t, err := onnx.NewTensor(b)
		if err != nil {
			log.Fatal(err)
		}
		img, err = images.TensorToImg(t)
		if err != nil {
			log.Fatal(err)
		}
		return t
	}
	if *imgF != "" {
		f, err := os.Open(*imgF)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		img, err = jpeg.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		inputT := tensor.New(tensor.WithShape(1, 3, hSize, wSize), tensor.Of(tensor.Float32))
		err = images.ImageToBCHW(img, inputT)
		if err != nil {
			log.Fatal(err)
		}
		return inputT
	}
	log.Fatal("Please speficy an input")
	return nil
}

func processOutput(t []tensor.Tensor, err error) {
	if err != nil {
		log.Fatal(err)
	}
	dense := t[0].(*tensor.Dense)
	must(dense.Reshape(125, gridHeight, gridWidth))
	data, err := native.Tensor3F32(dense)
	if err != nil {
		log.Fatal(err)
	}

	var boxes = make([]box, gridHeight*gridWidth*boxesPerCell)
	var counter int
	// https://github.com/pjreddie/darknet/blob/61c9d02ec461e30d55762ec7669d6a1d3c356fb2/src/yolo_layer.c#L159
	for cx := 0; cx < gridWidth; cx++ {
		for cy := 0; cy < gridHeight; cy++ {
			for b := 0; b < boxesPerCell; b++ {
				channel := b * (numClasses + 5)
				tx := data[channel][cx][cy]
				ty := data[channel+1][cx][cy]
				tw := data[channel+2][cx][cy]
				th := data[channel+3][cx][cy]
				tc := data[channel+4][cx][cy]
				tclasses := make([]float32, 20)
				for i := 0; i < 20; i++ {
					tclasses[i] = data[channel+5+i][cx][cy]
				}
				// The predicted tx and ty coordinates are relative to the location
				// of the grid cell; we use the logistic sigmoid to constrain these
				// coordinates to the range 0 - 1. Then we add the cell coordinates
				// (0-12) and multiply by the number of pixels per grid cell (32).
				// Now x and y represent center of the bounding box in the original
				// 416x416 image space.
				// https://github.com/hollance/Forge/blob/04109c856237faec87deecb55126d4a20fa4f59b/Examples/YOLO/YOLO/YOLO.swift#L154
				x := int((float32(cx) + sigmoid(tx)) * blockSize)
				y := int((float32(cy) + sigmoid(ty)) * blockSize)
				// The size of the bounding box, tw and th, is predicted relative to
				// the size of an "anchor" box. Here we also transform the width and
				// height into the original 416x416 image space.
				w := int(exp(tw) * anchors[2*b] * blockSize)
				h := int(exp(th) * anchors[2*b+1] * blockSize)

				boxes[counter] = box{
					gridcell:   []int{cx, cy},
					r:          image.Rect(max(y-w/2, 0), max(x-h/2, 0), min(y+w/2, wSize), min(x+h/2, hSize)),
					confidence: sigmoid64(tc),
					classes:    getOrderedElements(softmax(tclasses)),
				}
				counter++
			}
		}
	}
	sort.Sort(sort.Reverse(byConfidence(boxes)))
	printClassification(boxes)
	f, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}
	m := image.NewNRGBA(img.Bounds())

	draw.Draw(m, m.Bounds(), img, image.ZP, draw.Src)
	for _, b := range boxes {
		if b.confidence > drawingThreshold {
			drawRectangle(m, b.r, fmt.Sprintf("%v %2.2f%%", b.classes[0].class, b.classes[0].prob*100))
		}
	}

	if err := png.Encode(f, m); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

func printClassification(classification []box) {
	var classes []element
	for _, e := range classification {
		if e.confidence > threshold {
			if e.classes[0].prob > threshold {
				classes = append(classes, e.classes...)
				fmt.Printf("at (%v) with confidence %2.2f%%: %v\n", e.r, e.confidence, e.classes[:3])
			}
		}
	}
	sort.Sort(sort.Reverse(byProba(classes)))
	for _, e := range classes {
		if e.prob > 0.4 {
			fmt.Println(e)
		}
	}

}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type box struct {
	r          image.Rectangle
	gridcell   []int
	confidence float64
	classes    []element
}

type byProba []element

func (a byProba) Len() int           { return len(a) }
func (a byProba) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byProba) Less(i, j int) bool { return a[i].prob < a[j].prob }

type byGridCell []box
type byConfidence []box

func (a byGridCell) Len() int      { return len(a) }
func (a byGridCell) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byGridCell) Less(i, j int) bool {
	return a[i].gridcell[0] < a[j].gridcell[0] || a[i].gridcell[1] < a[j].gridcell[1]
}

func (a byConfidence) Len() int           { return len(a) }
func (a byConfidence) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byConfidence) Less(i, j int) bool { return a[i].confidence < a[j].confidence }

func sigmoid(sum float32) float32 {
	return float32(1.0 / (1.0 + math.Exp(float64(-sum))))
}
func sigmoid64(sum float32) float64 {
	return 1.0 / (1.0 + math.Exp(float64(-sum)))
}
func exp(val float32) float64 {
	return math.Exp(float64(val))
}

func softmax(a []float32) []float64 {
	var sum float64
	output := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		output[i] = math.Exp(float64(a[i]))
		sum += output[i]
	}
	for i := 0; i < len(output); i++ {
		output[i] = output[i] / sum
	}
	return output
}

type element struct {
	prob  float64
	class string
}

func getOrderedElements(input []float64) []element {
	elems := make([]element, len(input))
	for i := 0; i < len(elems); i++ {
		elems[i] = element{
			prob:  input[i],
			class: classes[i],
		}
	}
	sort.Sort(sort.Reverse(elements(elems)))
	return elems
}

type elements []element

func (a elements) Len() int           { return len(a) }
func (a elements) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a elements) Less(i, j int) bool { return a[i].prob < a[j].prob }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func drawRectangle(img *image.NRGBA, r image.Rectangle, label string) {
	col := color.RGBA{255, 0, 0, 255} // Red

	// HLine draws a horizontal line
	hLine := func(x1, y, x2 int) {
		for ; x1 <= x2; x1++ {
			img.Set(x1, y, col)
		}
	}

	// VLine draws a veritcal line
	vLine := func(x, y1, y2 int) {
		for ; y1 <= y2; y1++ {
			img.Set(x, y1, col)
		}
	}

	// Rect draws a rectangle utilizing HLine() and VLine()
	rect := func(r image.Rectangle) {
		hLine(r.Min.X, r.Max.Y, r.Max.X)
		hLine(r.Min.X, r.Min.Y, r.Max.X)
		vLine(r.Max.X, r.Min.Y, r.Max.Y)
		vLine(r.Min.X, r.Min.Y, r.Max.Y)
	}
	addLabel(img, r.Bounds().Min.X+5, r.Bounds().Min.Y+15, label)
	rect(r)
}

func addLabel(img *image.NRGBA, x, y int, label string) {
	col := color.NRGBA{0, 255, 0, 255}
	point := fixed.Point26_6{
		X: fixed.Int26_6(x * 64),
		Y: fixed.Int26_6(y * 64),
	}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
