package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("RoiAlign", "TestRoialign", NewTestRoialign)
}

/*
&ir.ModelProto{
    IrVersion:   5,
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
                Input:     {"X", "rois", "batch_indices"},
                Output:    {"Y"},
                Name:      "",
                OpType:    "RoiAlign",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:        "output_height",
                        RefAttrName: "",
                        DocString:   "",
                        Type:        2,
                        F:           0,
                        I:           5,
                        S:           nil,
                        T:           (*ir.TensorProto)(nil),
                        G:           (*ir.GraphProto)(nil),
                        Floats:      nil,
                        Ints:        nil,
                        Strings:     nil,
                        Tensors:     nil,
                        Graphs:      nil,
                    },
                    &ir.AttributeProto{
                        Name:        "output_width",
                        RefAttrName: "",
                        DocString:   "",
                        Type:        2,
                        F:           0,
                        I:           5,
                        S:           nil,
                        T:           (*ir.TensorProto)(nil),
                        G:           (*ir.GraphProto)(nil),
                        Floats:      nil,
                        Ints:        nil,
                        Strings:     nil,
                        Tensors:     nil,
                        Graphs:      nil,
                    },
                    &ir.AttributeProto{
                        Name:        "sampling_ratio",
                        RefAttrName: "",
                        DocString:   "",
                        Type:        2,
                        F:           0,
                        I:           2,
                        S:           nil,
                        T:           (*ir.TensorProto)(nil),
                        G:           (*ir.GraphProto)(nil),
                        Floats:      nil,
                        Ints:        nil,
                        Strings:     nil,
                        Tensors:     nil,
                        Graphs:      nil,
                    },
                    &ir.AttributeProto{
                        Name:        "spatial_scale",
                        RefAttrName: "",
                        DocString:   "",
                        Type:        1,
                        F:           1,
                        I:           0,
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
        Name:        "test_roialign",
        Initializer: nil,
        DocString:   "",
        Input:       {
            &ir.ValueInfoProto{
                Name: "X",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:10},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:10},
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
                Name: "rois",
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
            &ir.ValueInfoProto{
                Name: "batch_indices",
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
                Name: "Y",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
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
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestRoialign version: 5.
func NewTestRoialign() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "RoiAlign",
		Title:  "TestRoialign",
		ModelB: []byte{0x8, 0x5, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x81, 0x2, 0xa, 0x80, 0x1, 0xa, 0x1, 0x58, 0xa, 0x4, 0x72, 0x6f, 0x69, 0x73, 0xa, 0xd, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1, 0x59, 0x22, 0x8, 0x52, 0x6f, 0x69, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x2a, 0x14, 0xa, 0xd, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x5, 0xa0, 0x1, 0x2, 0x2a, 0x13, 0xa, 0xc, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x5, 0xa0, 0x1, 0x2, 0x2a, 0x15, 0xa, 0xe, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x2, 0xa0, 0x1, 0x2, 0x2a, 0x17, 0xa, 0xd, 0x73, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x15, 0x0, 0x0, 0x80, 0x3f, 0xa0, 0x1, 0x1, 0x12, 0xd, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x6f, 0x69, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x5a, 0x1b, 0xa, 0x1, 0x58, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0xa, 0xa, 0x2, 0x8, 0xa, 0x5a, 0x16, 0xa, 0x4, 0x72, 0x6f, 0x69, 0x73, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0x5a, 0x1b, 0xa, 0xd, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x62, 0x1b, 0xa, 0x1, 0x59, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0xa},

		/*

		   &ir.NodeProto{
		     Input:     []string{"X", "rois", "batch_indices"},
		     Output:    []string{"Y"},
		     Name:      "",
		     OpType:    "RoiAlign",
		     Attributes: ([]*ir.AttributeProto) (len=4 cap=4) {
		    (*ir.AttributeProto)(0xc000536380)(name:"output_height" type:INT i:5 ),
		    (*ir.AttributeProto)(0xc000536460)(name:"output_width" type:INT i:5 ),
		    (*ir.AttributeProto)(0xc000536540)(name:"sampling_ratio" type:INT i:2 ),
		    (*ir.AttributeProto)(0xc000536620)(name:"spatial_scale" type:FLOAT f:1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 1, 10, 10),
				tensor.WithBacking([]float32{0.2764, 0.715, 0.1958, 0.3416, 0.4638, 0.0259, 0.2963, 0.6518, 0.4856, 0.725, 0.9637, 0.0895, 0.2919, 0.6753, 0.0234, 0.6132, 0.8085, 0.5324, 0.8992, 0.4467, 0.3265, 0.8479, 0.9698, 0.2471, 0.9336, 0.1878, 0.4766, 0.4308, 0.34, 0.2162, 0.0206, 0.172, 0.2155, 0.4394, 0.0653, 0.3406, 0.7724, 0.3921, 0.2541, 0.5799, 0.4062, 0.2194, 0.4473, 0.4687, 0.7109, 0.9327, 0.9815, 0.632, 0.1728, 0.6119, 0.3097, 0.1283, 0.4984, 0.5068, 0.4279, 0.0173, 0.4388, 0.043, 0.4671, 0.7119, 0.1011, 0.8477, 0.4726, 0.1777, 0.9923, 0.4042, 0.1869, 0.7795, 0.9946, 0.9689, 0.1366, 0.3671, 0.7011, 0.6234, 0.9867, 0.5585, 0.6985, 0.5609, 0.8788, 0.9928, 0.5697, 0.8511, 0.6711, 0.9406, 0.8751, 0.7496, 0.165, 0.1049, 0.1559, 0.2514, 0.7012, 0.4056, 0.7879, 0.3461, 0.0415, 0.2998, 0.5094, 0.3727, 0.5482, 0.0502}),
			),

			tensor.New(
				tensor.WithShape(3, 4),
				tensor.WithBacking([]float32{0, 0, 9, 9, 0, 5, 4, 9, 5, 5, 9, 9}),
			),

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]int64{0, 0, 0}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 1, 5, 5),
				tensor.WithBacking([]float32{0.4664, 0.4466, 0.3405, 0.5688, 0.6068, 0.3714, 0.4296, 0.3835, 0.5562, 0.351, 0.2768, 0.4883, 0.5222, 0.5528, 0.4171, 0.4713, 0.4844, 0.6904, 0.492, 0.8774, 0.6239, 0.7125, 0.6289, 0.3355, 0.3495, 0.3022, 0.4305, 0.4696, 0.3978, 0.5423, 0.3656, 0.705, 0.5165, 0.3172, 0.7015, 0.2912, 0.5059, 0.6476, 0.6235, 0.8299, 0.5916, 0.7389, 0.7048, 0.8372, 0.8893, 0.6227, 0.6153, 0.7097, 0.6154, 0.4585, 0.2384, 0.3379, 0.3717, 0.61, 0.7601, 0.3767, 0.3785, 0.7147, 0.9243, 0.9727, 0.5749, 0.5826, 0.5709, 0.7619, 0.877, 0.5355, 0.2566, 0.2141, 0.2796, 0.36, 0.4365, 0.3504, 0.2887, 0.3661, 0.2349}),
			),
		},
	}
}
