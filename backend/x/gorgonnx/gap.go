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

func (g *gap) apply(gg *Graph, ns ...*Node) error {
	n := ns[0]
	children := getOrderedChildren(gg.g, n)
	var err error
	if len(children) != 1 {
		return errors.New("GlobalAveragePool: bad arity")
	}

	// Temporary, waiting for the operator to be implemented in Gorgonia
	// see https://github.com/gorgonia/gorgonia/pull/302
	n.gorgoniaNode, err = gorgonia.GlobalAveragePool2D(children[0].gorgoniaNode)
	return err
}

func (*gap) init(onnx.Operation) error {
	return nil
}
