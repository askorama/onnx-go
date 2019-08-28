package gorgonnx

import (
	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
)

// SPEC: https://github.com/onnx/onnx/blob/master/docs/Operators.md#BatchNormalization
// Gorgonia implem: https://godoc.org/gorgonia.org/gorgonia#BatchNorm

type unsqueeze struct {
	Axes []int64
}

func init() {
	register("Unsqueeze", newUnsqueeze)
}
func newUnsqueeze() operator {
	return &unsqueeze{}
}

func (a *unsqueeze) apply(g *Graph, ns ...*Node) error {
	n := ns[0]
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}

	tensor := children[0].gorgoniaNode
	dims := make([]int, len(a.Axes)+tensor.Dims())
	for k := range dims {
		dims[k] = -1
	}
	for _, v := range a.Axes {
		dims[v] = 1
	}
	var index int
	for k, v := range dims {
		if v == -1 {
			index = k
			break
		}
	}
	for i := 0; i < tensor.Dims(); i++ {
		dims[i+index] = tensor.Shape()[i]
	}
	n.gorgoniaNode, err = gorgonia.Reshape(tensor, dims)

	return err

}

func (a *unsqueeze) init(o onnx.Operation) error {
	return nil
}
