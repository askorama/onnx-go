package onnx

import (
	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"gonum.org/v1/gonum/graph/encoding"
)

// Unmarshal onnx encoded model proto data into a graph builder
func Unmarshal(data []byte, dst encoding.Builder) error {
	model := new(pb.ModelProto)
	err := model.Unmarshal(data)
	return err
}
