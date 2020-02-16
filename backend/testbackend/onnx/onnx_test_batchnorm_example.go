package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("BatchNormalization", "TestBatchnormExample", NewTestBatchnormExample)
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
                Input:     {"x", "s", "bias", "mean", "var"},
                Output:    {"y"},
                Name:      "",
                OpType:    "BatchNormalization",
                Domain:    "",
                Attribute: nil,
                DocString: "",
            },
        },
        Name:              "test_batchnorm_example",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
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
                Name: "s",
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
                Name: "bias",
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
                Name: "mean",
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
                Name: "var",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
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

// NewTestBatchnormExample version: 3.
func NewTestBatchnormExample() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "BatchNormalization",
		Title:  "TestBatchnormExample",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xce, 0x1, 0xa, 0x2e, 0xa, 0x1, 0x78, 0xa, 0x1, 0x73, 0xa, 0x4, 0x62, 0x69, 0x61, 0x73, 0xa, 0x4, 0x6d, 0x65, 0x61, 0x6e, 0xa, 0x3, 0x76, 0x61, 0x72, 0x12, 0x1, 0x79, 0x22, 0x12, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x6e, 0x6f, 0x72, 0x6d, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5a, 0x1b, 0xa, 0x1, 0x78, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x3, 0x5a, 0xf, 0xa, 0x1, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x5a, 0x12, 0xa, 0x4, 0x62, 0x69, 0x61, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x5a, 0x12, 0xa, 0x4, 0x6d, 0x65, 0x61, 0x6e, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x5a, 0x11, 0xa, 0x3, 0x76, 0x61, 0x72, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x62, 0x1b, 0xa, 0x1, 0x79, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x3, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x", "s", "bias", "mean", "var"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "BatchNormalization",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 2, 1, 3),
				tensor.WithBacking([]float32{-1, 0, 1, 2, 3, 4}),
			),

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{1, 1.5}),
			),

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{0, 1}),
			),

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{0, 3}),
			),

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{1, 1.5}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 2, 1, 3),
				tensor.WithBacking([]float32{-0.999995, 0, 0.999995, -0.22474074, 1, 2.2247407}),
			),
		},
	}
}
