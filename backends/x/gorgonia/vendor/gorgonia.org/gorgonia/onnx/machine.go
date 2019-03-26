package onnx

import "gorgonia.org/gorgonia/internal/engine"

// NewTapeMachine ...
func NewTapeMachine(g *Graph, opts ...engine.VMOpt) engine.VM {
	return engine.NewTapeMachine((*g).ExprGraph, opts...)
}
