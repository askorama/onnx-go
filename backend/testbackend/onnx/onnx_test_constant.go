package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Constant", "TestConstant", NewTestConstant)
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
                Input:     nil,
                Output:    {"values"},
                Name:      "",
                OpType:    "Constant",
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
                            Dims:         {5, 5},
                            DataType:     1,
                            Segment:      (*ir.TensorProto_Segment)(nil),
                            FloatData:    {1.764052391052246, 0.40015721321105957, 0.978738009929657, 2.2408931255340576, 1.8675580024719238, -0.9772778749465942, 0.9500884413719177, -0.15135720372200012, -0.10321885347366333, 0.4105985164642334, 0.14404356479644775, 1.4542734622955322, 0.7610377073287964, 0.12167501449584961, 0.44386324286460876, 0.3336743414402008, 1.4940791130065918, -0.2051582634449005, 0.3130677044391632, -0.8540957570075989, -2.5529897212982178, 0.653618574142456, 0.8644362092018127, -0.7421650290489197, 2.269754648208618},
                            Int32Data:    nil,
                            StringData:   nil,
                            Int64Data:    nil,
                            Name:         "const_tensor",
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
        Name:        "test_constant",
        Initializer: nil,
        DocString:   "",
        Input:       nil,
        Output:      {
            &ir.ValueInfoProto{
                Name: "values",
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

// NewTestConstant version: 3.
func NewTestConstant() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Constant",
		Title:  "TestConstant",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xc7, 0x1, 0xa, 0x9b, 0x1, 0x12, 0x6, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x8, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2a, 0x86, 0x1, 0xa, 0x5, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x7a, 0x8, 0x5, 0x8, 0x5, 0x10, 0x1, 0x22, 0x64, 0x78, 0xcc, 0xe1, 0x3f, 0x68, 0xe1, 0xcc, 0x3e, 0x93, 0x8e, 0x7a, 0x3f, 0xcb, 0x6a, 0xf, 0x40, 0x24, 0xc, 0xef, 0x3f, 0xe2, 0x2e, 0x7a, 0xbf, 0xff, 0x38, 0x73, 0x3f, 0x62, 0xfd, 0x1a, 0xbe, 0x68, 0x64, 0xd3, 0xbd, 0xf8, 0x39, 0xd2, 0x3e, 0x28, 0x80, 0x13, 0x3e, 0xa2, 0x25, 0xba, 0x3f, 0x5e, 0xd3, 0x42, 0x3f, 0xc0, 0x30, 0xf9, 0x3d, 0xb, 0x42, 0xe3, 0x3e, 0x5d, 0xd7, 0xaa, 0x3e, 0xfc, 0x3d, 0xbf, 0x3f, 0x2, 0x15, 0x52, 0xbe, 0x69, 0x4a, 0xa0, 0x3e, 0x5, 0xa6, 0x5a, 0xbf, 0x2f, 0x64, 0x23, 0xc0, 0x8c, 0x53, 0x27, 0x3f, 0xb1, 0x4b, 0x5d, 0x3f, 0x87, 0xfe, 0x3d, 0xbf, 0xa9, 0x43, 0x11, 0x40, 0x42, 0xc, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x5f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0xa0, 0x1, 0x4, 0x12, 0xd, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x62, 0x18, 0xa, 0x6, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string(nil),
		     Output:    []string{"values"},
		     Name:      "",
		     OpType:    "Constant",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc0001287e0)(name:"value" type:TENSOR t:<dims:5 dims:5 data_type:1 float_data:1.7640524 float_data:0.4001572 float_data:0.978738 float_data:2.2408931 float_data:1.867558 float_data:-0.9772779 float_data:0.95008844 float_data:-0.1513572 float_data:-0.10321885 float_data:0.41059852 float_data:0.14404356 float_data:1.4542735 float_data:0.7610377 float_data:0.121675014 float_data:0.44386324 float_data:0.33367434 float_data:1.4940791 float_data:-0.20515826 float_data:0.3130677 float_data:-0.85409576 float_data:-2.5529897 float_data:0.6536186 float_data:0.8644362 float_data:-0.742165 float_data:2.2697546 name:"const_tensor" > )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(5, 5),
				tensor.WithBacking([]float32{1.7640524, 0.4001572, 0.978738, 2.2408931, 1.867558, -0.9772779, 0.95008844, -0.1513572, -0.10321885, 0.41059852, 0.14404356, 1.4542735, 0.7610377, 0.121675014, 0.44386324, 0.33367434, 1.4940791, -0.20515826, 0.3130677, -0.85409576, -2.5529897, 0.6536186, 0.8644362, -0.742165, 2.2697546}),
			),
		},
	}
}
