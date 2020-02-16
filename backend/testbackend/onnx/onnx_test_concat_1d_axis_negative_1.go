package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Concat", "TestConcat1dAxisNegative1", NewTestConcat1dAxisNegative1)
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
                Input:     {"value0", "value1"},
                Output:    {"output"},
                Name:      "",
                OpType:    "Concat",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:          "axis",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          2,
                        F:             0,
                        I:             -1,
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
        Name:              "test_concat_1d_axis_negative_1",
        Initializer:       nil,
        SparseInitializer: nil,
        DocString:         "",
        Input:             {
            &ir.ValueInfoProto{
                Name: "value0",
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
                                },
                            },
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
            &ir.ValueInfoProto{
                Name: "value1",
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
                Name: "output",
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

// NewTestConcat1dAxisNegative1 version: 6.
func NewTestConcat1dAxisNegative1() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Concat",
		Title:  "TestConcat1dAxisNegative1",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x9a, 0x1, 0xa, 0x36, 0xa, 0x6, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x30, 0xa, 0x6, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x12, 0x6, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x6, 0x43, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x2a, 0x14, 0xa, 0x4, 0x61, 0x78, 0x69, 0x73, 0x18, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1, 0xa0, 0x1, 0x2, 0x12, 0x1e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x5f, 0x31, 0x64, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x31, 0x5a, 0x14, 0xa, 0x6, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x30, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x5a, 0x14, 0xa, 0x6, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x62, 0x14, 0xa, 0x6, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x4, 0x42, 0x2, 0x10, 0xb},

		/*

		   &ir.NodeProto{
		     Input:     []string{"value0", "value1"},
		     Output:    []string{"output"},
		     Name:      "",
		     OpType:    "Concat",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc0001e6700)(name:"axis" type:INT i:-1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{1, 2}),
			),

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{3, 4}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(4),
				tensor.WithBacking([]float32{1, 2, 3, 4}),
			),
		},
	}
}
