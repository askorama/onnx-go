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
	var exp, sum *gorgonia.Node
	if exp, err = gorgonia.Exp(a); err == nil {
		if sum, err = gorgonia.Sum(exp, s.axis); err == nil {
			if sum.IsScalar() {
				n.gorgoniaNode, err = gorgonia.HadamardDiv(exp, sum)
				return err
			}
			a, b, err := gorgonia.Broadcast(exp, sum, gorgonia.NewBroadcastPattern(nil, []byte{1}))
			if err != nil {
				return err
			}
			n.gorgoniaNode, err = gorgonia.Div(a, b)
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
