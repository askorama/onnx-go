package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("DepthToSpace", "TestDepthtospaceExample", NewTestDepthtospaceExample)
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
        Name:              "test_depthtospace_example",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:8},
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:4},
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

// NewTestDepthtospaceExample version: 5.
func NewTestDepthtospaceExample() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "DepthToSpace",
		Title:  "TestDepthtospaceExample",
		ModelB: []byte{0x8, 0x5, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x8d, 0x1, 0xa, 0x36, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0xc, 0x44, 0x65, 0x70, 0x74, 0x68, 0x54, 0x6f, 0x53, 0x70, 0x61, 0x63, 0x65, 0x2a, 0x10, 0xa, 0x9, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x2, 0xa0, 0x1, 0x2, 0x2a, 0xe, 0xa, 0x4, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x3, 0x44, 0x43, 0x52, 0xa0, 0x1, 0x3, 0x12, 0x19, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x68, 0x74, 0x6f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5a, 0x1b, 0xa, 0x1, 0x78, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x8, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x3, 0x62, 0x1b, 0xa, 0x1, 0x79, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x6, 0x42, 0x2, 0x10, 0xb},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "DepthToSpace",
		     Attributes: ([]*ir.AttributeProto) (len=2 cap=2) {
		    (*ir.AttributeProto)(0xc000217200)(name:"blocksize" type:INT i:2 ),
		    (*ir.AttributeProto)(0xc000217300)(name:"mode" type:STRING s:"DCR" )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 8, 2, 3),
				tensor.WithBacking([]float32{0, 1, 2, 3, 4, 5, 9, 10, 11, 12, 13, 14, 18, 19, 20, 21, 22, 23, 27, 28, 29, 30, 31, 32, 36, 37, 38, 39, 40, 41, 45, 46, 47, 48, 49, 50, 54, 55, 56, 57, 58, 59, 63, 64, 65, 66, 67, 68}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 2, 4, 6),
				tensor.WithBacking([]float32{0, 18, 1, 19, 2, 20, 36, 54, 37, 55, 38, 56, 3, 21, 4, 22, 5, 23, 39, 57, 40, 58, 41, 59, 9, 27, 10, 28, 11, 29, 45, 63, 46, 64, 47, 65, 12, 30, 13, 31, 14, 32, 48, 66, 49, 67, 50, 68}),
			),
		},
	}
}
