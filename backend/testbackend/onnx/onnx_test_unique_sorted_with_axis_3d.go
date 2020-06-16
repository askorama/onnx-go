package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Unique", "TestUniqueSortedWithAxis3d", NewTestUniqueSortedWithAxis3d)
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
                Input:     {"X"},
                Output:    {"Y", "indices", "inverse_indices", "counts"},
                Name:      "",
                OpType:    "Unique",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:          "axis",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          2,
                        F:             0,
                        I:             1,
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
                    &ir.AttributeProto{
                        Name:          "sorted",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          2,
                        F:             0,
                        I:             1,
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
        Name:              "test_unique_sorted_with_axis_3d",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:4},
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
                Name: "Y",
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
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
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
                                },
                            },
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
            &ir.ValueInfoProto{
                Name: "inverse_indices",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 7,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
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
                Name: "counts",
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

// NewTestUniqueSortedWithAxis3d version: 5.
func NewTestUniqueSortedWithAxis3d() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Unique",
		Title:  "TestUniqueSortedWithAxis3d",
		ModelB: []byte{0x8, 0x5, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xed, 0x1, 0xa, 0x4c, 0xa, 0x1, 0x58, 0x12, 0x1, 0x59, 0x12, 0x7, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0xf, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0x6, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x6, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x2a, 0xb, 0xa, 0x4, 0x61, 0x78, 0x69, 0x73, 0x18, 0x1, 0xa0, 0x1, 0x2, 0x2a, 0xd, 0xa, 0x6, 0x73, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x1, 0xa0, 0x1, 0x2, 0x12, 0x1f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x73, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x5f, 0x33, 0x64, 0x5a, 0x17, 0xa, 0x1, 0x58, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x2, 0x62, 0x17, 0xa, 0x1, 0x59, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x2, 0x62, 0x15, 0xa, 0x7, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x62, 0x1d, 0xa, 0xf, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x4, 0x62, 0x14, 0xa, 0x6, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x42, 0x2, 0x10, 0xb},

		/*

		   &ir.NodeProto{
		     Input:     []string{"X"},
		     Output:    []string{"Y", "indices", "inverse_indices", "counts"},
		     Name:      "",
		     OpType:    "Unique",
		     Attributes: ([]*ir.AttributeProto) (len=2 cap=2) {
		    (*ir.AttributeProto)(0xc0002d8100)(name:"axis" type:INT i:1 ),
		    (*ir.AttributeProto)(0xc0002d8200)(name:"sorted" type:INT i:1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 4, 2),
				tensor.WithBacking([]float32{1, 1, 0, 1, 2, 1, 0, 1, 1, 1, 0, 1, 2, 1, 0, 1}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 3, 2),
				tensor.WithBacking([]float32{0, 1, 1, 1, 2, 1, 0, 1, 1, 1, 2, 1}),
			),

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]int64{1, 0, 2}),
			),

			tensor.New(
				tensor.WithShape(4),
				tensor.WithBacking([]int64{1, 0, 2, 0}),
			),

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]int64{2, 1, 1}),
			),
		},
	}
}
