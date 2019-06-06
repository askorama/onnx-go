package gorgonnx

import (
	"errors"

	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

func init() {
	register("MaxPool", newMaxpool)
}

func newMaxpool() operator {
	return &maxpool{}
}

// maxpool to be compatible with:
//    https://github.com/onnx/onnx/blob/master/docs/Operators.md#maxpool
// and
//    https://godoc.org/gorgonia.org/gorgonia#MaxPool2D
// test with go test -run=TestONNX/maxpool
type maxpool struct {
	padding     string
	pad         []int
	stride      []int
	kernelShape tensor.Shape
}

func ceilDivInt(a, b int) int {
	return (a + b - 1) / b
}

func (c *maxpool) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}
	x := children[0].gorgoniaNode
	switch c.padding {
	case "SAME_UPPER":
		outputSpatialShape := make([]int, len(x.Shape()[2:]))
		for i, v := range x.Shape()[2:] {
			outputSpatialShape[i] = ceilDivInt(v, c.stride[i])
			// pad_shape[i] = (output_spatial_shape[i] - 1) * strides_spatial_shape[i] + kernel_spatial_shape[i] - input_spatial_shape[i]
			c.pad[i] = (outputSpatialShape[i]-1)*c.stride[i] + c.kernelShape[i] - v
		}
	default:
	}
	n.gorgoniaNode, err = gorgonia.MaxPool2D(
		children[0].gorgoniaNode,
		c.kernelShape,
		c.pad,
		c.stride)

	return err
}

func (c *maxpool) init(o onnx.Operation) error {
	var autoPad string
	if a, ok := o.Attributes["auto_pad"]; ok {
		if autoPad, ok = a.(string); !ok {
			return errors.New("autopad is not a string")
		}

	}
	switch autoPad {
	case "NOTSET":
	case "":
	default:
		return &onnx.ErrNotImplemented{
			Operator: "maxpool",
			Message:  "auto_pad " + autoPad + " not implemented",
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
	if len(c.kernelShape) != 2 {
		return &onnx.ErrNotImplemented{
			Operator: "maxpool",
			Message:  "maxpool for dim >2 not implemented",
		}
	}

	c.pad = []int{0, 0}
	pad, ok := o.Attributes["pads"]
	if ok {
		if pad, ok := pad.([]int64); ok {

			if len(pad) == 4 && (pad[0] != pad[1] || pad[2] != pad[3]) {
				return &onnx.ErrNotImplemented{
					Operator:       "maxpool",
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
	_, ok = o.Attributes["ceil_mode"]
	if ok {
		return &onnx.ErrNotImplemented{
			Operator: "maxpool",
			Message:  "ceil_mode not implemented",
		}
	}
	_, ok = o.Attributes["dilations"]
	if ok {
		return &onnx.ErrNotImplemented{
			Operator: "maxpool",
			Message:  "dilation not implemented",
		}
	}
	return nil
}
