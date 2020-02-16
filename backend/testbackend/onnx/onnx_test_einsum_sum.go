package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Einsum", "TestEinsumSum", NewTestEinsumSum)
}

/*
&ir.ModelProto{
    IrVersion:   6,
    OpsetImport: {
        &ir.OperatorSetIdProto{Domain:"", Version:12},
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
                OpType:    "Einsum",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:          "equation",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          3,
                        F:             0,
                        I:             0,
                        S:             {0x69, 0x6a, 0x2d, 0x3e, 0x69},
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
        Name:              "test_einsum_sum",
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

// NewTestEinsumSum version: 6.
func NewTestEinsumSum() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Einsum",
		Title:  "TestEinsumSum",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x5d, 0xa, 0x24, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x6, 0x45, 0x69, 0x6e, 0x73, 0x75, 0x6d, 0x2a, 0x14, 0xa, 0x8, 0x65, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5, 0x69, 0x6a, 0x2d, 0x3e, 0x69, 0xa0, 0x1, 0x3, 0x12, 0xf, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x69, 0x6e, 0x73, 0x75, 0x6d, 0x5f, 0x73, 0x75, 0x6d, 0x5a, 0x13, 0xa, 0x1, 0x78, 0x12, 0xe, 0xa, 0xc, 0x8, 0xb, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0x62, 0xf, 0xa, 0x1, 0x79, 0x12, 0xa, 0xa, 0x8, 0x8, 0xb, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x42, 0x2, 0x10, 0xc},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Einsum",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc000217400)(name:"equation" type:STRING s:"ij->i" )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4),
				tensor.WithBacking([]float64{1.764052345967664, 0.4001572083672233, 0.9787379841057392, 2.240893199201458, 1.8675579901499675, -0.977277879876411, 0.9500884175255894, -0.1513572082976979, -0.10321885179355784, 0.41059850193837233, 0.144043571160878, 1.454273506962975}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float64{5.3838407376420845, 1.689011319501448, 1.9056967282686674}),
			),
		},
	}
}
