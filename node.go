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

// Tensor ...
type Tensor interface {
	graph.Node
	SetValue(t tensor.Tensor) error
	GetValue() tensor.Tensor
}
