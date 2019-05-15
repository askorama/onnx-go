package gorgonnx

import (
	"errors"

	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
)

type softmax struct {
	axis int
}

func init() {
	register("Softmax", &softmax{})
}

func (s *softmax) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}
	a := children[0].gorgoniaNode
	var max, reshaped *gorgonia.Node
	if len(a.Shape()) > 2 {
		if s.axis > len(a.Shape()) {
			return errors.New("softmax cannot be applied on an axis > len(shape()) of the input")
		}
		row := 1
		col := 1
		for i, shape := range a.Shape() {
			if i < s.axis {
				row = row * shape
			} else {
				col = col * shape
			}
		}
		reshaped, err = gorgonia.Reshape(a, []int{row, col})
		if err != nil {
			return err
		}
	} else {
		reshaped = a
	}
	if max, err = gorgonia.Max(reshaped); err != nil {
		return err
	}
	a2, b2, err := gorgonia.Broadcast(reshaped, max, gorgonia.NewBroadcastPattern(nil, []byte{0, 1}))
	if err != nil {
		return err
	}
	output, err := gorgonia.Sub(a2, b2)
	if err != nil {
		return err
	}
	var exp, sum *gorgonia.Node
	if exp, err = gorgonia.Exp(output); err == nil {
		axis := 1
		if exp.IsScalar() {
			axis = 0
		}
		if sum, err = gorgonia.Sum(exp, axis); err == nil {
			if sum.IsScalar() {
				n.gorgoniaNode, err = gorgonia.HadamardDiv(exp, sum)
				return err
			}
			a, b, err := gorgonia.Broadcast(exp, sum, gorgonia.NewBroadcastPattern(nil, []byte{1}))
			if err != nil {
				return err
			}
			n.gorgoniaNode, err = gorgonia.HadamardDiv(a, b)
			return err
		}
		return err
	}
	return err
}

func (s *softmax) init(o onnx.Operation) error {
	axis, ok := o.Attributes["axis"]
	if !ok {
		s.axis = 1
		return nil
	}
	err := errors.New("axis in not an int")
	if axis, ok := axis.(int64); ok {
		s.axis = int(axis)
		err = nil
	}
	return err
}
