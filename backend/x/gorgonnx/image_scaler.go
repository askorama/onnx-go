package gorgonnx

import (
	"errors"

	"github.com/google/uuid"
	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

type imageScaler struct {
	bias  []float32
	scale float32
}

func init() {
	register("ImageScaler", newImageScaler)
}

func newImageScaler() operator {
	return &imageScaler{}
}

func (a *imageScaler) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	if len(children) != 1 {
		return errors.New("ImageScaler: bad number of children")
	}
	x := children[0].gorgoniaNode
	if x.Dtype() != tensor.Float32 {
		return &onnx.ErrNotImplemented{
			Operator: "ImageScaler",
			Message:  "Only float32 is supported",
		}
	}
	if len(x.Shape()) != 4 {
		return errors.New("Expected a 4D tensor [N,C,H,W]")
	}
	if len(a.bias) != x.Shape()[1] {
		return errors.New("bias should be the same size as the channel")
	}
	if a.scale != float32(1) {
		return &onnx.ErrNotImplemented{
			Operator:       "ImageScaler",
			AttributeName:  "scale",
			AttributeValue: a.scale,
		}
	}
	biasT := tensor.New(tensor.WithBacking(a.bias), tensor.Of(tensor.Float32))
	bias := gorgonia.NodeFromAny(g.exprgraph, biasT, gorgonia.WithName(uuid.New().String()))
	ax, bx, err := gorgonia.Broadcast(x, bias, gorgonia.NewBroadcastPattern(nil, []byte{0, 2, 3}))
	if err != nil {
		return err
	}
	n.gorgoniaNode, err = gorgonia.Add(ax, bx)
	if err != nil {
		return err
	}
	return err
}

func (a *imageScaler) init(o onnx.Operation) error {
	a.scale = 1

	bias, ok := o.Attributes["bias"]
	if !ok {
		return errors.New("imageScaler: expected bias attribute is not found")
	}
	err := errors.New("bias in not a []float32")
	if bias, ok := bias.([]float32); ok {
		a.bias = []float32(bias)
		err = nil
	}
	if err != nil {
		return err
	}

	if scale, ok := o.Attributes["scale"]; ok {
		err = errors.New("scale in not a float32")
		if scale, ok := scale.(float32); ok {
			a.scale = scale
			err = nil
		}
	}
	return err
}
