package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Split", "TestSplitEqualPartsDefaultAxis", NewTestSplitEqualPartsDefaultAxis)
}

/*
&ir.ModelProto{
    IrVersion:   3,
    OpsetImport: {
        &ir.OperatorSetIdProto{Domain:"", Version:9},
    },
    ProducerName:    "backend-test",
    ProducerVersion: "",
    Domain:          "",
    ModelVersion:    0,
    DocString:       "",
    Graph:           &ir.GraphProto{
        Node: {
            &ir.NodeProto{
                Input:     {"input"},
                Output:    {"output_1", "output_2", "output_3"},
                Name:      "",
                OpType:    "Split",
                Domain:    "",
                Attribute: nil,
                DocString: "",
            },
        },
        Name:              "test_split_equal_parts_default_axis",
        Initializer:       nil,
        SparseInitializer: nil,
        DocString:         "",
        Input:             {
            &ir.ValueInfoProto{
                Name: "input",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:6},
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
        Output: {
            &ir.ValueInfoProto{
                Name: "output_1",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
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
            &ir.ValueInfoProto{
                Name: "output_2",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
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
            &ir.ValueInfoProto{
                Name: "output_3",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
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
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestSplitEqualPartsDefaultAxis version: 3.
func NewTestSplitEqualPartsDefaultAxis() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Split",
		Title:  "TestSplitEqualPartsDefaultAxis",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xb0, 0x1, 0xa, 0x2c, 0xa, 0x5, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x31, 0x12, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x32, 0x12, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x33, 0x22, 0x5, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x12, 0x23, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x5a, 0x13, 0xa, 0x5, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x6, 0x62, 0x16, 0xa, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x31, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x62, 0x16, 0xa, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x32, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x62, 0x16, 0xa, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x33, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"input"},
		     Output:    []string{"output_1", "output_2", "output_3"},
		     Name:      "",
		     OpType:    "Split",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(6),
				tensor.WithBacking([]float32{1, 2, 3, 4, 5, 6}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{1, 2}),
			),

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{3, 4}),
			),

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{5, 6}),
			),
		},
	}
}
