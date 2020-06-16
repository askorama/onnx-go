package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("ConvTranspose", "TestConvtranspose3d", NewTestConvtranspose3d)
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
                Input:     {"X", "W"},
                Output:    {"Y"},
                Name:      "",
                OpType:    "ConvTranspose",
                Domain:    "",
                Attribute: nil,
                DocString: "",
            },
        },
        Name:              "test_convtranspose_3d",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
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
            &ir.ValueInfoProto{
                Name: "W",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:5},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:6},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:7},
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

// NewTestConvtranspose3d version: 3.
func NewTestConvtranspose3d() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "ConvTranspose",
		Title:  "TestConvtranspose3d",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x94, 0x1, 0xa, 0x18, 0xa, 0x1, 0x58, 0xa, 0x1, 0x57, 0x12, 0x1, 0x59, 0x22, 0xd, 0x43, 0x6f, 0x6e, 0x76, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x15, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x33, 0x64, 0x5a, 0x1f, 0xa, 0x1, 0x58, 0x12, 0x1a, 0xa, 0x18, 0x8, 0x1, 0x12, 0x14, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x5a, 0x1f, 0xa, 0x1, 0x57, 0x12, 0x1a, 0xa, 0x18, 0x8, 0x1, 0x12, 0x14, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x3, 0x62, 0x1f, 0xa, 0x1, 0x59, 0x12, 0x1a, 0xa, 0x18, 0x8, 0x1, 0x12, 0x14, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x6, 0xa, 0x2, 0x8, 0x7, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"X", "W"},
		     Output:    []string{"Y"},
		     Name:      "",
		     OpType:    "ConvTranspose",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 1, 3, 4, 5),
				tensor.WithBacking([]float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59}),
			),

			tensor.New(
				tensor.WithShape(1, 2, 3, 3, 3),
				tensor.WithBacking([]float32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 2, 5, 6, 7),
				tensor.WithBacking([]float32{0, 1, 3, 6, 9, 7, 4, 5, 12, 21, 27, 33, 24, 13, 15, 33, 54, 63, 72, 51, 27, 30, 63, 99, 108, 117, 81, 42, 25, 52, 81, 87, 93, 64, 33, 15, 31, 48, 51, 54, 37, 19, 20, 42, 66, 72, 78, 54, 28, 50, 104, 162, 174, 186, 128, 66, 90, 186, 288, 306, 324, 222, 114, 120, 246, 378, 396, 414, 282, 144, 90, 184, 282, 294, 306, 208, 106, 50, 102, 156, 162, 168, 114, 58, 60, 123, 189, 198, 207, 141, 72, 135, 276, 423, 441, 459, 312, 159, 225, 459, 702, 729, 756, 513, 261, 270, 549, 837, 864, 891, 603, 306, 195, 396, 603, 621, 639, 432, 219, 105, 213, 324, 333, 342, 231, 117, 60, 122, 186, 192, 198, 134, 68, 130, 264, 402, 414, 426, 288, 146, 210, 426, 648, 666, 684, 462, 234, 240, 486, 738, 756, 774, 522, 264, 170, 344, 522, 534, 546, 368, 186, 90, 182, 276, 282, 288, 194, 98, 40, 81, 123, 126, 129, 87, 44, 85, 172, 261, 267, 273, 184, 93, 135, 273, 414, 423, 432, 291, 147, 150, 303, 459, 468, 477, 321, 162, 105, 212, 321, 327, 333, 224, 113, 55, 111, 168, 171, 174, 117, 59, 0, 1, 3, 6, 9, 7, 4, 5, 12, 21, 27, 33, 24, 13, 15, 33, 54, 63, 72, 51, 27, 30, 63, 99, 108, 117, 81, 42, 25, 52, 81, 87, 93, 64, 33, 15, 31, 48, 51, 54, 37, 19, 20, 42, 66, 72, 78, 54, 28, 50, 104, 162, 174, 186, 128, 66, 90, 186, 288, 306, 324, 222, 114, 120, 246, 378, 396, 414, 282, 144, 90, 184, 282, 294, 306, 208, 106, 50, 102, 156, 162, 168, 114, 58, 60, 123, 189, 198, 207, 141, 72, 135, 276, 423, 441, 459, 312, 159, 225, 459, 702, 729, 756, 513, 261, 270, 549, 837, 864, 891, 603, 306, 195, 396, 603, 621, 639, 432, 219, 105, 213, 324, 333, 342, 231, 117, 60, 122, 186, 192, 198, 134, 68, 130, 264, 402, 414, 426, 288, 146, 210, 426, 648, 666, 684, 462, 234, 240, 486, 738, 756, 774, 522, 264, 170, 344, 522, 534, 546, 368, 186, 90, 182, 276, 282, 288, 194, 98, 40, 81, 123, 126, 129, 87, 44, 85, 172, 261, 267, 273, 184, 93, 135, 273, 414, 423, 432, 291, 147, 150, 303, 459, 468, 477, 321, 162, 105, 212, 321, 327, 333, 224, 113, 55, 111, 168, 171, 174, 117, 59}),
			),
		},
	}
}
