package gorgonnx

import (
	"fmt"

	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
)

func broadcast(a, b *Node) (*gorgonia.Node, *gorgonia.Node, error) {
	if sameDim(a.gorgoniaNode, b.gorgoniaNode) {
		return a.gorgoniaNode, b.gorgoniaNode, nil
	}
	// for NCHW tensors, the first dimension may be ommited and must be broadcasted
	// TODO find a smarter way to achieve this
	switch {
	case len(a.gorgoniaNode.Shape()) == 1 && len(b.gorgoniaNode.Shape()) != 1:
		//TODO
		return a.gorgoniaNode, b.gorgoniaNode, &onnx.ErrNotImplemented{
			Message: fmt.Sprintf("broadcast not yet implemented for shape %v, %v", a.gorgoniaNode.Shape(), b.gorgoniaNode.Shape()),
		}
	case len(a.gorgoniaNode.Shape()) != 1 && len(b.gorgoniaNode.Shape()) == 1:
		//TODO
		return a.gorgoniaNode, b.gorgoniaNode, &onnx.ErrNotImplemented{
			Message: fmt.Sprintf("broadcast not yet implemented for shape %v, %v", a.gorgoniaNode.Shape(), b.gorgoniaNode.Shape()),
		}
	case len(a.gorgoniaNode.Shape()) == 3 && len(b.gorgoniaNode.Shape()) == 4:
		// Reshape node a
		aR, err := gorgonia.Reshape(a.gorgoniaNode, append([]int{1}, a.gorgoniaNode.Shape()...))
		if err != nil {
			return nil, nil, err
		}
		return gorgonia.Broadcast(aR, b.gorgoniaNode, getBroadcastPattern(aR, b.gorgoniaNode))
	case len(a.gorgoniaNode.Shape()) == 4 && len(b.gorgoniaNode.Shape()) == 3:
		// Reshape node a
		bR, err := gorgonia.Reshape(b.gorgoniaNode, append([]int{1}, b.gorgoniaNode.Shape()...))
		if err != nil {
			return nil, nil, err
		}
		return gorgonia.Broadcast(a.gorgoniaNode, bR, getBroadcastPattern(a.gorgoniaNode, bR))
	default:
		return a.gorgoniaNode, b.gorgoniaNode, &onnx.ErrNotImplemented{
			Message: fmt.Sprintf("broadcast not yet implemented for shape %v, %v", a.gorgoniaNode.Shape(), b.gorgoniaNode.Shape()),
		}

	}
}

func sameDim(a, b *gorgonia.Node) bool {
	if len(a.Shape()) != len(b.Shape()) {
		return false
	}
	for i := 0; i < len(a.Shape()); i++ {
		if a.Shape()[i] != b.Shape()[i] {
			return false
		}
	}
	return true
}

func getBroadcastPattern(a, b *gorgonia.Node) gorgonia.BroadcastPattern {
	var leftAxes, rightAxes []byte
	for i := 0; i < len(a.Shape()); i++ {
		switch {
		case a.Shape()[i] == 1 && b.Shape()[i] != 1:
			leftAxes = append(leftAxes, byte(i))
		case a.Shape()[i] != 1 && b.Shape()[i] == 1:
			rightAxes = append(rightAxes, byte(i))
		}
	}
	return gorgonia.NewBroadcastPattern(leftAxes, rightAxes)

}
