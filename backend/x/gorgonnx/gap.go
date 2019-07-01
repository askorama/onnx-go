package gorgonnx

import (
	"errors"

	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
)

func init() {
	register("GlobalAveragePool", newGAP)
}

func newGAP() operator {
	return &gap{}
}

type gap struct{}

func (o *gap) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	var err error
	if len(children) != 1 {
		return errors.New("GlobalAveragePool: bad arity")
	}

	n.gorgoniaNode, err = gorgonia.GlobalAveragePool2D(children[0].gorgoniaNode)
	return err
}

func (*gap) init(onnx.Operation) error {
	return nil
}
