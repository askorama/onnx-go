package gorgonnx

import (
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/chewxy/hm"
	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

func init() {
	register("GlobalAveragePool", newGAP)
}

func newGAP() operator {
	return &gap{}
}

type gap struct{}

func (g *gap) apply(gg *Graph, ns ...*Node) error {
	n := ns[0]
	children := getOrderedChildren(gg.g, n)
	var err error
	if len(children) != 1 {
		return errors.New("GlobalAveragePool: bad arity")
	}

	// Temporary, waiting for the operator to be implemented in Gorgonia
	// see https://github.com/gorgonia/gorgonia/pull/302
	/*
		n.gorgoniaNode, err = gorgonia.GlobalAveragePool2D(children[0].gorgoniaNode)
	*/
	n.gorgoniaNode, err = gorgonia.ApplyOp(g, children[0].gorgoniaNode)
	return err
}

func (*gap) init(onnx.Operation) error {
	return nil
}

func (g *gap) Arity() int {
	return 1
}

func (g *gap) Type() hm.Type {
	t := gorgonia.TensorType{Dims: 4, Of: hm.TypeVariable('a')}
	return hm.NewFnType(t, t)
}

func (g *gap) InferShape(inputs ...gorgonia.DimSizer) (tensor.Shape, error) {
	b, err := inputs[0].DimSize(0)
	if err != nil {
		return nil, err
	}
	c, err := inputs[0].DimSize(1)
	if err != nil {
		return nil, err
	}
	// check if the shape is correct without doing type inference
	if _, err := inputs[0].DimSize(2); err != nil {
		return nil, err
	}
	if _, err := inputs[0].DimSize(3); err != nil {
		return nil, err
	}
	return tensor.Shape{b, c, 1, 1}, nil
}

func (g *gap) Do(inputs ...gorgonia.Value) (gorgonia.Value, error) {
	im := inputs[0]
	switch im.(type) {
	case tensor.Tensor:
		v := im.(tensor.Tensor)
		B, C, H, W := v.Shape()[0], v.Shape()[1], v.Shape()[2], v.Shape()[3]
		s, err := g.InferShape(v.Shape())
		if err != nil {
			return nil, err
		}
		output := tensor.New(tensor.Of(v.Dtype()), tensor.WithShape(s...))
		switch v.Dtype() {
		case tensor.Float64:
			err = setFloat64AtTensor(v, B, C, H, W, output)
			if err != nil {
				return nil, err
			}
		case tensor.Float32:
			err = setFloat32AtTensor(v, B, C, H, W, output)
			if err != nil {
				return nil, err
			}
		default:
			return nil, &onnx.ErrNotImplemented{
				Operator: "Global Average Pool",
				Message:  fmt.Sprintf("%v not implemented", v.Dtype()),
			}
		}

		return output, nil

	default:
		return nil, &onnx.ErrNotImplemented{
			Operator: "Global Average Pool",
			Message:  fmt.Sprintf("invalid input %v", inputs),
		}
	}
}

func setFloat64AtTensor(v tensor.Tensor, B, C, H, W int, output tensor.Tensor) error {
	for b := 0; b < B; b++ {
		for c := 0; c < C; c++ {
			var sum float64
			for h := 0; h < H; h++ {
				for w := 0; w < W; w++ {
					val, err := v.At(b, c, h, w)
					if err != nil {
						return err
					}
					sum += val.(float64)
				}
			}
			err := output.SetAt(sum/float64(H*W), b, c, 0, 0)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func setFloat32AtTensor(v tensor.Tensor, B, C, H, W int, output tensor.Tensor) error {
	for b := 0; b < B; b++ {
		for c := 0; c < C; c++ {
			var sum float32
			for h := 0; h < H; h++ {
				for w := 0; w < W; w++ {
					val, err := v.At(b, c, h, w)
					if err != nil {
						return err
					}
					sum += val.(float32)
				}
			}
			err := output.SetAt(sum/float32(H*W), b, c, 0, 0)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *gap) ReturnsPtr() bool {
	return false
}

func (g *gap) CallsExtern() bool {
	return false
}

func (g *gap) OverwritesInput() int {
	return -1
}

func (g *gap) WriteHash(h hash.Hash) {
	fmt.Fprintf(h, "GlobalAveragePool")
}

func (g *gap) Hashcode() uint32 {
	h := fnv.New32a()
	g.WriteHash(h)
	return h.Sum32()
}

func (g *gap) String() string {
	return "GlobalAveragePool"
}
