package onnx

import (
	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"gonum.org/v1/gonum/graph"
)

// Backend represent any backend able to receive a computation graph
type Backend interface {
	OperationCarrier
	graph.DirectedWeightedBuilder
}

// Operation defined by its name and its attribute
type Operation struct {
	Name       string
	Attributes []*pb.AttributeProto
}

// OperationCarrier should be amethod of the graph
// because the operation needs the topology of the graph
// to check the arity of the node for example
type OperationCarrier interface {
	ApplyOperation(Operation, graph.Node) error
}
