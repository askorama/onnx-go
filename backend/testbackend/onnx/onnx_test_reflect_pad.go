package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Pad", "TestReflectPad", NewTestReflectPad)
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
                Input:     {"x", "pads"},
                Output:    {"y"},
                Name:      "",
                OpType:    "Pad",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:          "mode",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          3,
                        F:             0,
                        I:             0,
                        S:             {0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74},
                        T:             (*ir.TensorProto)(nil),
                        G:             (*ir.GraphProto)(nil),
                        SparseTensor:  (*ir.SparseTensorProto)(nil),
                        Floats:        nil,
                        Ints:          nil,
                        Strings:       nil,
                        Tensors:       nil,
                        Graphs:        nil,
                        SparseTensors: nil,
                    },
                },
                DocString: "",
            },
        },
        Name:              "test_reflect_pad",
        Initializer:       nil,
        SparseInitializer: nil,
        DocString:         "",
        Input:             {
            &ir.ValueInfoProto{
                Name: "x",
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
                Name: "pads",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 7,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:8},
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
                            ElemType: 6,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
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

// NewTestReflectPad version: 6.
func NewTestReflectPad() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Pad",
		Title:  "TestReflectPad",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x87, 0x1, 0xa, 0x25, 0xa, 0x1, 0x78, 0xa, 0x4, 0x70, 0x61, 0x64, 0x73, 0x12, 0x1, 0x79, 0x22, 0x3, 0x50, 0x61, 0x64, 0x2a, 0x12, 0xa, 0x4, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x7, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0xa0, 0x1, 0x3, 0x12, 0x10, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x64, 0x5a, 0x1b, 0xa, 0x1, 0x78, 0x12, 0x16, 0xa, 0x14, 0x8, 0x6, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x5a, 0x12, 0xa, 0x4, 0x70, 0x61, 0x64, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x8, 0x62, 0x1b, 0xa, 0x1, 0x79, 0x12, 0x16, 0xa, 0x14, 0x8, 0x6, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x6, 0xa, 0x2, 0x8, 0x7, 0x42, 0x2, 0x10, 0xb},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x", "pads"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Pad",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc0001e6100)(name:"mode" type:STRING s:"reflect" )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 3, 4, 5),
				tensor.WithBacking([]int32{0, 0, 0, -1, 0, 0, -1, 0, 0, 0, 0, 0, 1, -1, 0, 0, 0, 0, 0, 0, -1, 0, 0, -1, 1, 1, 1, 0, -1, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, -1, -1, 0, -1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0}),
			),

			tensor.New(
				tensor.WithShape(8),
				tensor.WithBacking([]int64{0, 0, 1, 1, 0, 0, 1, 1}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 3, 6, 7),
				tensor.WithBacking([]int32{-1, 0, -1, 0, 0, 0, 0, 0, 0, 0, 0, -1, 0, -1, -1, 0, -1, 0, 0, 0, 0, 0, 0, 0, 1, -1, 0, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, -1, 0, -1, 1, 1, 1, 0, -1, 1, -1, 0, -1, 0, 0, -1, 1, -1, 1, 1, 1, 0, -1, 1, -1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, -1, 1, -1, -1, 0, -1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1}),
			),
		},
	}
}
