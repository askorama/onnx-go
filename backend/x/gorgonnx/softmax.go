package gorgonnx

import (
	"errors"

	"github.com/owulveryck/onnx-go"
)

type softmax struct {
	axis int
}

func init() {
	register("Softmax", &softmax{})
}

func (a *softmax) apply(g *Graph, n *Node) error {
	return &onnx.ErrNotImplemented{
		Operator: "Softmax",
	}
}

func (a *softmax) init(o onnx.Operation) error {
	axis, ok := o.Attributes["axis"]
	if !ok {
		a.axis = 1
		return nil
	}
	err := errors.New("axis in not an int")
	if axis, ok := axis.(int64); ok {
		a.axis = int(axis)
		err = nil
	}
	return err
}
