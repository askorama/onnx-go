package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("TopK", "TestTopKNegativeAxis", NewTestTopKNegativeAxis)
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
                Input:     {"x", "k"},
                Output:    {"values", "indices"},
                Name:      "",
                OpType:    "TopK",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:          "axis",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          2,
                        F:             0,
                        I:             -1,
                        S:             nil,
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
        Name:              "test_top_k_negative_axis",
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
            &ir.ValueInfoProto{
                Name: "k",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 7,
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
                Name: "values",
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
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestTopKNegativeAxis version: 6.
func NewTestTopKNegativeAxis() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "TopK",
		Title:  "TestTopKNegativeAxis",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xaa, 0x1, 0xa, 0x33, 0xa, 0x1, 0x78, 0xa, 0x1, 0x6b, 0x12, 0x6, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x7, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x22, 0x4, 0x54, 0x6f, 0x70, 0x4b, 0x2a, 0x14, 0xa, 0x4, 0x61, 0x78, 0x69, 0x73, 0x18, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1, 0xa0, 0x1, 0x2, 0x12, 0x18, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x70, 0x5f, 0x6b, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x5a, 0x13, 0xa, 0x1, 0x78, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0x5a, 0xf, 0xa, 0x1, 0x6b, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x1, 0x62, 0x18, 0xa, 0x6, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x3, 0x62, 0x19, 0xa, 0x7, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0xe, 0xa, 0xc, 0x8, 0x7, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x3, 0x42, 0x2, 0x10, 0xb},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x", "k"},
		     Output:    []string{"values", "indices"},
		     Name:      "",
		     OpType:    "TopK",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc000216f00)(name:"axis" type:INT i:-1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4),
				tensor.WithBacking([]float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}),
			),

			tensor.New(
				tensor.WithShape(1),
				tensor.WithBacking([]float32{3}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 3),
				tensor.WithBacking([]float32{3, 2, 1, 7, 6, 5, 11, 10, 9}),
			),

			tensor.New(
				tensor.WithShape(3, 3),
				tensor.WithBacking([]int64{3, 2, 1, 3, 2, 1, 3, 2, 1}),
			),
		},
	}
}
