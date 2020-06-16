package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("MeanVarianceNormalization", "TestMvn", NewTestMvn)
}

/*
&ir.ModelProto{
    IrVersion:   4,
    OpsetImport: {
        &ir.OperatorSetIdProto{Domain:"", Version:10},
    },
    ProducerName:    "backend-test",
    ProducerVersion: "",
    Domain:          "",
    ModelVersion:    0,
    DocString:       "",
    Graph:           &ir.GraphProto{
        Node: {
            &ir.NodeProto{
                Input:     {"X"},
                Output:    {"Y"},
                Name:      "",
                OpType:    "MeanVarianceNormalization",
                Domain:    "",
                Attribute: nil,
                DocString: "",
            },
        },
        Name:              "test_mvn",
        Initializer:       nil,
        SparseInitializer: nil,
        DocString:         "",
        Input:             {
            &ir.ValueInfoProto{
                Name: "X",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
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
                Name: "Y",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
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
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestMvn version: 4.
func NewTestMvn() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "MeanVarianceNormalization",
		Title:  "TestMvn",
		ModelB: []byte{0x8, 0x4, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x67, 0xa, 0x21, 0xa, 0x1, 0x58, 0x12, 0x1, 0x59, 0x22, 0x19, 0x4d, 0x65, 0x61, 0x6e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x8, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x76, 0x6e, 0x5a, 0x1b, 0xa, 0x1, 0x58, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x1, 0x62, 0x1b, 0xa, 0x1, 0x59, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x1, 0x42, 0x2, 0x10, 0xa},

		/*

		   &ir.NodeProto{
		     Input:     []string{"X"},
		     Output:    []string{"Y"},
		     Name:      "",
		     OpType:    "MeanVarianceNormalization",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 3, 3, 1),
				tensor.WithBacking([]float32{0.8439683, 0.5665144, 0.05836735, 0.02916367, 0.12964272, 0.5060197, 0.79538304, 0.9411346, 0.9546573, 0.17730942, 0.46192095, 0.26480448, 0.6746842, 0.01665257, 0.62473077, 0.9240844, 0.9722341, 0.11965699, 0.41356155, 0.9129373, 0.59330076, 0.81929934, 0.7862604, 0.11799799, 0.69248444, 0.54119414, 0.07513223}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 3, 3, 1),
				tensor.WithBacking([]float32{1.3546423, 0.33053496, -1.5450814, -1.2106764, -0.8925952, 0.29888135, 0.38083088, 0.81808794, 0.85865635, -1.1060555, -0.05552877, -0.78310335, 0.83281356, -1.250282, 0.67467856, 0.7669372, 0.9113869, -1.6463585, -0.23402764, 1.6092131, 0.42940593, 1.2906139, 1.1860244, -0.92945826, 0.0721334, -0.38174, -1.7799333}),
			),
		},
	}
}
