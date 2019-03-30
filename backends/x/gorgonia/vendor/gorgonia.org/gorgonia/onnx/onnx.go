package onnx

import (
	onnx "github.com/owulveryck/onnx-go"
	"gonum.org/v1/gonum/graph"
	"gorgonia.org/gorgonia/internal/engine"
	"gorgonia.org/gorgonia/node"
	"gorgonia.org/gorgonia/ops"
)

//START_DOC OMIT

// Graph ...
type Graph struct {
	*engine.ExprGraph
}

// NewGraph ...
func NewGraph() *Graph {
	return &Graph{engine.NewGraph()}
}

type operation interface {
	Constructor() func(g graph.WeightedDirected, n graph.Node) (interface{}, error)
}

// ApplyOperation ...
func (g Graph) ApplyOperation(op onnx.Operation, n graph.Node) error {
	s := op.Name
	var o operation
	switch s {
	case "MaxPool":
		o = NewMaxpool()
	case "Conv":
		o = NewConv()
	case "Relu":
		o = &Rectify{}
	case "Reshape":
		o = &Reshape{}
	case "Abs":
		o = &Abs{}
	case "Sign":
		o = &Sign{}
	case "Ceil":
		o = &Ceil{}
	case "Floor":
		o = &Floor{}
	case "Sin":
		o = &Sin{}
	case "Cos":
		o = &Cos{}
	case "Exp":
		o = &Exp{}
	case "Log":
		o = &Log{}
	case "Log2":
		o = &Log2{}
	case "Neg":
		o = &Neg{}
	case "Square":
		o = &Square{}
	case "Sqrt":
		o = &Sqrt{}
	case "Inverse":
		o = &Inverse{}
	case "InverseSqrt":
		o = &InverseSqrt{}
	case "Cube":
		o = &Cube{}
	case "Tanh":
		o = &Tanh{}
	case "Sigmoid":
		o = &Sigmoid{}
	case "Log1p":
		o = &Log1p{}
	case "Expm1":
		o = &Expm1{}
	case "Softplus":
		o = &Softplus{}
	case "Add":
		o = &Add{}
	case "Sub":
		o = &Sub{}
	case "MatMul":
		o = &Mul{}
	case "HadamardDiv":
		o = &HadamardDiv{}
	case "Pow":
		o = &Pow{}
	case "Lt":
		o = &Lt{}
	case "Gt":
		o = &Gt{}
	case "Lte":
		o = &Lte{}
	case "Gte":
		o = &Gte{}
	case "Eq":
		o = &Eq{}
	case "Ne":
		o = &Ne{}
	default:
		return &ErrNotImplemented{
			Operator: s,
		}
	}
	// TODO UnmarshalAttributes
	// TODO check the pointer
	err := onnx.UnmarshalAttributes(op.Attributes, o)
	if err != nil {
		return err
	}
	return g.apply(o.Constructor(), n)
}

// apply ...
func (g Graph) apply(operation func(g graph.WeightedDirected, n graph.Node) (interface{}, error), n graph.Node) error { // HL
	oper := func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		output, err := operation(g, n)
		if output == nil {
			panic(err)
		}
		return output.(ops.Op), err
	}
	return g.ApplyOp(engine.Operation(oper), n.(*engine.Node))
}

// END_DOC OMIT
