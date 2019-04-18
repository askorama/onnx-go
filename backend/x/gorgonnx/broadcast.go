package gorgonnx

import (
	"github.com/owulveryck/onnx-go"
)

func broadcast(a, b *Node) (*Node, *Node, error) {
	if len(a.gorgoniaNode.Shape()) != len(b.gorgoniaNode.Shape()) {
		return a, b, &onnx.ErrNotImplemented{
			Message: "broadcast not yet implemented",
		}

	}
	return a, b, nil

}
