package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Range", "TestRangeFloatTypePositiveDelta", NewTestRangeFloatTypePositiveDelta)
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
                Input:     {"start", "limit", "delta"},
                Output:    {"output"},
                Name:      "",
                OpType:    "Range",
                Domain:    "",
                Attribute: nil,
                DocString: "",
            },
        },
        Name:              "test_range_float_type_positive_delta",
        Initializer:       nil,
        SparseInitializer: nil,
        DocString:         "",
        Input:             {
            &ir.ValueInfoProto{
                Name: "start",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{},
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
            &ir.ValueInfoProto{
                Name: "limit",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{},
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
            &ir.ValueInfoProto{
                Name: "delta",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{},
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

// NewTestRangeFloatTypePositiveDelta version: 6.
func NewTestRangeFloatTypePositiveDelta() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Range",
		Title:  "TestRangeFloatTypePositiveDelta",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x95, 0x1, 0xa, 0x24, 0xa, 0x5, 0x73, 0x74, 0x61, 0x72, 0x74, 0xa, 0x5, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0xa, 0x5, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x6, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x24, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x5a, 0xf, 0xa, 0x5, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x6, 0xa, 0x4, 0x8, 0x1, 0x12, 0x0, 0x5a, 0xf, 0xa, 0x5, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x6, 0xa, 0x4, 0x8, 0x1, 0x12, 0x0, 0x5a, 0xf, 0xa, 0x5, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x6, 0xa, 0x4, 0x8, 0x1, 0x12, 0x0, 0x62, 0x14, 0xa, 0x6, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x42, 0x2, 0x10, 0xb},

		/*

		   &ir.NodeProto{
		     Input:     []string{"start", "limit", "delta"},
		     Output:    []string{"output"},
		     Name:      "",
		     OpType:    "Range",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1),
				tensor.WithBacking([]float32{1}),
			),

			tensor.New(
				tensor.WithShape(1),
				tensor.WithBacking([]float32{5}),
			),

			tensor.New(
				tensor.WithShape(1),
				tensor.WithBacking([]float32{2}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{1, 3}),
			),
		},
	}
}
