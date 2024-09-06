package gorgonnx

import (
	"errors"
	"github.com/owulveryck/onnx-go/internal/onnx/ir"
	"gorgonia.org/tensor"

	"github.com/owulveryck/onnx-go"
)

type cast struct {
	to ir.TensorProto_DataType
}

func init() {
	register("Cast", newCast)
}

func newCast() operator {
	return &cast{}
}

func (op *cast) apply(g *Graph, ns ...*Node) error {
	if len(ns) != 1 {
		return errors.New("cast: invalid number of nodes")
	}
	n := ns[0]

	ts := n.GetTensor()
	if ts == nil {
		return nil
	}
	ts, err := ts.Apply(func(a tensor.Tensor) (tensor.Tensor, error) {
		ts := tensor.New(
			tensor.Of(op.toDtype()),
			tensor.WithShape(a.Shape()...),
			tensor.WithEngine(a.Engine()),
		)
		it := a.Iterator()
		for !it.Done() {
			idx, err := it.Next()
			if err != nil {
				return nil, err
			}
			val, err := a.At(idx)
			if err != nil {
				return nil, err
			}
			err = ts.SetAt(val, it.Coord()...)
			if err != nil {
				return nil, err
			}
		}

		return nil, nil
	})
	if err != nil {
		return err
	}
	return n.SetTensor(ts)
}

func (op *cast) init(o onnx.Operation) error {
	to, ok := o.Attributes["to"]
	if !ok {
		return errors.New("cast: expected 'to' attribute is not found")
	}
	v, ok := to.(int64)
	if !ok {
		return errors.New("cast: expected 'to' attribute is not an int64")
	}
	op.to = ir.TensorProto_DataType(v)
	return nil
}

func (op *cast) toDtype() tensor.Dtype {
	switch op.to {
	case ir.TensorProto_FLOAT:
		return tensor.Float32
	case ir.TensorProto_DOUBLE:
		return tensor.Float64
	case ir.TensorProto_INT8:
		return tensor.Int8
	case ir.TensorProto_INT16:
		return tensor.Int16
	case ir.TensorProto_INT32:
		return tensor.Int32
	case ir.TensorProto_INT64:
		return tensor.Int64
	case ir.TensorProto_UINT8:
		return tensor.Uint8
	case ir.TensorProto_UINT16:
		return tensor.Uint16
	case ir.TensorProto_UINT32:
		return tensor.Uint32
	case ir.TensorProto_UINT64:
		return tensor.Uint64
	case ir.TensorProto_BOOL:
		return tensor.Bool
	case ir.TensorProto_STRING:
		return tensor.String
	default:
		panic("cast: unknown dtype")
	}
}
