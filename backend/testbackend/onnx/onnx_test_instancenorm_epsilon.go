package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("InstanceNormalization", "TestInstancenormEpsilon", NewTestInstancenormEpsilon)
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
                Input:     {"x", "s", "bias"},
                Output:    {"y"},
                Name:      "",
                OpType:    "InstanceNormalization",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:          "epsilon",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          1,
                        F:             0.009999999776482582,
                        I:             0,
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
        Name:              "test_instancenorm_epsilon",
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
                Name: "s",
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
            &ir.ValueInfoProto{
                Name: "bias",
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
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestInstancenormEpsilon version: 3.
func NewTestInstancenormEpsilon() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "InstanceNormalization",
		Title:  "TestInstancenormEpsilon",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xb5, 0x1, 0xa, 0x39, 0xa, 0x1, 0x78, 0xa, 0x1, 0x73, 0xa, 0x4, 0x62, 0x69, 0x61, 0x73, 0x12, 0x1, 0x79, 0x22, 0x15, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x11, 0xa, 0x7, 0x65, 0x70, 0x73, 0x69, 0x6c, 0x6f, 0x6e, 0x15, 0xa, 0xd7, 0x23, 0x3c, 0xa0, 0x1, 0x1, 0x12, 0x19, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x6e, 0x6f, 0x72, 0x6d, 0x5f, 0x65, 0x70, 0x73, 0x69, 0x6c, 0x6f, 0x6e, 0x5a, 0x1b, 0xa, 0x1, 0x78, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x5a, 0xf, 0xa, 0x1, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x5a, 0x12, 0xa, 0x4, 0x62, 0x69, 0x61, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x62, 0x1b, 0xa, 0x1, 0x79, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x", "s", "bias"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "InstanceNormalization",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc000217500)(name:"epsilon" type:FLOAT f:0.01 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 3, 4, 5),
				tensor.WithBacking([]float32{1.7640524, 0.4001572, 0.978738, 2.2408931, 1.867558, -0.9772779, 0.95008844, -0.1513572, -0.10321885, 0.41059852, 0.14404356, 1.4542735, 0.7610377, 0.121675014, 0.44386324, 0.33367434, 1.4940791, -0.20515826, 0.3130677, -0.85409576, -2.5529897, 0.6536186, 0.8644362, -0.742165, 2.2697546, -1.4543657, 0.045758516, -0.18718386, 1.5327792, 1.4693588, 0.15494743, 0.37816253, -0.88778573, -1.9807965, -0.34791216, 0.15634897, 1.2302907, 1.2023798, -0.3873268, -0.30230275, -1.048553, -1.420018, -1.7062702, 1.9507754, -0.5096522, -0.4380743, -1.2527953, 0.7774904, -1.6138978, -0.21274029, -0.89546657, 0.3869025, -0.51080513, -1.1806322, -0.028182229, 0.42833188, 0.06651722, 0.3024719, -0.6343221, -0.36274117, -0.67246044, -0.35955316, -0.8131463, -1.7262826, 0.17742614, -0.40178093, -1.6301984, 0.46278226, -0.9072984, 0.051945396, 0.7290906, 0.12898292, 1.1394007, -1.2348258, 0.40234163, -0.6848101, -0.87079716, -0.5788497, -0.31155252, 0.05616534, -1.1651498, 0.9008265, 0.46566245, -1.5362437, 1.4882522, 1.8958892, 1.1787796, -0.17992483, -1.0707526, 1.0544517, -0.40317693, 1.222445, 0.20827498, 0.97663903, 0.3563664, 0.7065732, 0.01050002, 1.7858706, 0.12691209, 0.40198937, 1.8831507, -1.347759, -1.270485, 0.9693967, -1.1731234, 1.9436212, -0.41361898, -0.7474548, 1.922942, 1.4805148, 1.867559, 0.90604466, -0.86122566, 1.9100649, -0.26800337, 0.8024564, 0.947252, -0.15501009, 0.61407936, 0.9222067}),
			),

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float32{0.37642553, -1.0994008, 0.2982382}),
			),

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float32{1.3263859, -0.69456786, -0.14963454}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 3, 4, 5),
				tensor.WithBacking([]float32{1.8517356, 1.251994, 1.5064116, 2.0614152, 1.8972497, 0.6462986, 1.4938136, 1.0094782, 1.0306458, 1.2565854, 1.1393739, 1.7155174, 1.4106829, 1.1295378, 1.2712127, 1.2227597, 1.733021, 0.9858204, 1.2136984, 0.7004651, 1.7114694, -1.245981, -1.4404176, 0.041348398, -2.736541, 0.69821006, -0.6853524, -0.47050995, -2.0568295, -1.998337, -0.7860572, -0.99192816, 0.17565417, 1.1837363, -0.32227048, -0.7873498, -1.7778447, -1.7521026, -0.28591838, -0.36433598, -0.37025386, -0.49566495, -0.59230715, 0.64235556, -0.18831447, -0.16414891, -0.43920854, 0.2462404, -0.5611211, -0.088073425, -0.31857005, 0.11437321, -0.18870372, -0.41484544, -0.025764398, 0.12836027, 0.0062072724, 0.08586843, -0.23040454, -0.13871554, 1.1640255, 1.3226438, 1.0927093, 0.6298244, 1.594848, 1.3012377, 0.67853117, 1.7395002, 1.044982, 1.5312396, 1.8744965, 1.5702913, 2.08249, 0.8789525, 1.7088617, 1.1577653, 1.0634851, 1.2114785, 1.3469762, 1.5333788, 1.1680491, -1.2577085, -0.74676245, 1.6037674, -1.947432, -2.426057, -1.5840659, 0.011251152, 1.0572131, -1.4380869, 0.27338165, -1.6353356, -0.44455203, -1.3467236, -0.61843294, -1.0296268, -0.21233538, -2.2968793, -0.34902012, -0.672001, 0.20946464, -0.6273186, -0.60730517, -0.02719134, -0.5820892, 0.22512603, -0.38538283, -0.47184405, 0.21977031, 0.105184704, 0.20542648, -0.043599084, -0.5013099, 0.21643522, -0.3476694, -0.07042773, -0.03292667, -0.3184049, -0.119216084, -0.03941323}),
			),
		},
	}
}
