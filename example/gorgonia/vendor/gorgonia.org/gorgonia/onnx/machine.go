package onnx

import "gorgonia.org/gorgonia/internal/engine"

// NewTapeMachine ...
func NewTapeMachine(g *Graph) engine.VM {
	return engine.NewTapeMachine((*g).ExprGraph)
}
