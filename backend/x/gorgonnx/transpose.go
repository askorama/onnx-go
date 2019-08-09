package gorgonnx

import (
	"errors"

	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
)

func init() {
	register("Transpose", newTranspose)
}

func newTranspose() operator {
	return &transpose{
		perm: make([]int64, 0),
	}
}

// transpose to be compatible with:
//    https://github.com/onnx/onnx/blob/master/docs/Operators.md#Transpose
// and
//    https://godoc.org/gorgonia.org/gorgonia#Transpose
// test with go test -run=TestONNX/Transpose
type transpose struct {
	perm []int64
}

func (t *transpose) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	var err error
	if len(children) != 1 {
		return errors.New("Transpose: bad arity")
	}
	perm := make([]int, len(t.perm))
	for i := 0; i < len(perm); i++ {
		perm[i] = int(t.perm[i])
	}
	n.gorgoniaNode, err = gorgonia.Transpose(children[0].gorgoniaNode, perm...)
	return err
}

func (t *transpose) init(o onnx.Operation) error {
	var ok bool
	var perm interface{}
	if perm, ok = o.Attributes["perm"]; !ok {
		return &onnx.ErrNotImplemented{
			Operator:      "Transpose",
			AttributeName: "perm",
			Message:       "enpty perm (Default transpose) not implemented yet",
		}
	}
	if t.perm, ok = perm.([]int64); !ok {
		return errors.New("transpose: perm is not an array of int")
	}
	return nil
}
