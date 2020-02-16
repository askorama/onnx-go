package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Einsum", "TestEinsumInnerProd", NewTestEinsumInnerProd)
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
                Input:     {"x", "y"},
                Output:    {"z"},
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
                        S:             {0x69, 0x2c, 0x69},
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
        Name:              "test_einsum_inner_prod",
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
                Name: "y",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 11,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
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
                Name: "z",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 11,
                            Shape:    &ir.TensorShapeProto{},
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

// NewTestEinsumInnerProd version: 6.
func NewTestEinsumInnerProd() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Einsum",
		Title:  "TestEinsumInnerProd",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x6e, 0xa, 0x25, 0xa, 0x1, 0x78, 0xa, 0x1, 0x79, 0x12, 0x1, 0x7a, 0x22, 0x6, 0x45, 0x69, 0x6e, 0x73, 0x75, 0x6d, 0x2a, 0x12, 0xa, 0x8, 0x65, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3, 0x69, 0x2c, 0x69, 0xa0, 0x1, 0x3, 0x12, 0x16, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x69, 0x6e, 0x73, 0x75, 0x6d, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x5a, 0xf, 0xa, 0x1, 0x78, 0x12, 0xa, 0xa, 0x8, 0x8, 0xb, 0x12, 0x4, 0xa, 0x2, 0x8, 0x5, 0x5a, 0xf, 0xa, 0x1, 0x79, 0x12, 0xa, 0xa, 0x8, 0x8, 0xb, 0x12, 0x4, 0xa, 0x2, 0x8, 0x5, 0x62, 0xb, 0xa, 0x1, 0x7a, 0x12, 0x6, 0xa, 0x4, 0x8, 0xb, 0x12, 0x0, 0x42, 0x2, 0x10, 0xc},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x", "y"},
		     Output:    []string{"z"},
		     Name:      "",
		     OpType:    "Einsum",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc000217100)(name:"equation" type:STRING s:"i,i" )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(5),
				tensor.WithBacking([]float64{1.764052345967664, 0.4001572083672233, 0.9787379841057392, 2.240893199201458, 1.8675579901499675}),
			),

			tensor.New(
				tensor.WithShape(5),
				tensor.WithBacking([]float64{-0.977277879876411, 0.9500884175255894, -0.1513572082976979, -0.10321885179355784, 0.41059850193837233}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1),
				tensor.WithBacking([]float32{-0.9564095667033291}),
			),
		},
	}
}
