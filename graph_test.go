package onnx

import (
	"testing"

	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"gonum.org/v1/gonum/graph/simple"
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
	float    = pb.TensorProto_DataType(pb.TensorProto_FLOAT)
	// sigmoidNeuron
	//                   1
	//  y = --------------------------
	//          -(w0.x0 + w1.x1 + w2)
	//     1+ e
	// see http://cs231n.github.io/optimization-2/ for a representation
	sigmoidNeuron = &pb.GraphProto{
		Node: []*pb.NodeProto{
			&pb.NodeProto{
				Input:  []string{x0, w0},
				Output: []string{"x0w0"},
				OpType: &mul,
			},
			&pb.NodeProto{
				Input:  []string{x1, w1},
				Output: []string{"x1w1"},
				OpType: &mul,
			},
			&pb.NodeProto{
				Input:  []string{"x0w0", "x1w1"},
				Output: []string{"x0w0+x1w1"},
				OpType: &add,
			},
			&pb.NodeProto{
				Input:  []string{"x0w0+x1w1", w2},
				Output: []string{"x0w0+x1w1+w2"},
				OpType: &add,
			},
			&pb.NodeProto{
				Input:  []string{"x0w0+x1w1+w2", minusOne},
				Output: []string{"-(x0w0+x1w1+w2)"},
				OpType: &mul,
			},
			&pb.NodeProto{
				Input:  []string{"-(x0w0+x1w1+w2)"},
				Output: []string{"exp(-(x0w0+x1w1+w2))"},
				OpType: &exp,
			},
			&pb.NodeProto{
				Input:  []string{one, "exp(-(x0w0+x1w1+w2))"},
				Output: []string{"1+exp(-(x0w0+x1w1+w2))"},
				OpType: &add,
			},
			&pb.NodeProto{
				Input:  []string{"1+exp(-(x0w0+x1w1+w2))", minusOne},
				Output: []string{y},
				OpType: &pow,
			},
		},
		Initializer: []*pb.TensorProto{},
		Input: []*pb.ValueInfoProto{
			newValueProtoScalar(minusOne),
			newValueProtoScalar(one),
			newValueProtoScalar(x0),
			newValueProtoScalar(w0),
			newValueProtoScalar(x1),
			newValueProtoScalar(w1),
			newValueProtoScalar(w2),
			newValueProtoScalar("x0w0"),
			newValueProtoScalar("x1w1"),
			newValueProtoScalar("x0w0+x1w1"),
			newValueProtoScalar("x0w0+x1w1+w2"),
			newValueProtoScalar("x0w0+x1w1+w2"),
			newValueProtoScalar("-(x0w0+x1w1+w2)"),
			newValueProtoScalar("exp(-(x0w0+x1w1+w2))"),
			newValueProtoScalar("1+exp(-(x0w0+x1w1+w2))"),
		},
		Output: []*pb.ValueInfoProto{
			newValueProtoScalar(y),
		},
	}
	sigmoidNeuronONNX []byte
)

func newValueProtoScalar(name string) *pb.ValueInfoProto {
	return &pb.ValueInfoProto{
		Name: &name,
		Type: &pb.TypeProto{
			Value: &pb.TypeProto_TensorType{
				TensorType: &pb.TypeProto_Tensor{
					ElemType: &float,
					Shape: &pb.TensorShapeProto{
						Dim: []*pb.TensorShapeProto_Dimension{
							&pb.TensorShapeProto_Dimension{
								Value: &pb.TensorShapeProto_Dimension_DimValue{
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
	model := &pb.ModelProto{
		Graph: sigmoidNeuron,
	}
	var err error
	sigmoidNeuronONNX, err = model.Marshal()
	if err != nil {
		panic(err)
	}
}

func TestUnmarshal(t *testing.T) {
	graph := simple.NewDirectedGraph()
	err := Unmarshal(sigmoidNeuronONNX, graph)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(graph)

}
