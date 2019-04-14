package gorgonnx

import (
	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
)

type hadamardProd struct{}

func init() {
	register("Mul", &hadamardProd{})
}

func (a *hadamardProd) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 2)
	if err != nil {
		return err
	}

	if len(children[0].gorgoniaNode.Shape()) != len(children[1].gorgoniaNode.Shape()) {
		return &onnx.ErrNotImplemented{
			Operator: n.operation.Name,
			Message:  "broadcast not yet implemented",
		}

	}

	n.gorgoniaNode, err = gorgonia.HadamardProd(
		children[0].gorgoniaNode,
		children[1].gorgoniaNode,
	)
	return err
}

func (a *hadamardProd) init(o onnx.Operation) error {
	return nil
}


type hadamardDiv struct{}

func init() {
	register("Div", &hadamardDiv{})
}

func (a *hadamardDiv) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 2)
	if err != nil {
		return err
	}

	if len(children[0].gorgoniaNode.Shape()) != len(children[1].gorgoniaNode.Shape()) {
		return &onnx.ErrNotImplemented{
			Operator: n.operation.Name,
			Message:  "broadcast not yet implemented",
		}

	}

	n.gorgoniaNode, err = gorgonia.HadamardDiv(
		children[0].gorgoniaNode,
		children[1].gorgoniaNode,
	)
	return err
}

func (a *hadamardDiv) init(o onnx.Operation) error {
	return nil
}


type sub struct{}

func init() {
	register("Sub", &sub{})
}

func (a *sub) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 2)
	if err != nil {
		return err
	}

	if len(children[0].gorgoniaNode.Shape()) != len(children[1].gorgoniaNode.Shape()) {
		return &onnx.ErrNotImplemented{
			Operator: n.operation.Name,
			Message:  "broadcast not yet implemented",
		}

	}

	n.gorgoniaNode, err = gorgonia.Sub(
		children[0].gorgoniaNode,
		children[1].gorgoniaNode,
	)
	return err
}

func (a *sub) init(o onnx.Operation) error {
	return nil
}


type add struct{}

func init() {
	register("Add", &add{})
}

func (a *add) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 2)
	if err != nil {
		return err
	}

	if len(children[0].gorgoniaNode.Shape()) != len(children[1].gorgoniaNode.Shape()) {
		return &onnx.ErrNotImplemented{
			Operator: n.operation.Name,
			Message:  "broadcast not yet implemented",
		}

	}

	n.gorgoniaNode, err = gorgonia.Add(
		children[0].gorgoniaNode,
		children[1].gorgoniaNode,
	)
	return err
}

func (a *add) init(o onnx.Operation) error {
	return nil
}


type cos struct{}

func init() {
	register("Cos", &cos{})
}

func (a *cos) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Cos(
		children[0].gorgoniaNode,
	)
	return err
}

func (a *cos) init(o onnx.Operation) error {
	return nil
}


type sin struct{}

func init() {
	register("Sin", &sin{})
}

func (a *sin) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Sin(
		children[0].gorgoniaNode,
	)
	return err
}

func (a *sin) init(o onnx.Operation) error {
	return nil
}


type tanh struct{}

func init() {
	register("Tanh", &tanh{})
}

func (a *tanh) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Tanh(
		children[0].gorgoniaNode,
	)
	return err
}

func (a *tanh) init(o onnx.Operation) error {
	return nil
}

