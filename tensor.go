package onnx

import (
	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"gorgonia.org/tensor"
)

// NewTensor from onnx value
func NewTensor(b []byte) (tensor.Tensor, error) {
	tp := new(pb.TensorProto)
	err := tp.XXX_Unmarshal(b)
	if err != nil {
		return nil, err
	}
	return tp.Tensor()
}
