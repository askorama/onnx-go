package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("DepthToSpace", "TestDepthtospaceDcrMode", NewTestDepthtospaceDcrMode)
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
                Input:     {"x"},
                Output:    {"y"},
                Name:      "",
                OpType:    "DepthToSpace",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:          "blocksize",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          2,
                        F:             0,
                        I:             2,
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
                        Name:          "mode",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          3,
                        F:             0,
                        I:             0,
                        S:             {0x44, 0x43, 0x52},
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
        Name:              "test_depthtospace_dcr_mode",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:8},
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
                Name: "y",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:6},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:6},
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

// NewTestDepthtospaceDcrMode version: 5.
func NewTestDepthtospaceDcrMode() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "DepthToSpace",
		Title:  "TestDepthtospaceDcrMode",
		ModelB: []byte{0x8, 0x5, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x8e, 0x1, 0xa, 0x36, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0xc, 0x44, 0x65, 0x70, 0x74, 0x68, 0x54, 0x6f, 0x53, 0x70, 0x61, 0x63, 0x65, 0x2a, 0x10, 0xa, 0x9, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x2, 0xa0, 0x1, 0x2, 0x2a, 0xe, 0xa, 0x4, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x3, 0x44, 0x43, 0x52, 0xa0, 0x1, 0x3, 0x12, 0x1a, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x68, 0x74, 0x6f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x64, 0x63, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5a, 0x1b, 0xa, 0x1, 0x78, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x3, 0x62, 0x1b, 0xa, 0x1, 0x79, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x6, 0xa, 0x2, 0x8, 0x6, 0x42, 0x2, 0x10, 0xb},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "DepthToSpace",
		     Attributes: ([]*ir.AttributeProto) (len=2 cap=2) {
		    (*ir.AttributeProto)(0xc000216c00)(name:"blocksize" type:INT i:2 ),
		    (*ir.AttributeProto)(0xc000216d00)(name:"mode" type:STRING s:"DCR" )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 8, 3, 3),
				tensor.WithBacking([]float32{0.5488135, 0.71518934, 0.60276335, 0.5448832, 0.4236548, 0.6458941, 0.4375872, 0.891773, 0.96366274, 0.3834415, 0.79172504, 0.5288949, 0.56804454, 0.92559665, 0.071036056, 0.0871293, 0.020218397, 0.83261985, 0.77815676, 0.87001216, 0.9786183, 0.7991586, 0.46147937, 0.7805292, 0.11827443, 0.639921, 0.14335328, 0.9446689, 0.5218483, 0.41466194, 0.2645556, 0.7742337, 0.45615032, 0.56843394, 0.0187898, 0.6176355, 0.6120957, 0.616934, 0.94374806, 0.6818203, 0.3595079, 0.43703195, 0.6976312, 0.06022547, 0.6667667, 0.67063785, 0.21038257, 0.12892629, 0.31542835, 0.36371076, 0.57019675, 0.43860152, 0.9883738, 0.10204481, 0.20887676, 0.16130951, 0.6531083, 0.2532916, 0.46631077, 0.2444256, 0.15896958, 0.11037514, 0.6563296, 0.13818295, 0.19658236, 0.36872518, 0.82099324, 0.09710128, 0.8379449, 0.09609841, 0.97645944, 0.4686512, 0.9767611, 0.6048455, 0.7392636, 0.039187793, 0.28280696, 0.12019656, 0.2961402, 0.11872772, 0.31798318, 0.41426298, 0.064147495, 0.6924721, 0.56660146, 0.2653895, 0.5232481, 0.09394051, 0.5759465, 0.9292962, 0.31856894, 0.6674104, 0.13179787, 0.7163272, 0.2894061, 0.18319136, 0.5865129, 0.020107547, 0.82894003, 0.004695476, 0.6778165, 0.27000797, 0.735194, 0.96218854, 0.24875315, 0.57615733, 0.5920419, 0.5722519, 0.22308163, 0.952749, 0.44712538, 0.84640867, 0.6994793, 0.29743695, 0.81379783, 0.39650574, 0.8811032, 0.5812729, 0.8817354, 0.6925316, 0.7252543, 0.50132436, 0.95608366, 0.6439902, 0.42385504, 0.6063932, 0.019193199, 0.30157483, 0.66017354, 0.2900776, 0.6180154, 0.4287687, 0.13547407, 0.29828233, 0.5699649, 0.59087276, 0.57432526, 0.6532008, 0.65210325, 0.43141845, 0.8965466, 0.36756188, 0.43586493, 0.89192337}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 2, 6, 6),
				tensor.WithBacking([]float32{0.5488135, 0.77815676, 0.71518934, 0.87001216, 0.60276335, 0.9786183, 0.6120957, 0.20887676, 0.616934, 0.16130951, 0.94374806, 0.6531083, 0.5448832, 0.7991586, 0.4236548, 0.46147937, 0.6458941, 0.7805292, 0.6818203, 0.2532916, 0.3595079, 0.46631077, 0.43703195, 0.2444256, 0.4375872, 0.11827443, 0.891773, 0.639921, 0.96366274, 0.14335328, 0.6976312, 0.15896958, 0.06022547, 0.11037514, 0.6667667, 0.6563296, 0.3834415, 0.9446689, 0.79172504, 0.5218483, 0.5288949, 0.41466194, 0.67063785, 0.13818295, 0.21038257, 0.19658236, 0.12892629, 0.36872518, 0.56804454, 0.2645556, 0.92559665, 0.7742337, 0.071036056, 0.45615032, 0.31542835, 0.82099324, 0.36371076, 0.09710128, 0.57019675, 0.8379449, 0.0871293, 0.56843394, 0.020218397, 0.0187898, 0.83261985, 0.6176355, 0.43860152, 0.09609841, 0.9883738, 0.97645944, 0.10204481, 0.4686512, 0.9767611, 0.31856894, 0.6048455, 0.6674104, 0.7392636, 0.13179787, 0.22308163, 0.019193199, 0.952749, 0.30157483, 0.44712538, 0.66017354, 0.039187793, 0.7163272, 0.28280696, 0.2894061, 0.12019656, 0.18319136, 0.84640867, 0.2900776, 0.6994793, 0.6180154, 0.29743695, 0.4287687, 0.2961402, 0.5865129, 0.11872772, 0.020107547, 0.31798318, 0.82894003, 0.81379783, 0.13547407, 0.39650574, 0.29828233, 0.8811032, 0.5699649, 0.41426298, 0.004695476, 0.064147495, 0.6778165, 0.6924721, 0.27000797, 0.5812729, 0.59087276, 0.8817354, 0.57432526, 0.6925316, 0.6532008, 0.56660146, 0.735194, 0.2653895, 0.96218854, 0.5232481, 0.24875315, 0.7252543, 0.65210325, 0.50132436, 0.43141845, 0.95608366, 0.8965466, 0.09394051, 0.57615733, 0.5759465, 0.5920419, 0.9292962, 0.5722519, 0.6439902, 0.36756188, 0.42385504, 0.43586493, 0.6063932, 0.89192337}),
			),
		},
	}
}
