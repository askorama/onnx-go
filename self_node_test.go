package onnx

import (
	"testing"

	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"gonum.org/v1/gonum/graph"
)

func TestDecodeProto_self(t *testing.T) {
	input := &pb.ModelProto{
		IrVersion: 5,
		OpsetImport: []*pb.OperatorSetIdProto{
			&pb.OperatorSetIdProto{
				Domain:  "",
				Version: 7,
			},
		},
		ProducerName:    "tf2onnx",
		ProducerVersion: "1.5.3",
		Domain:          "",
		ModelVersion:    0,
		DocString:       "",
		Graph: &pb.GraphProto{
			Node: []*pb.NodeProto{
				&pb.NodeProto{
					Input: []string{
						"x:0",
						"x:0",
					},
					Output: []string{
						"mul:0",
					},
					Name:      "mul",
					OpType:    "Mul",
					Domain:    "",
					Attribute: nil,
					DocString: "",
				},
			},
			Name:        "tf2onnx",
			Initializer: nil,
			DocString:   "converted from ./model_nowind_test/export/",
			Input: []*pb.ValueInfoProto{
				&pb.ValueInfoProto{
					Name: "x:0",
					Type: &pb.TypeProto{
						Value: &pb.TypeProto_TensorType{
							TensorType: &pb.TypeProto_Tensor{
								ElemType: 1,
								Shape: &pb.TensorShapeProto{
									Dim: []*pb.TensorShapeProto_Dimension{
										&pb.TensorShapeProto_Dimension{
											Value: &pb.TensorShapeProto_Dimension_DimValue{
												DimValue: 1,
											},
											Denotation: "",
										},
									},
								},
							},
						},
						Denotation: "",
					},
					DocString: "",
				},
			},
			Output: []*pb.ValueInfoProto{
				&pb.ValueInfoProto{
					Name: "mul:0",
					Type: &pb.TypeProto{
						Value: &pb.TypeProto_TensorType{
							TensorType: &pb.TypeProto_Tensor{
								ElemType: 1,
								Shape: &pb.TensorShapeProto{
									Dim: []*pb.TensorShapeProto_Dimension{
										&pb.TensorShapeProto_Dimension{
											Value: &pb.TensorShapeProto_Dimension_DimValue{
												DimValue: 1,
											},
											Denotation: "",
										},
									},
								},
							},
						},
						Denotation: "",
					},
					DocString: "",
				},
			},
			ValueInfo: nil,
		},
		MetadataProps: nil,
	}
	backend := newTestBackend()

	m := NewModel(backend)
	err := m.decodeProto(input)
	if err != nil {
		t.Fatal(err)
	}
	edges := backend.WeightedEdges()
	if edges.Len() != 2 {
		t.Fatal("expected 2 weighted edges")
	}
	ee := make([]graph.WeightedEdge, 2)
	for i := 0; edges.Next(); i++ {
		ee[i] = edges.WeightedEdge()
	}
	for i := 0; i < len(ee); i++ {
		if ee[i].From() == ee[i].To() && ee[i].Weight() == self {
			ee = ee[:len(ee)-1]
		}
		if ee[i].From() != ee[i].To() && ee[i].Weight() == 1 {
			ee = ee[:len(ee)-1]
		}
	}

}
