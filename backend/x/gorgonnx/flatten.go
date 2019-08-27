package gorgonnx

import (
	"errors"

	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
)

func init() {
	register("Flatten", newFlatten)
}

func newFlatten() operator {
	return &flatten{}
}

type flatten struct {
	axis int64
}

func (f *flatten) apply(g *Graph, ns ...*Node) error {
	n := ns[0]
	children := getOrderedChildren(g.g, n)
	var err error
	if len(children) != 1 {
		return errors.New("flatten: bad arity")
	}
	s := make([]int, len(children[0].gorgoniaNode.Shape()))
	copy(s, children[0].gorgoniaNode.Shape())
	var x, y int
	x = 1
	y = 1
	for i := 0; i < len(s); i++ {
		if i < int(f.axis) {
			x *= s[i]
		}
		if i >= int(f.axis) {
			y *= s[i]
		}
	}
	n.gorgoniaNode, err = gorgonia.Reshape(children[0].gorgoniaNode, []int{x, y})
	return err
}

func (f *flatten) init(o onnx.Operation) error {
	f.axis = 1
	axis, ok := o.Attributes["axis"]
	if ok {
		if f.axis, ok = axis.(int64); !ok {
			return errors.New("axis is not an int64")
		}
	}
	return nil
}
