package main

import (
	"flag"
	"fmt"
	"image"
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
	"gorgonia.org/tensor"
	"gorgonia.org/tensor/native"
)

// The 416x416 image is divided into a 13x13 grid. Each of these grid cells
// will predict 5 bounding boxes (boxesPerCell). A bounding box consists of
// five data items: x, y, width, height, and a confidence score. Each grid
// cell also predicts which class each bounding box belongs to.
//
const (
	h, w         = 416, 416
	blockSize    = 32
	gridHeight   = 13
	gridWidth    = 13
	boxesPerCell = 5
	numClasses   = 20
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
	anchors = []float32{1.08, 1.19, 3.42, 4.41, 6.63, 11.38, 9.42, 5.11, 16.62, 10.52}
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
		inputT := tensor.New(tensor.WithShape(1, 3, 416, 416), tensor.Of(tensor.Float32))
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

	var classification = make([]box, gridHeight*gridWidth*boxesPerCell)
	var counter int
	// https://github.com/pjreddie/darknet/blob/61c9d02ec461e30d55762ec7669d6a1d3c356fb2/src/yolo_layer.c#L159
	for b := 0; b < len(data)-25; b += 25 {
		for cx := 0; cx < gridHeight; cx++ {
			for cy := 0; cy < gridWidth; cy++ {
				tx := data[b][cx][cy]
				ty := data[b+1][cx][cy]
				tw := data[b+2][cx][cy]
				th := data[b+3][cx][cy]
				tc := data[b+4][cx][cy]
				tclasses := data[b+5 : b+24][cx][cy]

				classification[counter] = box{
					gridcell:    []int{cx, cy},
					boundindBox: b,
					// The predicted tx and ty coordinates are relative to the location
					// of the grid cell; we use the logistic sigmoid to constrain these
					// coordinates to the range 0 - 1. Then we add the cell coordinates
					// (0-12) and multiply by the number of pixels per grid cell (32).
					// Now x and y represent center of the bounding box in the original
					// 416x416 image space.
					// https://github.com/hollance/Forge/blob/04109c856237faec87deecb55126d4a20fa4f59b/Examples/YOLO/YOLO/YOLO.swift#L154
					x: (float32(cx) + sigmoid(tx)) * blockSize,
					y: (float32(cy) + sigmoid(ty)) * blockSize,
					// The size of the bounding box, tw and th, is predicted relative to
					// the size of an "anchor" box. Here we also transform the width and
					// height into the original 416x416 image space.
					//w:          exp(tw) * anchors[2*b] * blockSize,
					//h:          exp(th) * anchors[2*b+1] * blockSize,
					w:          tw, // TODO use the anchor
					h:          th, // TODO use the anchor
					confidence: sigmoid64(tc),
					classes:    getOrderedElements(softmax(tclasses)),
				}
				counter++
			}
		}
	}
	sort.Sort(sort.Reverse(byConfidence(classification)))
	//sort.Sort(sort.Reverse(byGridCell(classification)))
	for _, e := range classification[:15] {
		fmt.Printf("%v: %v\n", e.confidence, e.classes[:3])
	}
	f, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}
	m := image.NewRGBA(img.Bounds())

	draw.Draw(m, m.Bounds(), img, image.ZP, draw.Src)

	if err := png.Encode(f, m); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type box struct {
	gridcell    []int
	boundindBox int
	x, y, w, h  float32
	confidence  float64
	classes     []element
}

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
func exp(val float32) float32 {
	return float32(math.Exp(float64(val)))
}

func softmax(a []float32) []float32 {
	max := -float32(math.MaxFloat32)
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	output := make([]float32, len(a))
	var sum float32
	for i := 0; i < len(a); i++ {
		sum += float32(math.Exp(float64(a[i] - max)))
	}
	for i := 0; i < len(a); i++ {
		output[i] = float32(math.Exp(float64(a[i]-max))) / sum
	}
	return output
}

type element struct {
	prob  float32
	class string
}

func getOrderedElements(input []float32) []element {
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
