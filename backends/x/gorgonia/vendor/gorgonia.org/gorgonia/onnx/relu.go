package onnx

import (
	"gonum.org/v1/gonum/graph"
	"gorgonia.org/gorgonia/internal/engine"
)

// Rectify operator ...
type Rectify struct{}

// Constructor to fulfil the interface ...
func (r *Rectify) Constructor() func(g graph.WeightedDirected, n graph.Node) (interface{}, error) {
	return func(g graph.WeightedDirected, n graph.Node) (interface{}, error) {
		return engine.NewRectifyOperation()(g, n.(*engine.Node))
	}
}
