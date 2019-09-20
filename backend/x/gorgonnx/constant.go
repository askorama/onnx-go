package gorgonnx

import (
	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
)

type constant struct {
	value *gorgonia.Node
}

func init() {
	register("Constant", newConstant)
}

func newConstant() operator {
	return &constant{}
}

func (a *constant) apply(g *Graph, ns ...*Node) error {
	n := ns[0]
	n.gorgoniaNode = a.value
	return nil
}

func (a *constant) init(o onnx.Operation) error {
	a.value = gorgonia.NewConstant(o.Attributes["value"])
	return nil
}
