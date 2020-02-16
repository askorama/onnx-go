package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Identity", "TestIdentity", NewTestIdentity)
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
                Input:     {"x"},
                Output:    {"y"},
                Name:      "",
                OpType:    "Identity",
                Domain:    "",
                Attribute: nil,
                DocString: "",
            },
        },
        Name:              "test_identity",
        Initializer:       nil,
        SparseInitializer: nil,
        DocString:         "",
        Input:             {
            &ir.ValueInfoProto{
                Name: "x",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
                                        Denotation: "",
                                    },
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
                Name: "y",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
                                        Denotation: "",
                                    },
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
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestIdentity version: 3.
func NewTestIdentity() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Identity",
		Title:  "TestIdentity",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x5b, 0xa, 0x10, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x8, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0xd, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5a, 0x1b, 0xa, 0x1, 0x78, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x2, 0x62, 0x1b, 0xa, 0x1, 0x79, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x2, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Identity",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 1, 2, 2),
				tensor.WithBacking([]float32{1, 2, 3, 4}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 1, 2, 2),
				tensor.WithBacking([]float32{1, 2, 3, 4}),
			),
		},
	}
}
