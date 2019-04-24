package gorgonnx

import (
	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

func init() {
	register("Conv", &conv{})
}

// conv to be compatible with:
//    https://github.com/onnx/onnx/blob/master/docs/Operators.md#Conv
// and
//    https://godoc.org/gorgonia.org/gorgonia#Conv2d
// test with go test -run=TestONNX/Conv
type conv struct {
	pad         []int
	stride      []int
	dilation    []int
	group       int
	kernelShape tensor.Shape
}

func (c *conv) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 2)
	if err != nil {
		return err
	}
	n.gorgoniaNode, err = gorgonia.Conv2d(
		children[0].gorgoniaNode,
		children[1].gorgoniaNode,
		c.kernelShape,
		c.pad,
		c.stride,
		c.dilation)
	return err
}

func (c *conv) init(o onnx.Operation) error {
	autoPad, ok := o.Attributes["auto_pad"]
	if ok && autoPad.(string) != "NOTSET" {
		return &onnx.ErrNotImplemented{
			Operator: "Conv",
			Message:  "auto_pad " + autoPad.(string) + " not implemented",
		}
	}
	kernelShape, ok := o.Attributes["kernel_shape"]
	if ok {
		if kernelShape, ok := kernelShape.([]int64); ok {
			c.kernelShape = make([]int, len(kernelShape))
			for i := 0; i < len(kernelShape); i++ {
				c.kernelShape[i] = int(kernelShape[i])
			}
		}
	}
	c.pad = []int{0, 0}
	pad, ok := o.Attributes["pads"]
	if ok {
		if pad, ok := pad.([]int64); ok {

			if len(pad) == 4 && (pad[0] != pad[1] || pad[2] != pad[3]) {
				return &onnx.ErrNotImplemented{
					Operator:       "Conv",
					AttributeName:  "pads",
					AttributeValue: pad,
					Message:        "Asymetric padding",
				}
			}

			if len(pad) == 4 {
				for i := 0; i < 2; i++ {
					c.pad[i] = int(pad[2*i])
				}
			} else if len(pad) == 2 {
				for i := 0; i < 2; i++ {
					c.pad[i] = int(pad[i])
				}
			}
		}
	}
	c.stride = []int{1, 1}
	stride, ok := o.Attributes["strides"]
	if ok {
		if stride, ok := stride.([]int64); ok {
			if len(stride) == 4 {
				for i := 0; i < 2; i++ {
					c.stride[i] = int(stride[2*i])
				}
			} else if len(stride) == 2 {
				for i := 0; i < 2; i++ {
					c.stride[i] = int(stride[i])
				}
			}
		}
	}
	dilation, ok := o.Attributes["dilations"]
	if ok {
		if dilation, ok := dilation.([]int64); ok {
			c.dilation = make([]int, len(dilation))
			for i := 0; i < len(dilation); i++ {
				c.dilation[i] = int(dilation[i])
			}
		}
	}
	return nil
}
