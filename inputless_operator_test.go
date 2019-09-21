package onnx

import (
	"testing"

	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
)

func TestDecodeProto_inputless(t *testing.T) {
	proto := &pb.ModelProto{
		IrVersion: 3,
		OpsetImport: []*pb.OperatorSetIdProto{
			{
				Domain:  "",
				Version: 9,
			},
		},
		ProducerName:    "backend-test",
		ProducerVersion: "",
		Domain:          "",
		ModelVersion:    0,
		DocString:       "",
		Graph: &pb.GraphProto{
			Node: []*pb.NodeProto{
				{
					Input: nil,
					Output: []string{
						"values",
					},
					Name:   "",
					OpType: "Constant",
					Domain: "",
					Attribute: []*pb.AttributeProto{
						{
							Name:        "value",
							RefAttrName: "",
							DocString:   "",
							Type:        4,
							F:           0,
							I:           0,
							S:           nil,
							T: &pb.TensorProto{
								Dims: []int64{
									5,
									5,
								},
								DataType: 1,
								Segment:  nil,
								FloatData: []float32{
									1.7640524,
									0.4001572,
									0.978738,
									2.2408931,
									1.867558,
									-0.9772779,
									0.95008844,
									-0.1513572,
									-0.10321885,
									0.41059852,
									0.14404356,
									1.4542735,
									0.7610377,
									0.121675014,
									0.44386324,
									0.33367434,
									1.4940791,
									-0.20515826,
									0.3130677,
									-0.85409576,
									-2.5529897,
									0.6536186,
									0.8644362,
									-0.742165,
									2.2697546,
								},
								Int32Data:  nil,
								StringData: nil,
								Int64Data:  nil,
								Name:       "const_tensor",
								DocString:  "",
								RawData:    nil,
								DoubleData: nil,
								Uint64Data: nil,
							},
							G:       nil,
							Floats:  nil,
							Ints:    nil,
							Strings: nil,
							Tensors: nil,
							Graphs:  nil,
						},
					},
					DocString: "",
				},
			},
			Name:        "test_constant",
			Initializer: nil,
			DocString:   "",
			Input:       nil,
			Output: []*pb.ValueInfoProto{
				{
					Name: "values",
					Type: &pb.TypeProto{
						Value: &pb.TypeProto_TensorType{
							TensorType: &pb.TypeProto_Tensor{
								ElemType: 1,
								Shape: &pb.TensorShapeProto{
									Dim: []*pb.TensorShapeProto_Dimension{
										{
											Value: &pb.TensorShapeProto_Dimension_DimValue{
												DimValue: 5,
											},
											Denotation: "",
										},
										{
											Value: &pb.TensorShapeProto_Dimension_DimValue{
												DimValue: 5,
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
	err := m.decodeProto(proto)
	if err != nil {
		t.Fatal(err)
	}
	if backend.Nodes().Len() != 2 {
		t.Fail()
	}
	if backend.Edges().Len() != 1 {
		t.Fail()

	}
}
