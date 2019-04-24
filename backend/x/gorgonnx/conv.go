package gorgonnx

import (
	"github.com/owulveryck/onnx-go"
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
	return &onnx.ErrNotImplemented{
		Operator: "Conv",
	}
}

func (c *conv) init(o onnx.Operation) error {
	autoPad, ok := o.Attributes["auto_pad"]
	if ok && autoPad.(string) != "NOTSET" {
		return &onnx.ErrNotImplemented{
			Operator: "Conv",
			Message:  "auto_pad " + autoPad.(string) + " not implemented",
		}
	}
	// ex: "kernel_shape":[]int64{3, 3}, "pads":[]int64{1, 0, 1, 0}, "strides":[]int64{2, 2}, "auto_pad": string{"NOTSET"}
	return nil
}
