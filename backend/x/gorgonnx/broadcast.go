package gorgonnx

import (
	"fmt"

	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
)

func broadcast(a, b *Node) (*gorgonia.Node, *gorgonia.Node, error) {
	return ggnBroadcast(a.gorgoniaNode, b.gorgoniaNode)
}

func ggnBroadcast(a, b *gorgonia.Node) (*gorgonia.Node, *gorgonia.Node, error) {
	if sameDim(a, b) {
		return a, b, nil
	}
	return ggnReshapedBroadcast(a, b)
}

func ggnReshapedBroadcast(a, b *gorgonia.Node) (*gorgonia.Node, *gorgonia.Node, error) {
	// for NCHW tensors, the first dimension may be omitted and must be broadcasted
	// TODO find a smarter way to achieve this
	reshapedA := a
	reshapedB := b
	var err error
	switch {
	case len(a.Shape()) == 0:
		reshapedA, err = reshapeNode(a, b, false)
	case len(b.Shape()) == 0:
		reshapedB, err = reshapeNode(b, a, false)
	case len(a.Shape()) == 1 && len(b.Shape()) != 1:
		reshapedA, err = reshapeNode(a, b, true)
	case len(a.Shape()) != 1 && len(b.Shape()) == 1:
		reshapedB, err = reshapeNode(b, a, true)
	case len(a.Shape()) == 2 && len(b.Shape()) == 2:
		// No reshaping needed
	case len(a.Shape()) == 3 && len(b.Shape()) == 4:
		reshapedA, err = gorgonia.Reshape(a, append([]int{1}, a.Shape()...))
	case len(a.Shape()) == 4 && len(b.Shape()) == 3:
		reshapedB, err = gorgonia.Reshape(b, append([]int{1}, b.Shape()...))
	default:
		return a, b, &onnx.ErrNotImplemented{
			Message: fmt.Sprintf("broadcast not yet implemented for shape %v, %v", a.Shape(), b.Shape()),
		}

	}
	if err != nil {
		return nil, nil, err
	}
	return gorgonia.Broadcast(reshapedA, reshapedB, getBroadcastPattern(reshapedA, reshapedB))
}

func reshapeNode(n1, n2 *gorgonia.Node, findAxis bool) (*gorgonia.Node, error) {
	dimN2 := n2.Shape()
	dimReshapeN1 := make([]int, len(dimN2))
	for i := 0; i < len(dimN2); i++ {
		dimReshapeN1[i] = 1
		if findAxis && dimN2[i] == n1.Shape()[0] {
			// Make an educated guess: find the axis that has the same dimension
			dimReshapeN1[i] = dimN2[i]
		}
	}
	reshapeN1, err := gorgonia.Reshape(n1, dimReshapeN1)
	if err != nil {
		return nil, err
	}
	return reshapeN1, nil
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
