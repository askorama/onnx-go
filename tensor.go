package onnx

import (
	"github.com/owulveryck/onnx-go/internal/onnx/ir"
	"gorgonia.org/tensor"
)

// NewTensor from onnx value
func NewTensor(b []byte) (tensor.Tensor, error) {
	tp := new(ir.TensorProto)
	err := tp.XXX_Unmarshal(b)
	if err != nil {
		return nil, err
	}
	return tp.Tensor()
}
