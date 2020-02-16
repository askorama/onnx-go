package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("AveragePool", "TestAveragepool1dDefault", NewTestAveragepool1dDefault)
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
                OpType:    "AveragePool",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:          "kernel_shape",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          7,
                        F:             0,
                        I:             0,
                        S:             nil,
                        T:             (*ir.TensorProto)(nil),
                        G:             (*ir.GraphProto)(nil),
                        SparseTensor:  (*ir.SparseTensorProto)(nil),
                        Floats:        nil,
                        Ints:          {2},
                        Strings:       nil,
                        Tensors:       nil,
                        Graphs:        nil,
                        SparseTensors: nil,
                    },
                },
                DocString: "",
            },
        },
        Name:              "test_averagepool_1d_default",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:32},
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:31},
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

// NewTestAveragepool1dDefault version: 3.
func NewTestAveragepool1dDefault() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "AveragePool",
		Title:  "TestAveragepool1dDefault",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x79, 0xa, 0x28, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0xb, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x2a, 0x13, 0xa, 0xc, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x40, 0x2, 0xa0, 0x1, 0x7, 0x12, 0x1b, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x31, 0x64, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5a, 0x17, 0xa, 0x1, 0x78, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x20, 0x62, 0x17, 0xa, 0x1, 0x79, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x1f, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "AveragePool",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc000217000)(name:"kernel_shape" type:INTS ints:2 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 3, 32),
				tensor.WithBacking([]float32{1.7640524, 0.4001572, 0.978738, 2.2408931, 1.867558, -0.9772779, 0.95008844, -0.1513572, -0.10321885, 0.41059852, 0.14404356, 1.4542735, 0.7610377, 0.121675014, 0.44386324, 0.33367434, 1.4940791, -0.20515826, 0.3130677, -0.85409576, -2.5529897, 0.6536186, 0.8644362, -0.742165, 2.2697546, -1.4543657, 0.045758516, -0.18718386, 1.5327792, 1.4693588, 0.15494743, 0.37816253, -0.88778573, -1.9807965, -0.34791216, 0.15634897, 1.2302907, 1.2023798, -0.3873268, -0.30230275, -1.048553, -1.420018, -1.7062702, 1.9507754, -0.5096522, -0.4380743, -1.2527953, 0.7774904, -1.6138978, -0.21274029, -0.89546657, 0.3869025, -0.51080513, -1.1806322, -0.028182229, 0.42833188, 0.06651722, 0.3024719, -0.6343221, -0.36274117, -0.67246044, -0.35955316, -0.8131463, -1.7262826, 0.17742614, -0.40178093, -1.6301984, 0.46278226, -0.9072984, 0.051945396, 0.7290906, 0.12898292, 1.1394007, -1.2348258, 0.40234163, -0.6848101, -0.87079716, -0.5788497, -0.31155252, 0.05616534, -1.1651498, 0.9008265, 0.46566245, -1.5362437, 1.4882522, 1.8958892, 1.1787796, -0.17992483, -1.0707526, 1.0544517, -0.40317693, 1.222445, 0.20827498, 0.97663903, 0.3563664, 0.7065732}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 3, 31),
				tensor.WithBacking([]float32{1.0821048, 0.68944764, 1.6098156, 2.0542254, 0.44514006, -0.013594717, 0.3993656, -0.12728803, 0.15368983, 0.27732104, 0.7991585, 1.1076555, 0.44135636, 0.28276914, 0.3887688, 0.9138767, 0.64446044, 0.05395472, -0.270514, -1.7035427, -0.9496856, 0.75902736, 0.06113559, 0.7637948, 0.40769446, -0.7043036, -0.07071267, 0.6727977, 1.5010691, 0.8121531, 0.26655498, -1.4342911, -1.1643543, -0.095781595, 0.6933198, 1.2163353, 0.4075265, -0.34481478, -0.67542785, -1.2342855, -1.5631441, 0.12225258, 0.7205616, -0.47386324, -0.8454348, -0.23765248, -0.4182037, -0.91331905, -0.55410343, -0.25428203, -0.06195131, -0.8457187, -0.60440725, 0.20007482, 0.24742454, 0.18449455, -0.1659251, -0.49853164, -0.5176008, -0.5160068, -0.5863497, -1.2697145, -0.112177394, -1.0159897, -0.58370805, -0.22225806, -0.4276765, 0.39051798, 0.42903674, 0.6341918, -0.047712564, -0.41624212, -0.14123423, -0.77780366, -0.7248234, -0.4452011, -0.1276936, -0.55449224, -0.13216165, 0.68324447, -0.5352906, -0.023995757, 1.6920707, 1.5373344, 0.49942738, -0.62533873, -0.008150458, 0.3256374, 0.40963405, 0.71536, 0.592457, 0.6665027, 0.5314698}),
			),
		},
	}
}
