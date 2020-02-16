package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Gemm", "TestGemmDefaultZeroBias", NewTestGemmDefaultZeroBias)
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
                Input:     {"a", "b", "c"},
                Output:    {"y"},
                Name:      "",
                OpType:    "Gemm",
                Domain:    "",
                Attribute: nil,
                DocString: "",
            },
        },
        Name:              "test_gemm_default_zero_bias",
        Initializer:       nil,
        SparseInitializer: nil,
        DocString:         "",
        Input:             {
            &ir.ValueInfoProto{
                Name: "a",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:5},
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
                Name: "b",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:5},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:4},
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
                Name: "c",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:4},
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:4},
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

// NewTestGemmDefaultZeroBias version: 6.
func NewTestGemmDefaultZeroBias() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Gemm",
		Title:  "TestGemmDefaultZeroBias",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x85, 0x1, 0xa, 0x12, 0xa, 0x1, 0x61, 0xa, 0x1, 0x62, 0xa, 0x1, 0x63, 0x12, 0x1, 0x79, 0x22, 0x4, 0x47, 0x65, 0x6d, 0x6d, 0x12, 0x1b, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x67, 0x65, 0x6d, 0x6d, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x7a, 0x65, 0x72, 0x6f, 0x5f, 0x62, 0x69, 0x61, 0x73, 0x5a, 0x13, 0xa, 0x1, 0x61, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x5, 0x5a, 0x13, 0xa, 0x1, 0x62, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x4, 0x5a, 0x13, 0xa, 0x1, 0x63, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x4, 0x62, 0x13, 0xa, 0x1, 0x79, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0x42, 0x2, 0x10, 0xb},

		/*

		   &ir.NodeProto{
		     Input:     []string{"a", "b", "c"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Gemm",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 5),
				tensor.WithBacking([]float32{0.5488135, 0.71518934, 0.60276335, 0.5448832, 0.4236548, 0.6458941, 0.4375872, 0.891773, 0.96366274, 0.3834415, 0.79172504, 0.5288949, 0.56804454, 0.92559665, 0.071036056}),
			),

			tensor.New(
				tensor.WithShape(5, 4),
				tensor.WithBacking([]float32{0.0871293, 0.020218397, 0.83261985, 0.77815676, 0.87001216, 0.9786183, 0.7991586, 0.46147937, 0.7805292, 0.11827443, 0.639921, 0.14335328, 0.9446689, 0.5218483, 0.41466194, 0.2645556, 0.7742337, 0.45615032, 0.56843394, 0.0187898}),
			),

			tensor.New(
				tensor.WithShape(1, 4),
				tensor.WithBacking([]float32{0, 0, 0, 0}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4),
				tensor.WithBacking([]float32{1.9832575, 1.2598817, 1.8809859, 0.9956285, 2.3402529, 1.2245564, 2.0757055, 1.0945301, 1.9018837, 1.116203, 1.8695695, 1.1877979}),
			),
		},
	}
}
