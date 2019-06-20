package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"

	"github.com/kelseyhightower/envconfig"
	"github.com/nfnt/resize"
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
	hSize, wSize  = 416, 416
	blockSize     = 32
	gridHeight    = 13
	gridWidth     = 13
	boxesPerCell  = 5
	numClasses    = 20
	envConfPrefix = "yolo"
)

type configuration struct {
	ConfidenceThreshold float64 `envconfig:"confidence_threshold" default:"0.30" required:"true"`
	ClassProbaThreshold float64 `envconfig:"proba_threshold" default:"0.90" required:"true"`
}

func init() {
	err := envconfig.Process(envConfPrefix, &config)
	if err != nil {
		panic(err)
	}
}

var (
	model   = flag.String("model", "model.onnx", "path to the model file")
	imgF    = flag.String("img", "", "path of an input jpeg image (use - for stdin)")
	outputF = flag.String("output", "", "path of an output png file (use - for stdout)")
	silent  = flag.Bool("s", false, "silent mode (useful if output is -)")
	img     image.Image
	classes = []string{"aeroplane", "bicycle", "bird", "boat", "bottle",
		"bus", "car", "cat", "chair", "cow",
		"diningtable", "dog", "horse", "motorbike", "person",
		"pottedplant", "sheep", "sofa", "train", "tv/monitor"}
	anchors     = []float64{1.08, 1.19, 3.42, 4.41, 6.63, 11.38, 9.42, 5.11, 16.62, 10.52}
	scaleFactor = float32(1) // The scale factor to resize the image to hSize*wSize
	config      configuration
)

func main() {
	h := flag.Bool("h", false, "help")
	flag.Parse()
	if *h {
		flag.Usage()
		envconfig.Usage(envConfPrefix, &config)
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
	if *imgF == "" {
		flag.Usage()
		os.Exit(1)
	}
	var f io.Reader
	var err error
	if *imgF == "-" {
		f = os.Stdin
	} else {
		f, err = os.Open(*imgF)
		if err != nil {
			log.Fatal(err)
		}
		defer f.(*os.File).Close()
	}
	img, err = jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	// find the resize scale
	imgRescaled := image.NewNRGBA(image.Rect(0, 0, wSize, hSize))
	color := color.RGBA{0, 0, 0, 255}

	draw.Draw(imgRescaled, imgRescaled.Bounds(), &image.Uniform{color}, image.ZP, draw.Src)
	var m image.Image
	if (img.Bounds().Max.X - img.Bounds().Min.X) > (img.Bounds().Max.Y - img.Bounds().Min.Y) {
		scaleFactor = float32(img.Bounds().Max.Y-img.Bounds().Min.Y) / float32(hSize)
		m = resize.Resize(0, hSize, img, resize.Lanczos3)
	} else {
		scaleFactor = float32(img.Bounds().Max.X-img.Bounds().Min.X) / float32(wSize)
		m = resize.Resize(wSize, 0, img, resize.Lanczos3)
	}
	switch m.(type) {
	case *image.NRGBA:
		draw.Draw(imgRescaled, imgRescaled.Bounds(), m.(*image.NRGBA), image.ZP, draw.Src)
	case *image.YCbCr:
		draw.Draw(imgRescaled, imgRescaled.Bounds(), m.(*image.YCbCr), image.ZP, draw.Src)
	default:
		log.Fatal("unhandled type")
	}

	inputT := tensor.New(tensor.WithShape(1, 3, hSize, wSize), tensor.Of(tensor.Float32))
	//err = images.ImageToBCHW(img, inputT)
	err = images.ImageToBCHW(imgRescaled, inputT)
	if err != nil {
		log.Fatal(err)
	}
	return inputT
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
	boxes = sanitize(boxes)
	if !*silent {
		printClassification(boxes)
	}
	if *outputF != "" {
		drawClassification(boxes)
	}
}

func printClassification(boxes []box) {
	var elements []element
	for _, box := range boxes {
		if box.classes[0].prob > config.ConfidenceThreshold {
			elements = append(elements, box.classes...)
			fmt.Printf("at (%v) with confidence %2.2f%%: %v\n", box.r, box.confidence, box.classes[:3])
		}
	}
	sort.Sort(sort.Reverse(byProba(elements)))
	for _, c := range elements {
		if c.prob > 0.4 {
			fmt.Println(c)
		}
	}

}
func drawClassification(boxes []box) {
	if *outputF == "" {
		return
	}
	var f io.Writer
	var err error
	if *outputF == "-" {
		f = os.Stdout
	} else {
		f, err = os.Create(*outputF)
		if err != nil {
			log.Fatal(err)
		}
		defer f.(*os.File).Close()
	}
	m := image.NewNRGBA(img.Bounds())

	draw.Draw(m, m.Bounds(), img, image.ZP, draw.Src)
	for _, b := range boxes {
		drawRectangle(m, b.r, fmt.Sprintf("%v %2.2f%%", b.classes[0].class, b.classes[0].prob*100))
	}

	if err := png.Encode(f, m); err != nil {
		log.Fatal(err)
	}

}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type element struct {
	prob  float64
	class string
}

type byProba []element

func (b byProba) Len() int           { return len(b) }
func (b byProba) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byProba) Less(i, j int) bool { return b[i].prob < b[j].prob }

