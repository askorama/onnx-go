package onnx_test

import (
	"github.com/gogo/protobuf/proto"
	"github.com/owulveryck/onnx-go/internal/onnx/ir"
)

var (
	x0       = "x0"
	w0       = "w0"
	x1       = "x1"
	w1       = "w1"
	w2       = "w2"
	y        = "y"
	exp      = "Exp"
	mul      = "MatMul"
	add      = "Add"
	minusOne = "-1"
	one      = "1"
	pow      = "Pow"
	float    = ir.TensorProto_DataType(ir.TensorProto_FLOAT)
	// sigmoidNeuron
	//                   1
	//  y = --------------------------
	//          -(w0.x0 + w1.x1 + w2)
	//     1+ e
	// see http://cs231n.github.io/optimization-2/ for a representation
	sigmoidNeuron = &ir.GraphProto{
		Node: []*ir.NodeProto{
			{
				Input:  []string{x0, w0},
				Output: []string{"x0w0"},
				OpType: mul,
			},
			{
				Input:  []string{x1, w1},
				Output: []string{"x1w1"},
				OpType: mul,
			},
			{
				Input:  []string{"x0w0", "x1w1"},
				Output: []string{"x0w0+x1w1"},
				OpType: add,
			},
			{
				Input:  []string{"x0w0+x1w1", w2},
				Output: []string{"x0w0+x1w1+w2"},
				OpType: add,
			},
			{
				Input:  []string{"x0w0+x1w1+w2", minusOne},
				Output: []string{"-(x0w0+x1w1+w2)"},
				OpType: mul,
			},
			{
				Input:  []string{"-(x0w0+x1w1+w2)"},
				Output: []string{"exp(-(x0w0+x1w1+w2))"},
				OpType: exp,
			},
			{
				Input:  []string{one, "exp(-(x0w0+x1w1+w2))"},
				Output: []string{"1+exp(-(x0w0+x1w1+w2))"},
				OpType: add,
			},
			{
				Input:  []string{"1+exp(-(x0w0+x1w1+w2))", minusOne},
				Output: []string{y},
				OpType: pow,
			},
		},
		Initializer: []*ir.TensorProto{},
		Input: []*ir.ValueInfoProto{
			newValueProtoScalar(minusOne),
			newValueProtoScalar(one),
			newValueProtoScalar(x0),
			newValueProtoScalar(w0),
			newValueProtoScalar(x1),
			newValueProtoScalar(w1),
			newValueProtoScalar(w2),
		},
		Output: []*ir.ValueInfoProto{
			newValueProtoScalar("x0w0"),
			newValueProtoScalar("x1w1"),
			newValueProtoScalar("x0w0+x1w1"),
			newValueProtoScalar("x0w0+x1w1+w2"),
			newValueProtoScalar("-(x0w0+x1w1+w2)"),
			newValueProtoScalar("exp(-(x0w0+x1w1+w2))"),
			newValueProtoScalar("1+exp(-(x0w0+x1w1+w2))"),
			newValueProtoScalar(y),
		},
	}
	sigmoidNeuronONNX []byte
)

func newValueProtoScalar(name string) *ir.ValueInfoProto {
	return &ir.ValueInfoProto{
		Name: name,
		Type: &ir.TypeProto{
			Value: &ir.TypeProto_TensorType{
				TensorType: &ir.TypeProto_Tensor{
					ElemType: int32(ir.TensorProto_FLOAT),
					Shape: &ir.TensorShapeProto{
						Dim: []*ir.TensorShapeProto_Dimension{
							{
								Value: &ir.TensorShapeProto_Dimension_DimValue{
									DimValue: int64(1),
								},
							},
						},
					},
				},
			},
		},
	}
}

func init() {
	model := &ir.ModelProto{
		Graph: sigmoidNeuron,
	}
	var err error
	//sigmoidNeuronONNX, err = model.Marshal()
	sigmoidNeuronONNX, err = proto.Marshal(model)
	if err != nil {
		panic(err)
	}
}
