package gorgonnx

import (
	"errors"

	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
)

type leakyRELU struct {
	alpha float32
}

func init() {
	register("LeakyRelu", newLeakyRELU)
}

func newLeakyRELU() operator {
	return &leakyRELU{}
}

func (l *leakyRELU) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	var nodes = make([]*gorgonia.Node, len(children))
	for i := 0; i < len(children); i++ {
		nodes[i] = children[i].gorgoniaNode
	}
	var err error
	n.gorgoniaNode, err = gorgonia.LeakyRelu(nodes[0], float64(l.alpha))
	return err
}

func (l *leakyRELU) init(o onnx.Operation) error {
	l.alpha = 0.01
	if alpha, ok := o.Attributes["alpha"]; ok {
		if alpha, ok := alpha.(float32); ok {
			l.alpha = alpha
			return nil
		}
		return errors.New("alpha in not a float32")
	}
	return nil
}
