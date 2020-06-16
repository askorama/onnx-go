package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("CumSum", "TestCumsum2dAxis1", NewTestCumsum2dAxis1)
}

/*
&ir.ModelProto{
    IrVersion:   5,
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
                Input:     {"x", "axis"},
                Output:    {"y"},
                Name:      "",
                OpType:    "CumSum",
                Domain:    "",
                Attribute: nil,
                DocString: "",
            },
        },
        Name:              "test_cumsum_2d_axis_1",
        Initializer:       nil,
        SparseInitializer: nil,
        DocString:         "",
        Input:             {
            &ir.ValueInfoProto{
                Name: "x",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 11,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
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
                Name: "axis",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 6,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
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
                            ElemType: 11,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
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

// NewTestCumsum2dAxis1 version: 5.
func NewTestCumsum2dAxis1() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "CumSum",
		Title:  "TestCumsum2dAxis1",
		ModelB: []byte{0x8, 0x5, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x6b, 0xa, 0x14, 0xa, 0x1, 0x78, 0xa, 0x4, 0x61, 0x78, 0x69, 0x73, 0x12, 0x1, 0x79, 0x22, 0x6, 0x43, 0x75, 0x6d, 0x53, 0x75, 0x6d, 0x12, 0x15, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x75, 0x6d, 0x73, 0x75, 0x6d, 0x5f, 0x32, 0x64, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x5f, 0x31, 0x5a, 0x13, 0xa, 0x1, 0x78, 0x12, 0xe, 0xa, 0xc, 0x8, 0xb, 0x12, 0x8, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x3, 0x5a, 0x12, 0xa, 0x4, 0x61, 0x78, 0x69, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x6, 0x12, 0x4, 0xa, 0x2, 0x8, 0x1, 0x62, 0x13, 0xa, 0x1, 0x79, 0x12, 0xe, 0xa, 0xc, 0x8, 0xb, 0x12, 0x8, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x3, 0x42, 0x2, 0x10, 0xb},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x", "axis"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "CumSum",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 3),
				tensor.WithBacking([]float64{1, 2, 3, 4, 5, 6}),
			),

			tensor.New(
				tensor.WithShape(1),
				tensor.WithBacking([]float32{1}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 3),
				tensor.WithBacking([]float64{1, 3, 6, 4, 9, 15}),
			),
		},
	}
}
