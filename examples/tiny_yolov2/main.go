package main

import (
	"flag"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
	"github.com/owulveryck/onnx-go/internal/x/images"
	"gorgonia.org/tensor"
)

var (
	model   = flag.String("model", "model.onnx", "path to the model file")
	inputT  = flag.String("inputT", "", "path of an input tensor for testing")
	imgF    = flag.String("img", "", "path of an input tensor for testing")
	img     image.Image
	classes = []string{"aeroplane", "bicycle", "bird", "boat", "bottle",
		"bus", "car", "cat", "chair", "cow",
		"diningtable", "dog", "horse", "motorbike", "person",
		"pottedplant", "sheep", "sofa", "train", "tv/monitor"}
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
	switch {
	case *inputT != "":
		if _, err := os.Stat(*inputT); err != nil && os.IsNotExist(err) {
			log.Fatalf("%v does not exist", *inputT)
		}
		b, err := ioutil.ReadFile(*inputT)
		if err != nil {
			log.Fatal(err)
		}
		inputT, err := onnx.NewTensor(b)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(inputT)
		img, err = images.TensorToImg(inputT)
		if err != nil {
			log.Println(err)
		}
		return inputT

	case *imgF != "":
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
	default:
		log.Fatal("Please speficy an input")
	}
	return nil
}

func processOutput(t []tensor.Tensor, err error) {
	if err != nil {
		log.Fatal(err)
	}
	log.Println(t[0].Data())
	f, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
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
