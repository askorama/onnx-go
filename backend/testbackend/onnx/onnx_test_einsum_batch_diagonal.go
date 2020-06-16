package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Einsum", "TestEinsumBatchDiagonal", NewTestEinsumBatchDiagonal)
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
                        S:             {0x2e, 0x2e, 0x2e, 0x69, 0x69, 0x20, 0x2d, 0x3e, 0x2e, 0x2e, 0x2e, 0x69},
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
        Name:              "test_einsum_batch_diagonal",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:5},
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

// NewTestEinsumBatchDiagonal version: 6.
func NewTestEinsumBatchDiagonal() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Einsum",
		Title:  "TestEinsumBatchDiagonal",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x77, 0xa, 0x2b, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x6, 0x45, 0x69, 0x6e, 0x73, 0x75, 0x6d, 0x2a, 0x1b, 0xa, 0x8, 0x65, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc, 0x2e, 0x2e, 0x2e, 0x69, 0x69, 0x20, 0x2d, 0x3e, 0x2e, 0x2e, 0x2e, 0x69, 0xa0, 0x1, 0x3, 0x12, 0x1a, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x69, 0x6e, 0x73, 0x75, 0x6d, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x69, 0x61, 0x67, 0x6f, 0x6e, 0x61, 0x6c, 0x5a, 0x17, 0xa, 0x1, 0x78, 0x12, 0x12, 0xa, 0x10, 0x8, 0xb, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x5, 0x62, 0x13, 0xa, 0x1, 0x79, 0x12, 0xe, 0xa, 0xc, 0x8, 0xb, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0xc},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Einsum",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc000217a00)(name:"equation" type:STRING s:"...ii ->...i" )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 5, 5),
				tensor.WithBacking([]float64{1.764052345967664, 0.4001572083672233, 0.9787379841057392, 2.240893199201458, 1.8675579901499675, -0.977277879876411, 0.9500884175255894, -0.1513572082976979, -0.10321885179355784, 0.41059850193837233, 0.144043571160878, 1.454273506962975, 0.7610377251469934, 0.12167501649282841, 0.44386323274542566, 0.33367432737426683, 1.4940790731576061, -0.20515826376580087, 0.31306770165090136, -0.8540957393017248, -2.5529898158340787, 0.6536185954403606, 0.8644361988595057, -0.7421650204064419, 2.2697546239876076, -1.4543656745987648, 0.04575851730144607, -0.1871838500258336, 1.5327792143584575, 1.469358769900285, 0.1549474256969163, 0.37816251960217356, -0.8877857476301128, -1.980796468223927, -0.3479121493261526, 0.15634896910398005, 1.2302906807277207, 1.2023798487844113, -0.3873268174079523, -0.30230275057533557, -1.0485529650670926, -1.4200179371789752, -1.7062701906250126, 1.9507753952317897, -0.5096521817516535, -0.4380743016111864, -1.2527953600499262, 0.7774903558319101, -1.6138978475579515, -0.2127402802139687, -0.8954665611936756, 0.386902497859262, -0.510805137568873, -1.180632184122412, -0.028182228338654868, 0.42833187053041766, 0.06651722238316789, 0.3024718977397814, -0.6343220936809636, -0.3627411659871381, -0.672460447775951, -0.3595531615405413, -0.813146282044454, -1.7262826023316769, 0.17742614225375283, -0.4017809362082619, -1.6301983469660446, 0.4627822555257742, -0.9072983643832422, 0.05194539579613895, 0.7290905621775369, 0.12898291075741067, 1.1394006845433007, -1.2348258203536526, 0.402341641177549}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 5),
				tensor.WithBacking([]float64{1.764052345967664, 0.9500884175255894, 0.7610377251469934, 0.31306770165090136, 2.2697546239876076, -1.4543656745987648, 0.37816251960217356, 1.2023798487844113, 1.9507753952317897, -0.2127402802139687, -0.8954665611936756, 0.06651722238316789, -0.813146282044454, -0.9072983643832422, 0.402341641177549}),
			),
		},
	}
}
