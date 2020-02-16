package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("ConstantOfShape", "TestConstantofshapeFloatOnes", NewTestConstantofshapeFloatOnes)
}

/*
&ir.ModelProto{
    IrVersion:   4,
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
                OpType:    "ConstantOfShape",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:        "value",
                        RefAttrName: "",
                        DocString:   "",
                        Type:        4,
                        F:           0,
                        I:           0,
                        S:           nil,
                        T:           &ir.TensorProto{
                            Dims:         {1},
                            DataType:     1,
                            Segment:      (*ir.TensorProto_Segment)(nil),
                            FloatData:    {1},
                            Int32Data:    nil,
                            StringData:   nil,
                            Int64Data:    nil,
                            Name:         "value",
                            DocString:    "",
                            RawData:      nil,
                            ExternalData: nil,
                            DataLocation: 0,
                            DoubleData:   nil,
                            Uint64Data:   nil,
                        },
                        G:       (*ir.GraphProto)(nil),
                        Floats:  nil,
                        Ints:    nil,
                        Strings: nil,
                        Tensors: nil,
                        Graphs:  nil,
                    },
                },
                DocString: "",
            },
        },
        Name:        "test_constantofshape_float_ones",
        Initializer: nil,
        DocString:   "",
        Input:       {
            &ir.ValueInfoProto{
                Name: "x",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 7,
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:4},
                                        Denotation: "",
                                    },
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

// NewTestConstantofshapeFloatOnes version: 4.
func NewTestConstantofshapeFloatOnes() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "ConstantOfShape",
		Title:  "TestConstantofshapeFloatOnes",
		ModelB: []byte{0x8, 0x4, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x83, 0x1, 0xa, 0x36, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0xf, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4f, 0x66, 0x53, 0x68, 0x61, 0x70, 0x65, 0x2a, 0x1d, 0xa, 0x5, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x11, 0x8, 0x1, 0x10, 0x1, 0x22, 0x4, 0x0, 0x0, 0x80, 0x3f, 0x42, 0x5, 0x76, 0x61, 0x6c, 0x75, 0x65, 0xa0, 0x1, 0x4, 0x12, 0x1f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x6f, 0x66, 0x73, 0x68, 0x61, 0x70, 0x65, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x6f, 0x6e, 0x65, 0x73, 0x5a, 0xf, 0xa, 0x1, 0x78, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x62, 0x17, 0xa, 0x1, 0x79, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x2, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "ConstantOfShape",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc0001282a0)(name:"value" type:TENSOR t:<dims:1 data_type:1 float_data:1 name:"value" > )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]int64{4, 3, 2}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(4, 3, 2),
				tensor.WithBacking([]float32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}),
			),
		},
	}
}
