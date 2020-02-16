package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("ReduceLogSum", "TestReduceLogSumAscAxes", NewTestReduceLogSumAscAxes)
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
                Input:     {"data"},
                Output:    {"reduced"},
                Name:      "",
                OpType:    "ReduceLogSum",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:        "axes",
                        RefAttrName: "",
                        DocString:   "",
                        Type:        7,
                        F:           0,
                        I:           0,
                        S:           nil,
                        T:           (*ir.TensorProto)(nil),
                        G:           (*ir.GraphProto)(nil),
                        Floats:      nil,
                        Ints:        {0, 1},
                        Strings:     nil,
                        Tensors:     nil,
                        Graphs:      nil,
                    },
                    &ir.AttributeProto{
                        Name:        "keepdims",
                        RefAttrName: "",
                        DocString:   "",
                        Type:        2,
                        F:           0,
                        I:           0,
                        S:           nil,
                        T:           (*ir.TensorProto)(nil),
                        G:           (*ir.GraphProto)(nil),
                        Floats:      nil,
                        Ints:        nil,
                        Strings:     nil,
                        Tensors:     nil,
                        Graphs:      nil,
                    },
                },
                DocString: "",
            },
        },
        Name:        "test_reduce_log_sum_asc_axes",
        Initializer: nil,
        DocString:   "",
        Input:       {
            &ir.ValueInfoProto{
                Name: "data",
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
                Name: "reduced",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
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
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestReduceLogSumAscAxes version: 3.
func NewTestReduceLogSumAscAxes() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "ReduceLogSum",
		Title:  "TestReduceLogSumAscAxes",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x90, 0x1, 0xa, 0x3d, 0xa, 0x4, 0x64, 0x61, 0x74, 0x61, 0x12, 0x7, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x64, 0x22, 0xc, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x53, 0x75, 0x6d, 0x2a, 0xd, 0xa, 0x4, 0x61, 0x78, 0x65, 0x73, 0x40, 0x0, 0x40, 0x1, 0xa0, 0x1, 0x7, 0x2a, 0xf, 0xa, 0x8, 0x6b, 0x65, 0x65, 0x70, 0x64, 0x69, 0x6d, 0x73, 0x18, 0x0, 0xa0, 0x1, 0x2, 0x12, 0x1c, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x75, 0x6d, 0x5f, 0x61, 0x73, 0x63, 0x5f, 0x61, 0x78, 0x65, 0x73, 0x5a, 0x1a, 0xa, 0x4, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x62, 0x15, 0xa, 0x7, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x64, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"data"},
		     Output:    []string{"reduced"},
		     Name:      "",
		     OpType:    "ReduceLogSum",
		     Attributes: ([]*ir.AttributeProto) (len=2 cap=2) {
		    (*ir.AttributeProto)(0xc000176620)(name:"axes" type:INTS ints:0 ints:1 ),
		    (*ir.AttributeProto)(0xc000176700)(name:"keepdims" type:INT )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{0.15896958, 0.11037514, 0.6563296, 0.13818295, 0.19658236, 0.36872518, 0.82099324, 0.09710128, 0.8379449, 0.09609841, 0.97645944, 0.4686512, 0.9767611, 0.6048455, 0.7392636, 0.039187793, 0.28280696, 0.12019656, 0.2961402, 0.11872772, 0.31798318, 0.41426298, 0.064147495, 0.6924721, 0.56660146, 0.2653895, 0.5232481, 0.09394051, 0.5759465, 0.9292962, 0.31856894, 0.6674104, 0.13179787, 0.7163272, 0.2894061, 0.18319136, 0.5865129, 0.020107547, 0.82894003, 0.004695476, 0.6778165, 0.27000797, 0.735194, 0.96218854, 0.24875315, 0.57615733, 0.5920419, 0.5722519, 0.22308163, 0.952749, 0.44712538, 0.84640867, 0.6994793, 0.29743695, 0.81379783, 0.39650574, 0.8811032, 0.5812729, 0.8817354, 0.6925316}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(5),
				tensor.WithBacking([]float32{1.553096, 1.866221, 1.5578456, 1.9537709, 1.7313905}),
			),
		},
	}
}