type box struct {
	r          image.Rectangle
	gridcell   []int
	confidence float64
	classes    []element
}

type byConfidence []box

func (b byConfidence) Len() int           { return len(b) }
func (b byConfidence) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byConfidence) Less(i, j int) bool { return b[i].confidence < b[j].confidence }

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

func getOrderedElements(input []float64) []element {
	elems := make([]element, len(input))
	for i := 0; i < len(elems); i++ {
		elems[i] = element{
			prob:  input[i],
			class: classes[i],
		}
	}
	sort.Sort(sort.Reverse(byProba(elems)))
	return elems
}

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

	minX := int(float32(r.Min.X) * scaleFactor)
	maxX := int(float32(r.Max.X) * scaleFactor)
	minY := int(float32(r.Min.Y) * scaleFactor)
	maxY := int(float32(r.Max.Y) * scaleFactor)
	// Rect draws a rectangle utilizing HLine() and VLine()
	rect := func(r image.Rectangle) {
		hLine(minX, maxY, maxX)
		hLine(minX, maxY, maxX)
		hLine(minX, minY, maxX)
		vLine(maxX, minY, maxY)
		vLine(minX, minY, maxY)
	}
	addLabel(img, minX+5, minY+15, label)
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

// from https://medium.com/@jonathan_hui/real-time-object-detection-with-yolo-yolov2-28b1b93e2088
// 1- Sort the predictions by the confidence scores.
// 2- Start from the top scores, ignore any current prediction if we find any previous predictions that have the same class and IoU > 0.5 with the current prediction.
// 3- Repeat step 2 until all predictions are checked.
func sanitize(boxes []box) []box {
	sort.Sort(sort.Reverse(byConfidence(boxes)))

	for i := 1; i < len(boxes); i++ {
		if boxes[i].confidence < config.ConfidenceThreshold {
			boxes = boxes[:i]
			break
		}
		if boxes[i].classes[0].prob < config.ClassProbaThreshold {
			boxes = boxes[:i]
			break
		}
		for j := i + 1; j < len(boxes); {
			iou := iou(boxes[i].r, boxes[j].r)
			if iou > 0.5 && boxes[i].classes[0].class == boxes[j].classes[0].class {
				boxes = append(boxes[:j], boxes[j+1:]...)
				continue
			}
			j++
		}
	}
	return boxes
}

// evaluate the intersection over union of two rectangles
func iou(r1, r2 image.Rectangle) float64 {
	// get the intesection rectangle
	intersection := image.Rect(
		max(r1.Min.X, r2.Min.X),
		max(r1.Min.Y, r2.Min.Y),
		min(r1.Max.X, r2.Max.X),
		min(r1.Max.Y, r2.Max.Y),
	)
	// compute the area of intersection rectangle
	interArea := area(intersection)
	r1Area := area(r1)
	r2Area := area(r2)
	// compute the intersection over union by taking the intersection
	// area and dividing it by the sum of prediction + ground-truth
	// areas - the interesection area
	return float64(interArea) / float64(r1Area+r2Area-interArea)
}

func area(r image.Rectangle) int {
	return max(0, r.Max.X-r.Min.X-1) * max(0, r.Max.Y-r.Min.Y-1)
}
