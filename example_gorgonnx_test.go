package onnx_test

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
	"gorgonia.org/tensor"
)

var input tensor.Tensor

func Example_gorgonia() {
	// Create a backend receiver
	backend := gorgonnx.NewGraph()
	// Create a model and set the execution backend
	model := onnx.NewModel(backend)

	// read the onnx model
	b, _ := ioutil.ReadFile("model.onnx")
	// Decode it into the model
	err := model.UnmarshalBinary(b)
	if err != nil {
		log.Fatal(err)
	}
	// Set the first input, the number depends of the model
	model.SetInput(0, input)
	err = backend.Run()
	if err != nil {
		log.Fatal(err)
	}
	// Check error
	output, _ := model.GetOutputTensors()
	// write the first output to stdout
	fmt.Println(output[0])
}
