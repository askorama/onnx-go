package onnx_test

import (
	"log"
	"os"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/simple"
)

func Example_simple() {
	// START SIMPLE
	// Create a backend receiver
	backend := simple.NewSimpleGraph()
	// Create a model and set the execution backend
	model := onnx.NewModel(backend)

	// read the onnx model
	b, _ := os.ReadFile("model.onnx")
	// Decode it into the model
	err := model.UnmarshalBinary(b)
	// END SIMPLE
	if err != nil {
		log.Fatal(err)
	}

}
