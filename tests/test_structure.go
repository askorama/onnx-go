package tests

import (
	"github.com/owulveryck/onnx-go/internal/pb-onnx"
	"gorgonia.org/tensor"
)

// TestCase describes an integration test
type TestCase struct {
	Title   string
	Model   *pb.ModelProto
	Inputs  []tensor.Tensor
	Outputs []tensor.Tensor
}
