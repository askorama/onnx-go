package onnx

import (
	"gonum.org/v1/gonum/graph"
	"gorgonia.org/gorgonia/internal/engine"
	"gorgonia.org/tensor"
)

// Maxpool operator ...
type Maxpool struct {
	AutoPad      string  `attributeName:"auto_pad"`
	StorageOrder int64   `attributeName:"storage_order"`
	KernelShape  []int64 `attributeName:"kernel_shape" required:"true"`
	Pads         []int64 `attributeName:"pads"`
	Strides      []int64 `attributeName:"strides"`
	strides      []int
	pads         []int
	kernelShape  tensor.Shape
}

// NewMaxpool with default values
func NewMaxpool() *Maxpool {
	return &Maxpool{
		AutoPad:      "NOTSET",
		StorageOrder: 0,
		Strides:      []int64{1, 1},
		Pads:         []int64{0, 0},
	}
}

// Constructor to fulfil the interface ...
func (m *Maxpool) Constructor() func(g graph.WeightedDirected, n graph.Node) (interface{}, error) {
	return func(g graph.WeightedDirected, n graph.Node) (interface{}, error) {
		if m.StorageOrder != 0 {
			return nil, &ErrNotImplemented{
				Operator:       "Maxpool",
				AttributeName:  "storage_order",
				AttributeValue: m.StorageOrder,
				Message:        "attribute not implemented for a value != 0",
			}
		}
		switch m.AutoPad {
		case "NOTSET":
		case "VALID":
			m.pads = []int{0, 0}
		case "SAME_UPPER":

			return nil, &ErrNotImplemented{
				Operator:       "Maxpool",
				AttributeName:  "auto_pad",
				AttributeValue: m.AutoPad,
				Message:        "Padding is buggy",
			}
			/*
				outputHeight := int(math.Ceil(float64(input[0].Shape()[2]) / float64(o.Strides[0])))
				outputWidth := int(math.Ceil(float64(input[0].Shape()[3]) / float64(o.Strides[1])))
				o.Pads[0] = int(math.Max(float64((outputHeight-1)*o.Strides[0]+o.KernelShape[0]-input[0].Shape()[2]), float64(0))) / 2
				o.Pads[1] = int(math.Max(float64((outputWidth-1)*o.Strides[1]+o.KernelShape[1]-input[0].Shape()[3]), float64(0))) / 2
			*/
		case "SAME_LOWER":
			return nil, &ErrNotImplemented{
				Operator:       "Maxpool",
				AttributeName:  "auto_pad",
				AttributeValue: m.AutoPad,
				Message:        "Padding is buggy",
			}
		default:
			return nil, &ErrNotImplemented{
				Operator:       "Maxpool",
				AttributeName:  "auto_pad",
				AttributeValue: m.AutoPad,
				Message:        "Invalide value",
			}

		}

		if len(m.Pads) == 4 && (m.Pads[0] != m.Pads[1] || m.Pads[2] != m.Pads[3]) {
			return nil, &ErrNotImplemented{
				Operator:       "Maxpool",
				AttributeName:  "pads",
				AttributeValue: m.Pads,
				Message:        "Asymetric padding",
			}
		}
		m.pads = make([]int, 2)
		if len(m.Pads) == 4 {
			for i := 0; i < 2; i++ {
				m.pads[i] = int(m.Pads[2*i])
			}
		}
		if m.pads[0] != 0 || m.pads[1] != 0 {
			return nil, &ErrNotImplemented{
				Operator:       "Maxpool",
				AttributeName:  "pads",
				AttributeValue: m.pads,
				Message:        "Padding is buggy",
			}

		}
		if len(m.KernelShape) != 2 {
			return nil, &ErrNotImplemented{
				Operator: "Maxpool",
				Message:  "Not implemented for dimension != 2",
			}
		}
		m.strides = make([]int, len(m.Strides))
		for i, v := range m.Strides {
			m.strides[i] = int(v)
		}
		m.kernelShape = make([]int, len(m.KernelShape))
		for i, v := range m.KernelShape {
			m.kernelShape[i] = int(v)
		}

		return engine.NewMaxPool2DOperation(m.kernelShape, m.pads, m.strides)(g, n.(*engine.Node))
	}
}
