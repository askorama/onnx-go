package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Cast", "TestCastDOUBLEToFLOAT", NewTestCastDOUBLEToFLOAT)
}

/*
&ir.ModelProto{
    IrVersion:   4,
    OpsetImport: {
        &ir.OperatorSetIdProto{Domain:"", Version:10},
    },
    ProducerName:    "backend-test",
    ProducerVersion: "",
    Domain:          "",
    ModelVersion:    0,
    DocString:       "",
    Graph:           &ir.GraphProto{
        Node: {
            &ir.NodeProto{
                Input:     {"input"},
                Output:    {"output"},
                Name:      "",
                OpType:    "Cast",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:          "to",
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
        Name:              "test_cast_DOUBLE_to_FLOAT",
        Initializer:       nil,
        SparseInitializer: nil,
        DocString:         "",
        Input:             {
            &ir.ValueInfoProto{
                Name: "input",
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
                Name: "output",
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

// NewTestCastDOUBLEToFLOAT version: 4.
func NewTestCastDOUBLEToFLOAT() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Cast",
		Title:  "TestCastDOUBLEToFLOAT",
		ModelB: []byte{0x8, 0x4, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x70, 0xa, 0x20, 0xa, 0x5, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x6, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x4, 0x43, 0x61, 0x73, 0x74, 0x2a, 0x9, 0xa, 0x2, 0x74, 0x6f, 0x18, 0x1, 0xa0, 0x1, 0x2, 0x12, 0x19, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x5f, 0x74, 0x6f, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x5a, 0x17, 0xa, 0x5, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0xe, 0xa, 0xc, 0x8, 0xb, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0x62, 0x18, 0xa, 0x6, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0x42, 0x2, 0x10, 0xa},

		/*

		   &ir.NodeProto{
		     Input:     []string{"input"},
		     Output:    []string{"output"},
		     Name:      "",
		     OpType:    "Cast",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc000216e00)(name:"to" type:INT i:1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4),
				tensor.WithBacking([]float64{0.31542835092418386, 0.3637107709426226, 0.5701967704178796, 0.43860151346232035, 0.9883738380592262, 0.10204481074802807, 0.2088767560948347, 0.16130951788499626, 0.6531083254653984, 0.2532916025397821, 0.4663107728563063, 0.24442559200160274}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4),
				tensor.WithBacking([]float32{0.31542835, 0.36371076, 0.57019675, 0.43860152, 0.9883738, 0.10204481, 0.20887676, 0.16130951, 0.6531083, 0.2532916, 0.46631077, 0.2444256}),
			),
		},
	}
}
