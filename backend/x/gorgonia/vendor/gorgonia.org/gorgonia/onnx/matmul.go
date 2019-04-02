package onnx

import (
	"gonum.org/v1/gonum/graph"
	"gorgonia.org/gorgonia/internal/engine"
)

// Mul performs a multiplication
type Mul struct{}

// Constructor to fulfil the interface ...
func (a *Mul) Constructor() func(g graph.WeightedDirected, n graph.Node) (interface{}, error) {
	return func(g graph.WeightedDirected, n graph.Node) (interface{}, error) {
		return engine.NewMulOperation()(g, n.(*engine.Node))
	}
}
