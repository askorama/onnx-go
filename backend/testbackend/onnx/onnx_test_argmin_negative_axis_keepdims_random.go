package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("ArgMin", "TestArgminNegativeAxisKeepdimsRandom", NewTestArgminNegativeAxisKeepdimsRandom)
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
                Input:     {"data"},
                Output:    {"result"},
                Name:      "",
                OpType:    "ArgMin",
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
                    &ir.AttributeProto{
                        Name:          "keepdims",
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
        Name:              "test_argmin_negative_axis_keepdims_random",
        Initializer:       nil,
        SparseInitializer: nil,
        DocString:         "",
        Input:             {
            &ir.ValueInfoProto{
                Name: "data",
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
                Name: "result",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 7,
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
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

// NewTestArgminNegativeAxisKeepdimsRandom version: 6.
func NewTestArgminNegativeAxisKeepdimsRandom() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "ArgMin",
		Title:  "TestArgminNegativeAxisKeepdimsRandom",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xa4, 0x1, 0xa, 0x3d, 0xa, 0x4, 0x64, 0x61, 0x74, 0x61, 0x12, 0x6, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x6, 0x41, 0x72, 0x67, 0x4d, 0x69, 0x6e, 0x2a, 0x14, 0xa, 0x4, 0x61, 0x78, 0x69, 0x73, 0x18, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1, 0xa0, 0x1, 0x2, 0x2a, 0xf, 0xa, 0x8, 0x6b, 0x65, 0x65, 0x70, 0x64, 0x69, 0x6d, 0x73, 0x18, 0x1, 0xa0, 0x1, 0x2, 0x12, 0x29, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x72, 0x67, 0x6d, 0x69, 0x6e, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x64, 0x69, 0x6d, 0x73, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5a, 0x1a, 0xa, 0x4, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0x62, 0x1c, 0xa, 0x6, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0xa, 0x10, 0x8, 0x7, 0x12, 0xc, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x1, 0x42, 0x2, 0x10, 0xb},

		/*

		   &ir.NodeProto{
		     Input:     []string{"data"},
		     Output:    []string{"result"},
		     Name:      "",
		     OpType:    "ArgMin",
		     Attributes: ([]*ir.AttributeProto) (len=2 cap=2) {
		    (*ir.AttributeProto)(0xc000217500)(name:"axis" type:INT i:-1 ),
		    (*ir.AttributeProto)(0xc000217600)(name:"keepdims" type:INT i:1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 3, 4),
				tensor.WithBacking([]float32{0.9762701, 4.303787, 2.0552676, 0.89766365, -1.526904, 2.9178822, -1.2482557, 7.83546, 9.273255, -2.3311696, 5.834501, 0.5778984, 1.3608912, 8.511932, -8.579279, -8.257414, -9.595632, 6.6523967, 5.563135, 7.400243, 9.572367, 5.9831715, -0.77041274, 5.6105833}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 3, 1),
				tensor.WithBacking([]int64{3, 0, 1, 2, 0, 2}),
			),
		},
	}
}
