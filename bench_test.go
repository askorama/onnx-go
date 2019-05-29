package onnx_test

import (
	"testing"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
	"github.com/owulveryck/onnx-go/internal/examples/mnist"
	"gorgonia.org/tensor"
)

func BenchmarkUnmarshalBinary(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Create a backend receiver
		backend := gorgonnx.NewGraph()
		// Create a model and set the execution backend
		model := onnx.NewModel(backend)

		// Decode it into the model
		err := model.UnmarshalBinary(mnist.GetMnist())
		if err != nil {
			b.Fatal(err)
		}
		// Set the first input, the number depends of the model
		input := tensor.New(tensor.WithShape(1, 1, 28, 28), tensor.Of(tensor.Float32))
		model.SetInput(0, input)
		err = backend.Run()
		if err != nil {
			b.Fatal(err)
		}
		// Check error
	}
}
