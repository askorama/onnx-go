package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("GatherND", "TestGatherndExampleInt32", NewTestGatherndExampleInt32)
}

/*
&ir.ModelProto{
    IrVersion:   6,
    OpsetImport: {
        &ir.OperatorSetIdProto{Domain:"", Version:11},
    },
    ProducerName:    "backend-test",
    ProducerVersion: "",
    Domain:          "",
    ModelVersion:    0,
    DocString:       "",
    Graph:           &ir.GraphProto{
        Node: {
            &ir.NodeProto{
                Input:     {"data", "indices"},
                Output:    {"output"},
                Name:      "",
                OpType:    "GatherND",
                Domain:    "",
                Attribute: nil,
                DocString: "",
            },
        },
        Name:              "test_gathernd_example_int32",
        Initializer:       nil,
        SparseInitializer: nil,
        DocString:         "",
        Input:             {
            &ir.ValueInfoProto{
                Name: "data",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 6,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
                                        Denotation: "",
                                    },
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
                Name: "indices",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 7,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
                                        Denotation: "",
                                    },
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
        Output: {
            &ir.ValueInfoProto{
                Name: "output",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 6,
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

// NewTestGatherndExampleInt32 version: 6.
func NewTestGatherndExampleInt32() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "GatherND",
		Title:  "TestGatherndExampleInt32",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x89, 0x1, 0xa, 0x21, 0xa, 0x4, 0x64, 0x61, 0x74, 0x61, 0xa, 0x7, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0x6, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x8, 0x47, 0x61, 0x74, 0x68, 0x65, 0x72, 0x4e, 0x44, 0x12, 0x1b, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x67, 0x61, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x64, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5a, 0x16, 0xa, 0x4, 0x64, 0x61, 0x74, 0x61, 0x12, 0xe, 0xa, 0xc, 0x8, 0x6, 0x12, 0x8, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x2, 0x5a, 0x19, 0xa, 0x7, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0xe, 0xa, 0xc, 0x8, 0x7, 0x12, 0x8, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x2, 0x62, 0x14, 0xa, 0x6, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0xa, 0xa, 0x8, 0x8, 0x6, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x42, 0x2, 0x10, 0xb},

		/*

		   &ir.NodeProto{
		     Input:     []string{"data", "indices"},
		     Output:    []string{"output"},
		     Name:      "",
		     OpType:    "GatherND",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 2),
				tensor.WithBacking([]int32{0, 1, 2, 3}),
			),

			tensor.New(
				tensor.WithShape(2, 2),
				tensor.WithBacking([]int64{0, 0, 1, 1}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]int32{0, 3}),
			),
		},
	}
}
