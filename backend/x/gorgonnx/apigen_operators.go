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

	x, y, err := broadcast(children[0], children[1])
	if err != nil {
		err, ok := err.(*onnx.ErrNotImplemented)
		if ok {
			err.Operator = "Mul / hadamardProd"
		}
		return err
	}
	n.gorgoniaNode, err = gorgonia.HadamardProd(x, y)

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

	x, y, err := broadcast(children[0], children[1])
	if err != nil {
		err, ok := err.(*onnx.ErrNotImplemented)
		if ok {
			err.Operator = "Div / hadamardDiv"
		}
		return err
	}
	n.gorgoniaNode, err = gorgonia.HadamardDiv(x, y)

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

	x, y, err := broadcast(children[0], children[1])
	if err != nil {
		err, ok := err.(*onnx.ErrNotImplemented)
		if ok {
			err.Operator = "Sub / sub"
		}
		return err
	}
	n.gorgoniaNode, err = gorgonia.Sub(x, y)

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

	x, y, err := broadcast(children[0], children[1])
	if err != nil {
		err, ok := err.(*onnx.ErrNotImplemented)
		if ok {
			err.Operator = "Add / add"
		}
		return err
	}
	n.gorgoniaNode, err = gorgonia.Add(x, y)

	return err
}

func (a *add) init(o onnx.Operation) error {
	return nil
}


type abs struct{}

func init() {
	register("Abs", &abs{})
}

func (a *abs) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Abs(
		children[0].gorgoniaNode,
	)

	return err
}

func (a *abs) init(o onnx.Operation) error {
	return nil
}


type sign struct{}

func init() {
	register("Sign", &sign{})
}

func (a *sign) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Sign(
		children[0].gorgoniaNode,
	)

	return err
}

func (a *sign) init(o onnx.Operation) error {
	return nil
}


type ceil struct{}

func init() {
	register("Ceil", &ceil{})
}

func (a *ceil) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Ceil(
		children[0].gorgoniaNode,
	)

	return err
}

func (a *ceil) init(o onnx.Operation) error {
	return nil
}


type floor struct{}

func init() {
	register("Floor", &floor{})
}

func (a *floor) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Floor(
		children[0].gorgoniaNode,
	)

	return err
}

func (a *floor) init(o onnx.Operation) error {
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


type exp struct{}

func init() {
	register("Exp", &exp{})
}

func (a *exp) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Exp(
		children[0].gorgoniaNode,
	)

	return err
}

func (a *exp) init(o onnx.Operation) error {
	return nil
}


type logarithm struct{}

func init() {
	register("Log", &logarithm{})
}

func (a *logarithm) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Log(
		children[0].gorgoniaNode,
	)

	return err
}

func (a *logarithm) init(o onnx.Operation) error {
	return nil
}


type log2 struct{}

func init() {
	register("Log2", &log2{})
}

func (a *log2) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Log2(
		children[0].gorgoniaNode,
	)

	return err
}

func (a *log2) init(o onnx.Operation) error {
	return nil
}


type relu struct{}

func init() {
	register("Relu", &relu{})
}

func (a *relu) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Rectify(
		children[0].gorgoniaNode,
	)

	return err
}

func (a *relu) init(o onnx.Operation) error {
	return nil
}


type neg struct{}

func init() {
	register("Neg", &neg{})
}

func (a *neg) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Neg(
		children[0].gorgoniaNode,
	)

	return err
}

func (a *neg) init(o onnx.Operation) error {
	return nil
}


type square struct{}

func init() {
	register("Square", &square{})
}

func (a *square) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Square(
		children[0].gorgoniaNode,
	)

	return err
}

func (a *square) init(o onnx.Operation) error {
	return nil
}


type sqrt struct{}

func init() {
	register("Sqrt", &sqrt{})
}

func (a *sqrt) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Sqrt(
		children[0].gorgoniaNode,
	)

	return err
}

func (a *sqrt) init(o onnx.Operation) error {
	return nil
}


type inverse struct{}

func init() {
	register("Inverse", &inverse{})
}

func (a *inverse) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Inverse(
		children[0].gorgoniaNode,
	)

	return err
}

func (a *inverse) init(o onnx.Operation) error {
	return nil
}


type cube struct{}

func init() {
	register("Cube", &cube{})
}

func (a *cube) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Cube(
		children[0].gorgoniaNode,
	)

	return err
}

func (a *cube) init(o onnx.Operation) error {
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


type sigmoid struct{}

func init() {
	register("Sigmoid", &sigmoid{})
}

func (a *sigmoid) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Sigmoid(
		children[0].gorgoniaNode,
	)

	return err
}

func (a *sigmoid) init(o onnx.Operation) error {
	return nil
}


type log1p struct{}

func init() {
	register("Log1p", &log1p{})
}

func (a *log1p) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Log1p(
		children[0].gorgoniaNode,
	)

	return err
}

func (a *log1p) init(o onnx.Operation) error {
	return nil
}


type expm1 struct{}

func init() {
	register("Expm1", &expm1{})
}

func (a *expm1) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Expm1(
		children[0].gorgoniaNode,
	)

	return err
}

func (a *expm1) init(o onnx.Operation) error {
	return nil
}


type softplus struct{}

func init() {
	register("Softplus", &softplus{})
}

func (a *softplus) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	n.gorgoniaNode, err = gorgonia.Softplus(
		children[0].gorgoniaNode,
	)

	return err
}

func (a *softplus) init(o onnx.Operation) error {
	return nil
}

