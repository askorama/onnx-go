package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Gemm", "TestGemmAllAttributes", NewTestGemmAllAttributes)
}

/*
&ir.ModelProto{
    IrVersion:   6,
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
                Input:     {"a", "b", "c"},
                Output:    {"y"},
                Name:      "",
                OpType:    "Gemm",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:          "alpha",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          1,
                        F:             0.25,
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
                    &ir.AttributeProto{
                        Name:          "beta",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          1,
                        F:             0.3499999940395355,
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
                    &ir.AttributeProto{
                        Name:          "transA",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          2,
                        F:             0,
                        I:             1,
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
                        Name:          "transB",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          2,
                        F:             0,
                        I:             1,
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
        Name:              "test_gemm_all_attributes",
        Initializer:       nil,
        SparseInitializer: nil,
        DocString:         "",
        Input:             {
            &ir.ValueInfoProto{
                Name: "a",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:4},
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
                Name: "b",
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
            &ir.ValueInfoProto{
                Name: "c",
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
        },
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestGemmAllAttributes version: 6.
func NewTestGemmAllAttributes() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Gemm",
		Title:  "TestGemmAllAttributes",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xc1, 0x1, 0xa, 0x51, 0xa, 0x1, 0x61, 0xa, 0x1, 0x62, 0xa, 0x1, 0x63, 0x12, 0x1, 0x79, 0x22, 0x4, 0x47, 0x65, 0x6d, 0x6d, 0x2a, 0xf, 0xa, 0x5, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x15, 0x0, 0x0, 0x80, 0x3e, 0xa0, 0x1, 0x1, 0x2a, 0xe, 0xa, 0x4, 0x62, 0x65, 0x74, 0x61, 0x15, 0x33, 0x33, 0xb3, 0x3e, 0xa0, 0x1, 0x1, 0x2a, 0xd, 0xa, 0x6, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x41, 0x18, 0x1, 0xa0, 0x1, 0x2, 0x2a, 0xd, 0xa, 0x6, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x42, 0x18, 0x1, 0xa0, 0x1, 0x2, 0x12, 0x18, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x67, 0x65, 0x6d, 0x6d, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x5a, 0x13, 0xa, 0x1, 0x61, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x3, 0x5a, 0x13, 0xa, 0x1, 0x62, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x4, 0x5a, 0x13, 0xa, 0x1, 0x63, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x5, 0x62, 0x13, 0xa, 0x1, 0x79, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0xb},

		/*

		   &ir.NodeProto{
		     Input:     []string{"a", "b", "c"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Gemm",
		     Attributes: ([]*ir.AttributeProto) (len=4 cap=4) {
		    (*ir.AttributeProto)(0xc000216d00)(name:"alpha" type:FLOAT f:0.25 ),
		    (*ir.AttributeProto)(0xc000216e00)(name:"beta" type:FLOAT f:0.35 ),
		    (*ir.AttributeProto)(0xc000216f00)(name:"transA" type:INT i:1 ),
		    (*ir.AttributeProto)(0xc000217000)(name:"transB" type:INT i:1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(4, 3),
				tensor.WithBacking([]float32{0.5488135, 0.71518934, 0.60276335, 0.5448832, 0.4236548, 0.6458941, 0.4375872, 0.891773, 0.96366274, 0.3834415, 0.79172504, 0.5288949}),
			),

			tensor.New(
				tensor.WithShape(5, 4),
				tensor.WithBacking([]float32{0.56804454, 0.92559665, 0.071036056, 0.0871293, 0.020218397, 0.83261985, 0.77815676, 0.87001216, 0.9786183, 0.7991586, 0.46147937, 0.7805292, 0.11827443, 0.639921, 0.14335328, 0.9446689, 0.5218483, 0.41466194, 0.2645556, 0.7742337}),
			),

			tensor.New(
				tensor.WithShape(1, 5),
				tensor.WithBacking([]float32{0.45615032, 0.56843394, 0.0187898, 0.6176355, 0.6120957}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 5),
				tensor.WithBacking([]float32{0.37979913, 0.4836736, 0.37501436, 0.42580932, 0.44547832, 0.39233345, 0.63644016, 0.52356744, 0.5240351, 0.5636829, 0.42334533, 0.63895106, 0.49747032, 0.49676925, 0.5259358}),
			),
		},
	}
}
