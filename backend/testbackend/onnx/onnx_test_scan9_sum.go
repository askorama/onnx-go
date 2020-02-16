package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Scan", "TestScan9Sum", NewTestScan9Sum)
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
                Input:     {"initial", "x"},
                Output:    {"y", "z"},
                Name:      "",
                OpType:    "Scan",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:        "body",
                        RefAttrName: "",
                        DocString:   "",
                        Type:        5,
                        F:           0,
                        I:           0,
                        S:           nil,
                        T:           (*ir.TensorProto)(nil),
                        G:           &ir.GraphProto{
                            Node: {
                                &ir.NodeProto{
                                    Input:     {"sum_in", "next"},
                                    Output:    {"sum_out"},
                                    Name:      "",
                                    OpType:    "Add",
                                    Domain:    "",
                                    Attribute: nil,
                                    DocString: "",
                                },
                                &ir.NodeProto{
                                    Input:     {"sum_out"},
                                    Output:    {"scan_out"},
                                    Name:      "",
                                    OpType:    "Identity",
                                    Domain:    "",
                                    Attribute: nil,
                                    DocString: "",
                                },
                            },
                            Name:        "scan_body",
                            Initializer: nil,
                            DocString:   "",
                            Input:       {
                                &ir.ValueInfoProto{
                                    Name: "sum_in",
                                    Type: &ir.TypeProto{
                                        Value: &ir.TypeProto_TensorType{
                                            TensorType: &ir.TypeProto_Tensor{
                                                ElemType: 1,
                                                Shape:    &ir.TensorShapeProto{
                                                    Dim: {
                                                        &!%v(DEPTH EXCEEDED),
                                                    },
                                                },
                                            },
                                        },
                                        Denotation: "",
                                    },
                                    DocString: "",
                                },
                                &ir.ValueInfoProto{
                                    Name: "next",
                                    Type: &ir.TypeProto{
                                        Value: &ir.TypeProto_TensorType{
                                            TensorType: &ir.TypeProto_Tensor{
                                                ElemType: 1,
                                                Shape:    &ir.TensorShapeProto{
                                                    Dim: {
                                                        &!%v(DEPTH EXCEEDED),
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
                                    Name: "sum_out",
                                    Type: &ir.TypeProto{
                                        Value: &ir.TypeProto_TensorType{
                                            TensorType: &ir.TypeProto_Tensor{
                                                ElemType: 1,
                                                Shape:    &ir.TensorShapeProto{
                                                    Dim: {
                                                        &!%v(DEPTH EXCEEDED),
                                                    },
                                                },
                                            },
                                        },
                                        Denotation: "",
                                    },
                                    DocString: "",
                                },
                                &ir.ValueInfoProto{
                                    Name: "scan_out",
                                    Type: &ir.TypeProto{
                                        Value: &ir.TypeProto_TensorType{
                                            TensorType: &ir.TypeProto_Tensor{
                                                ElemType: 1,
                                                Shape:    &ir.TensorShapeProto{
                                                    Dim: {
                                                        &!%v(DEPTH EXCEEDED),
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
                        Floats:  nil,
                        Ints:    nil,
                        Strings: nil,
                        Tensors: nil,
                        Graphs:  nil,
                    },
                    &ir.AttributeProto{
                        Name:        "num_scan_inputs",
                        RefAttrName: "",
                        DocString:   "",
                        Type:        2,
                        F:           0,
                        I:           1,
                        S:           nil,
                        T:           (*ir.TensorProto)(nil),
                        G:           (*ir.GraphProto)(nil),
                        Floats:      nil,
                        Ints:        nil,
                        Strings:     nil,
                        Tensors:     nil,
                        Graphs:      nil,
                    },
                },
                DocString: "",
            },
        },
        Name:        "test_scan9_sum",
        Initializer: nil,
        DocString:   "",
        Input:       {
            &ir.ValueInfoProto{
                Name: "initial",
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
                Name: "x",
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
                                },
                            },
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
            &ir.ValueInfoProto{
                Name: "z",
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
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestScan9Sum version: 3.
func NewTestScan9Sum() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Scan",
		Title:  "TestScan9Sum",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xc5, 0x2, 0xa, 0xe0, 0x1, 0xa, 0x7, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x12, 0x1, 0x7a, 0x22, 0x4, 0x53, 0x63, 0x61, 0x6e, 0x2a, 0xad, 0x1, 0xa, 0x4, 0x62, 0x6f, 0x64, 0x79, 0x32, 0xa1, 0x1, 0xa, 0x1c, 0xa, 0x6, 0x73, 0x75, 0x6d, 0x5f, 0x69, 0x6e, 0xa, 0x4, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x7, 0x73, 0x75, 0x6d, 0x5f, 0x6f, 0x75, 0x74, 0x22, 0x3, 0x41, 0x64, 0x64, 0xa, 0x1d, 0xa, 0x7, 0x73, 0x75, 0x6d, 0x5f, 0x6f, 0x75, 0x74, 0x12, 0x8, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x6f, 0x75, 0x74, 0x22, 0x8, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x9, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x5a, 0x14, 0xa, 0x6, 0x73, 0x75, 0x6d, 0x5f, 0x69, 0x6e, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x5a, 0x12, 0xa, 0x4, 0x6e, 0x65, 0x78, 0x74, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x62, 0x15, 0xa, 0x7, 0x73, 0x75, 0x6d, 0x5f, 0x6f, 0x75, 0x74, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x62, 0x16, 0xa, 0x8, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x6f, 0x75, 0x74, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0xa0, 0x1, 0x5, 0x2a, 0x16, 0xa, 0xf, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x1, 0xa0, 0x1, 0x2, 0x12, 0xe, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x63, 0x61, 0x6e, 0x39, 0x5f, 0x73, 0x75, 0x6d, 0x5a, 0x15, 0xa, 0x7, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x5a, 0x13, 0xa, 0x1, 0x78, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x2, 0x62, 0xf, 0xa, 0x1, 0x79, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x62, 0x13, 0xa, 0x1, 0x7a, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x2, 0x42, 0x4, 0xa, 0x0, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"initial", "x"},
		     Output:    []string{"y", "z"},
		     Name:      "",
		     OpType:    "Scan",
		     Attributes: ([]*ir.AttributeProto) (len=2 cap=2) {
		    (*ir.AttributeProto)(0xc0001761c0)(name:"body" type:GRAPH g:<node:<input:"sum_in" input:"next" output:"sum_out" op_type:"Add" > node:<input:"sum_out" output:"scan_out" op_type:"Identity" > name:"scan_body" input:<name:"sum_in" type:<tensor_type:<elem_type:1 shape:<dim:<dim_value:2 > > > > > input:<name:"next" type:<tensor_type:<elem_type:1 shape:<dim:<dim_value:2 > > > > > output:<name:"sum_out" type:<tensor_type:<elem_type:1 shape:<dim:<dim_value:2 > > > > > output:<name:"scan_out" type:<tensor_type:<elem_type:1 shape:<dim:<dim_value:2 > > > > > > ),
		    (*ir.AttributeProto)(0xc0001762a0)(name:"num_scan_inputs" type:INT i:1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{0, 0}),
			),

			tensor.New(
				tensor.WithShape(3, 2),
				tensor.WithBacking([]float32{1, 2, 3, 4, 5, 6}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{9, 12}),
			),

			tensor.New(
				tensor.WithShape(3, 2),
				tensor.WithBacking([]float32{1, 2, 4, 6, 9, 12}),
			),
		},
	}
}
