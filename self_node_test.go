package onnx

import (
	"testing"

	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
)

func TestDecodeProto_self(t *testing.T) {
	_ = pb.ModelProto{
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
}
