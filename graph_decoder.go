package onnx

import (
	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"gonum.org/v1/gonum/graph"
)

// Unmarshal onnx encoded model proto data into a graph builder
func Unmarshal(data []byte, dst graph.DirectedBuilder) error {
	model := new(pb.ModelProto)
	err := model.Unmarshal(data)

	return err
}
