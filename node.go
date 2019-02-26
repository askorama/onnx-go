package onnx

import (
	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"gonum.org/v1/gonum/graph"
	"gorgonia.org/tensor"
)

// Namer is a node that know its own name
type Namer interface {
	graph.Node
	SetName(string)
	GetName() string
}

// Documenter is an interface that describe any object able to document itself
type Documenter interface {
	graph.Node
	SetDescription(string)
	GetDescription() string
}

// TensorCarrier is a graph which can turn a node into a tensor
type TensorCarrier interface {
	ApplyTensor(t tensor.Tensor) error
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
