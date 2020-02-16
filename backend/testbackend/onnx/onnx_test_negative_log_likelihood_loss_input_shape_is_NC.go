package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("NegativeLogLikelihoodLoss", "TestNegativeLogLikelihoodLossInputShapeIsNC", NewTestNegativeLogLikelihoodLossInputShapeIsNC)
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
                Input:     {"input", "target"},
                Output:    {"loss"},
                Name:      "",
                OpType:    "NegativeLogLikelihoodLoss",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:          "reduction",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          3,
                        F:             0,
                        I:             0,
                        S:             {0x6e, 0x6f, 0x6e, 0x65},
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
        Name:              "test_negative_log_likelihood_loss_input_shape_is_NC",
        Initializer:       nil,
        SparseInitializer: nil,
        DocString:         "",
        Input:             {
            &ir.ValueInfoProto{
                Name: "input",
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
                Name: "target",
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
                Name: "loss",
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

// NewTestNegativeLogLikelihoodLossInputShapeIsNC version: 6.
func NewTestNegativeLogLikelihoodLossInputShapeIsNC() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "NegativeLogLikelihoodLoss",
		Title:  "TestNegativeLogLikelihoodLossInputShapeIsNC",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xc0, 0x1, 0xa, 0x46, 0xa, 0x5, 0x69, 0x6e, 0x70, 0x75, 0x74, 0xa, 0x6, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x4, 0x6c, 0x6f, 0x73, 0x73, 0x22, 0x19, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x6b, 0x65, 0x6c, 0x69, 0x68, 0x6f, 0x6f, 0x64, 0x4c, 0x6f, 0x73, 0x73, 0x2a, 0x14, 0xa, 0x9, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4, 0x6e, 0x6f, 0x6e, 0x65, 0xa0, 0x1, 0x3, 0x12, 0x33, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x6c, 0x69, 0x68, 0x6f, 0x6f, 0x64, 0x5f, 0x6c, 0x6f, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x5f, 0x69, 0x73, 0x5f, 0x4e, 0x43, 0x5a, 0x17, 0xa, 0x5, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x5, 0x5a, 0x14, 0xa, 0x6, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0xa, 0xa, 0x8, 0x8, 0x6, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x62, 0x12, 0xa, 0x4, 0x6c, 0x6f, 0x73, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x42, 0x2, 0x10, 0xc},

		/*

		   &ir.NodeProto{
		     Input:     []string{"input", "target"},
		     Output:    []string{"loss"},
		     Name:      "",
		     OpType:    "NegativeLogLikelihoodLoss",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc0001e6700)(name:"reduction" type:STRING s:"none" )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 5),
				tensor.WithBacking([]float32{0.5488135, 0.71518934, 0.60276335, 0.5448832, 0.4236548, 0.6458941, 0.4375872, 0.891773, 0.96366274, 0.3834415, 0.79172504, 0.5288949, 0.56804454, 0.92559665, 0.071036056}),
			),

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]int32{1, 4, 3}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float32{-0.71518934, -0.3834415, -0.92559665}),
			),
		},
	}
}
