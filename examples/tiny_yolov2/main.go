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
	s := make([]int, 3)
	copy(s, dense.Shape()[1:])
	// Ignore the first dimension
	err = dense.Reshape(s...)
	if err != nil {
		log.Fatal(err)
	}

	features, err := native.Tensor3F32(dense)
	if err != nil {
		log.Fatal(err)
	}
	var classification = make([]box, 13*13*5)
	var counter int
	for b := 0; b < len(features); b += 25 {
		for cx := 0; cx < len(features[b]); cx++ {
			for cy := 0; cy < len(features[b][cx]); cy++ {
				bb := (b + 1) / 25
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
					x: (float32(cx) + sigmoid(features[b][cx][cy])) * blockSize,
					y: (float32(cy) + sigmoid(features[b+1][cx][cy])) * blockSize,
					// The size of the bounding box, tw and th, is predicted relative to
					// the size of an "anchor" box. Here we also transform the width and
					// height into the original 416x416 image space.
					w:          exp(features[b+2][cx][cy]) * anchors[2*bb] * blockSize,
					h:          exp(features[b+3][cx][cy]) * anchors[2*bb+1] * blockSize,
					confidence: sigmoid(features[b+4][cx][cy]),
					classes:    softmax(features[b+5 : b+24][cx][cy]),
				}
				counter++
			}
		}
	}
	sort.Sort(sort.Reverse(byConfidence(classification)))
	//sort.Sort(sort.Reverse(byGridCell(classification)))
	for _, e := range classification[:6] {
		fmt.Println(e)
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
	gridcell               []int
	boundindBox            int
	x, y, w, h, confidence float32
	classes                []float32
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
func exp(val float32) float32 {
	return float32(math.Exp(float64(val)))
}

func softmax(a []float32) []float32 {
	output := make([]float32, len(a))
	var sum float32
	for i := 0; i < len(a); i++ {
		sum += float32(math.Exp(float64(a[i])))
	}
	for i := 0; i < len(a); i++ {
		output[i] = float32(math.Exp(float64(a[i]))) / sum
	}
	return output
}
