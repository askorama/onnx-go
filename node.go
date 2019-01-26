package onnx

import (
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
