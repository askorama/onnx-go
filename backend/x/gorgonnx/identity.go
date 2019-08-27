package gorgonnx

import "github.com/owulveryck/onnx-go"

type identity struct{}

func init() {
	register("Identity", newIdentity)
}

func newIdentity() operator {
	return &identity{}
}

func (a *identity) apply(g *Graph, ns ...*Node) error {
	n := ns[0]
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode = children[0].gorgoniaNode
	return err
}

func (a *identity) init(o onnx.Operation) error {
	return nil
}
