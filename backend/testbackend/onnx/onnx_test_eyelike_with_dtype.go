package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("EyeLike", "TestEyelikeWithDtype", NewTestEyelikeWithDtype)
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
                OpType:    "EyeLike",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:          "dtype",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          2,
                        F:             0,
                        I:             11,
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
        Name:              "test_eyelike_with_dtype",
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

// NewTestEyelikeWithDtype version: 3.
func NewTestEyelikeWithDtype() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "EyeLike",
		Title:  "TestEyelikeWithDtype",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x62, 0xa, 0x1d, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x7, 0x45, 0x79, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x2a, 0xc, 0xa, 0x5, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0xb, 0xa0, 0x1, 0x2, 0x12, 0x17, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x79, 0x65, 0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x74, 0x79, 0x70, 0x65, 0x5a, 0x13, 0xa, 0x1, 0x78, 0x12, 0xe, 0xa, 0xc, 0x8, 0x6, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0x62, 0x13, 0xa, 0x1, 0x79, 0x12, 0xe, 0xa, 0xc, 0x8, 0xb, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "EyeLike",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc000216f00)(name:"dtype" type:INT i:11 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4),
				tensor.WithBacking([]int32{44, 47, 64, 67, 67, 9, 83, 21, 36, 87, 70, 88}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4),
				tensor.WithBacking([]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0}),
			),
		},
	}
}
