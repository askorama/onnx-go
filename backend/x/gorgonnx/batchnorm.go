package gorgonnx

import (
	"errors"

	"github.com/davecgh/go-spew/spew"
	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
)

// SPEC: https://github.com/onnx/onnx/blob/master/docs/Operators.md#BatchNormalization
// Gorgonia implem: https://godoc.org/gorgonia.org/gorgonia#BatchNorm

type batchnorm struct {
	epsilon  float64
	momentum float64
}

func init() {
	register("BatchNormalization", newBatchNorm)
}
func newBatchNorm() operator {
	return &batchnorm{}
}

func (b *batchnorm) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 5)
	if err != nil {
		return err
	}
	spew.Dump(children[1].t.Data())
	n.gorgoniaNode, _, _, _, err = gorgonia.BatchNorm(children[0].gorgoniaNode,
		children[1].gorgoniaNode,
		children[2].gorgoniaNode,
		b.momentum,
		b.epsilon)
	return err
	/*
		return &onnx.ErrNotImplemented{
			Operator: "BatchNormalization",
		}
	*/
}

func (b *batchnorm) init(o onnx.Operation) error {
	b.epsilon = 1e-5
	b.momentum = 0.9
	if e, ok := o.Attributes["epsilon"]; ok {
		if v, ok := e.(float32); ok {
			b.epsilon = float64(v)
		} else {
			return errors.New("epsilon is not a float64")
		}
	}
	if e, ok := o.Attributes["momentum"]; ok {
		if v, ok := e.(float32); ok {
			b.momentum = float64(v)
		} else {
			return errors.New("momentum is not a float64")
		}
	}
	return nil
}
