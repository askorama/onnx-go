package gorgonnx

import (
	"github.com/owulveryck/onnx-go"
	"gorgonia.org/tensor"
)

func init() {
	//register("Conv", &conv{})
}

// conv to be compatible with:
//    https://github.com/onnx/onnx/blob/master/docs/Operators.md#Conv
// and
//    https://godoc.org/gorgonia.org/gorgonia#Conv2d
type conv struct {
	pad         []int
	stride      []int
	dilation    []int
	group       int
	kernelShape tensor.Shape
}

func (c *conv) apply(g *Graph, n *Node) error {
	return nil
}
func (c *conv) init(o onnx.Operation) error {
	return nil
}
