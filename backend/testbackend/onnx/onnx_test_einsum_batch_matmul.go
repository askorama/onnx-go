package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Einsum", "TestEinsumBatchMatmul", NewTestEinsumBatchMatmul)
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
                        S:             {0x62, 0x69, 0x6a, 0x2c, 0x20, 0x62, 0x6a, 0x6b, 0x20, 0x2d, 0x3e, 0x20, 0x62, 0x69, 0x6b},
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
        Name:              "test_einsum_batch_matmul",
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
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
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
                Name: "z",
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
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
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

// NewTestEinsumBatchMatmul version: 6.
func NewTestEinsumBatchMatmul() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Einsum",
		Title:  "TestEinsumBatchMatmul",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x98, 0x1, 0xa, 0x31, 0xa, 0x1, 0x78, 0xa, 0x1, 0x79, 0x12, 0x1, 0x7a, 0x22, 0x6, 0x45, 0x69, 0x6e, 0x73, 0x75, 0x6d, 0x2a, 0x1e, 0xa, 0x8, 0x65, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf, 0x62, 0x69, 0x6a, 0x2c, 0x20, 0x62, 0x6a, 0x6b, 0x20, 0x2d, 0x3e, 0x20, 0x62, 0x69, 0x6b, 0xa0, 0x1, 0x3, 0x12, 0x18, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x69, 0x6e, 0x73, 0x75, 0x6d, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6d, 0x61, 0x74, 0x6d, 0x75, 0x6c, 0x5a, 0x17, 0xa, 0x1, 0x78, 0x12, 0x12, 0xa, 0x10, 0x8, 0xb, 0x12, 0xc, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x3, 0x5a, 0x17, 0xa, 0x1, 0x79, 0x12, 0x12, 0xa, 0x10, 0x8, 0xb, 0x12, 0xc, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0x62, 0x17, 0xa, 0x1, 0x7a, 0x12, 0x12, 0xa, 0x10, 0x8, 0xb, 0x12, 0xc, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x4, 0x42, 0x2, 0x10, 0xc},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x", "y"},
		     Output:    []string{"z"},
		     Name:      "",
		     OpType:    "Einsum",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc000216b00)(name:"equation" type:STRING s:"bij, bjk -> bik" )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(5, 2, 3),
				tensor.WithBacking([]float64{1.764052345967664, 0.4001572083672233, 0.9787379841057392, 2.240893199201458, 1.8675579901499675, -0.977277879876411, 0.9500884175255894, -0.1513572082976979, -0.10321885179355784, 0.41059850193837233, 0.144043571160878, 1.454273506962975, 0.7610377251469934, 0.12167501649282841, 0.44386323274542566, 0.33367432737426683, 1.4940790731576061, -0.20515826376580087, 0.31306770165090136, -0.8540957393017248, -2.5529898158340787, 0.6536185954403606, 0.8644361988595057, -0.7421650204064419, 2.2697546239876076, -1.4543656745987648, 0.04575851730144607, -0.1871838500258336, 1.5327792143584575, 1.469358769900285}),
			),

			tensor.New(
				tensor.WithShape(5, 3, 4),
				tensor.WithBacking([]float64{0.1549474256969163, 0.37816251960217356, -0.8877857476301128, -1.980796468223927, -0.3479121493261526, 0.15634896910398005, 1.2302906807277207, 1.2023798487844113, -0.3873268174079523, -0.30230275057533557, -1.0485529650670926, -1.4200179371789752, -1.7062701906250126, 1.9507753952317897, -0.5096521817516535, -0.4380743016111864, -1.2527953600499262, 0.7774903558319101, -1.6138978475579515, -0.2127402802139687, -0.8954665611936756, 0.386902497859262, -0.510805137568873, -1.180632184122412, -0.028182228338654868, 0.42833187053041766, 0.06651722238316789, 0.3024718977397814, -0.6343220936809636, -0.3627411659871381, -0.672460447775951, -0.3595531615405413, -0.813146282044454, -1.7262826023316769, 0.17742614225375283, -0.4017809362082619, -1.6301983469660446, 0.4627822555257742, -0.9072983643832422, 0.05194539579613895, 0.7290905621775369, 0.12898291075741067, 1.1394006845433007, -1.2348258203536526, 0.402341641177549, -0.6848100909403132, -0.8707971491818818, -0.5788496647644155, -0.31155253212737266, 0.05616534222974544, -1.1651498407833565, 0.9008264869541871, 0.46566243973045984, -1.5362436862772237, 1.4882521937955997, 1.8958891760305832, 1.1787795711596507, -0.17992483581235091, -1.0707526215105425, 1.0544517269311366}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(5, 2, 4),
				tensor.WithBacking([]float64{-0.2449756530890492, 0.43378746218127495, -2.1000494618012033, -4.402913186196848, 0.07600044906940934, 1.4348463760343186, 1.3329337654266886, -0.8054871218198842, -1.339058906750601, 1.6957947068814332, -0.18721484237474798, -0.26214604663510777, -2.1833023980833026, 1.475639994712103, -1.1845844105749435, -1.9274786285080947, -0.45955462759799476, -0.48439320150081144, 0.047553420605969635, 0.008108102915400511, -0.7903073724831986, -0.04487949481849868, -1.0189144324021775, -0.3538450680307544, -2.1602497048529434, 1.7830316105326642, 0.9659331698437077, 2.5487191966595786, -0.7338795719749507, 0.9222226800617044, 1.0381872981434612, -0.6038736884980718, -1.3304520633035486, 2.353508536662, -4.858063196916035, -0.664410908697119, 2.5041254116379257, -2.629609770935472, 0.9259395066178495, 4.286727244116296}),
			),
		},
	}
}
