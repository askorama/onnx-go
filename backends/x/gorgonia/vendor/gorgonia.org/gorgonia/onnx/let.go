package onnx

import (
	"gorgonia.org/gorgonia/internal/engine"
	"gorgonia.org/gorgonia/node"
)

// Let binds a value.Value to a node that is a variable. A variable is represented as a *Node with no Op.
// It is equivalent to :
//		x = 2
func Let(n node.Node, be interface{}) error {
	return engine.UnsafeLet(n.(*engine.Node), be)
}
