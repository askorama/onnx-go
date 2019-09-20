package gorgonnx

import (
	"errors"

	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
)

// Element-wise min of each of the input tensors (with Numpy-style broadcasting support)
// Specifications: https://github.com/onnx/onnx/blob/master/docs/Operators.md#Min
type min struct{}

func init() {
	register("Min", newmin)
}

func newmin() operator {
	return &min{}
}

func (a *min) apply(g *Graph, n ...*Node) error {
	if len(n) != 1 {
		return errors.New("wrong number of input nodes")
	}
	children := getOrderedChildren(g.g, n[0])
	err := checkMinimumCondition(children, 1)
	if err != nil {
		return err
	}

	output := children[0].gorgoniaNode
	others := children[1:]

	for _, child := range others {
		x, y, err := ggnBroadcast(output, child.gorgoniaNode)
		if err != nil {
			err, ok := err.(*onnx.ErrNotImplemented)
			if ok {
				err.Operator = "Min"
			}
			return err
		}

		comp, err := gorgonia.Lt(x, y, true)
		if err != nil {
			return err
		}
		leftTerm, err := gorgonia.HadamardProd(x, comp)
		if err != nil {
			return err
		}
		c := gorgonia.NewConstant(float32(1.0))
		complementComp, err := gorgonia.Sub(c, comp)
		if err != nil {
			return err
		}
		rightTerm, err := gorgonia.HadamardProd(y, complementComp)
		if err != nil {
			return err
		}
		output, err = gorgonia.Add(leftTerm, rightTerm)
		if err != nil {
			return err
		}
	}

	n[0].gorgoniaNode = output
	return err
}

func (a *min) init(o onnx.Operation) error {
	return nil
}

// Element-wise max of each of the input tensors (with Numpy-style broadcasting support)
// Specifications: https://github.com/onnx/onnx/blob/master/docs/Operators.md#Max
type max struct{}

func init() {
	register("Max", newmax)
}

func newmax() operator {
	return &max{}
}

func (a *max) apply(g *Graph, n ...*Node) error {
	if len(n) != 1 {
		return errors.New("wrong number of input nodes")
	}
	children := getOrderedChildren(g.g, n[0])
	err := checkMinimumCondition(children, 1)
	if err != nil {
		return err
	}

	output := children[0].gorgoniaNode
	others := children[1:]

	for _, child := range others {
		x, y, err := ggnBroadcast(output, child.gorgoniaNode)
		if err != nil {
			err, ok := err.(*onnx.ErrNotImplemented)
			if ok {
				err.Operator = "Max"
			}
			return err
		}

		comp, err := gorgonia.Gt(x, y, true)
		if err != nil {
			return err
		}
		leftTerm, err := gorgonia.HadamardProd(x, comp)
		if err != nil {
			return err
		}
		c := gorgonia.NewConstant(float32(1.0))
		complementComp, err := gorgonia.Sub(c, comp)
		if err != nil {
			return err
		}
		rightTerm, err := gorgonia.HadamardProd(y, complementComp)
		if err != nil {
			return err
		}
		output, err = gorgonia.Add(leftTerm, rightTerm)
		if err != nil {
			return err
		}
	}

	n[0].gorgoniaNode = output
	return err
}

func (a *max) init(o onnx.Operation) error {
	return nil
}
