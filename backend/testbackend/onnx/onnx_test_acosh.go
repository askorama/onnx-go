package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Acosh", "TestAcosh", NewTestAcosh)
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
                OpType:    "Acosh",
                Domain:    "",
                Attribute: nil,
                DocString: "",
            },
        },
        Name:        "test_acosh",
        Initializer: nil,
        DocString:   "",
        Input:       {
            &ir.ValueInfoProto{
                Name: "x",
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
        },
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestAcosh version: 3.
func NewTestAcosh() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Acosh",
		Title:  "TestAcosh",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x4d, 0xa, 0xd, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x5, 0x41, 0x63, 0x6f, 0x73, 0x68, 0x12, 0xa, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x6f, 0x73, 0x68, 0x5a, 0x17, 0xa, 0x1, 0x78, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x62, 0x17, 0xa, 0x1, 0x79, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Acosh",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{5.9393215, 7.436704, 6.4248705, 5.903949, 4.8128934, 6.813047, 4.938285, 9.025957, 9.672965, 4.4509735, 8.125525, 5.760054, 6.112401, 9.33037, 1.6393245, 1.7841637, 1.1819656, 8.493579, 8.00341, 8.83011, 9.807565, 8.192427, 5.153314, 8.024762, 2.0644698, 6.7592893, 2.2901795, 9.50202, 5.696635, 4.7319574, 3.3810005, 7.9681034, 5.105353, 6.115906, 1.1691082, 6.5587196, 6.5088615, 6.552406, 9.493732, 7.1363826, 4.235571, 4.9332876, 7.278681, 1.5420293, 7.0009003, 7.035741, 2.893443, 2.1603367, 3.8388553, 4.273397, 6.131771, 4.9474134, 9.895365, 1.9184033, 2.879891, 2.4517856, 6.877975, 3.2796245, 5.196797, 3.1998303}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{2.4675784, 2.6950235, 2.5472116, 2.4615178, 2.253474, 2.6065567, 2.2797523, 2.890169, 2.9597995, 2.1734052, 2.7843494, 2.4364724, 2.4967072, 2.9235377, 1.0778373, 1.1822617, 0.59447473, 2.828974, 2.769089, 2.8680928, 2.973692, 2.7926116, 2.3232377, 2.7717743, 1.3534045, 2.5985475, 1.4702908, 2.9418712, 2.4252284, 2.2361293, 1.8886944, 2.7646327, 2.3137043, 2.4972882, 0.5736651, 2.5680795, 2.5603578, 2.567105, 2.9409935, 2.6534078, 2.1224294, 2.2787185, 2.6733441, 0.99910486, 2.6340458, 2.6390612, 1.7242991, 1.4049425, 2.0209084, 2.1315765, 2.4999144, 2.2816381, 2.9826505, 1.2685117, 1.7192944, 1.5455111, 2.6161444, 1.8567783, 2.3318014, 1.8308822}),
			),
		},
	}
}
