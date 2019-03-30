package onnx

import (
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
	Attributes map[string]interface{}
}

// OperationCarrier should be a method of the graph
// because the operation needs the topology of the graph
// to check the arity of the node for example
type OperationCarrier interface {
	ApplyOperation(Operation, graph.Node) error
}
