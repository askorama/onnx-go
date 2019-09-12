package onnx

import (
	"math"

	"gonum.org/v1/gonum/graph"
)

// SelfEdge is the weight of a self edge in the graph
const SelfEdge = math.MaxFloat64

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
	// ApplyOperation on the graph nodes
	// graph.Node is an array because it allows to handle multiple output
	// for example a split operation returns n nodes...
	ApplyOperation(Operation, ...graph.Node) error
}
