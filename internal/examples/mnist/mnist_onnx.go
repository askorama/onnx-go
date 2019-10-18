package mnist

import (
	"github.com/owulveryck/onnx-go/internal/onnx/ir"
)

// Mnist model represented in an onnx compatible format
var Mnist = &ir.ModelProto{
	IrVersion: 3,
	OpsetImport: []*ir.OperatorSetIdProto{
		{
			Domain:  "",
			Version: 7,
		},
	},
	ProducerName:    "CNTK",
	ProducerVersion: "2.5.1",
	Domain:          "ai.cntk",
	ModelVersion:    1,
	Graph: &ir.GraphProto{
		Node: []*ir.NodeProto{
			{
				Input:     []string{"Parameter193", "Parameter193_reshape1_shape"},
				Output:    []string{"Parameter193_reshape1"},
				Name:      "Times212_reshape1",
				OpType:    "Reshape",
				Domain:    "",
				DocString: "",
			},
			{
				Input:  []string{"Input3", "Parameter5"},
				Output: []string{"Convolution28_Output_0"},
				Name:   "Convolution28",
				OpType: "Conv",
				Domain: "",
				Attribute: []*ir.AttributeProto{
					{
						Name: "kernel_shape",
						Type: 7,
						Ints: []int64{5, 5},
					},
					{
						Name: "strides",
						Type: 7,
						Ints: []int64{1, 1},
					},
					{
						Name: "auto_pad",
						Type: 3,
						S:    []byte{0x53, 0x41, 0x4d, 0x45, 0x5f, 0x55, 0x50, 0x50, 0x45, 0x52}, // SAME_UPPER
					},
					{
						Name: "group",
						Type: 2,
						I:    1,
					},
					{
						Name: "dilations",
						Type: 7,
						Ints: []int64{1, 1},
					},
				},
				DocString: "",
			},
			{
				Input:     []string{"Convolution28_Output_0", "Parameter6"},
				Output:    []string{"Plus30_Output_0"},
				Name:      "Plus30",
				OpType:    "Add",
				Domain:    "",
				DocString: "",
			},
			{
				Input:     []string{"Plus30_Output_0"},
				Output:    []string{"ReLU32_Output_0"},
				Name:      "ReLU32",
				OpType:    "Relu",
				Domain:    "",
				DocString: "",
			},
			{
				Input:  []string{"ReLU32_Output_0"},
				Output: []string{"Pooling66_Output_0"},
				Name:   "Pooling66",
				OpType: "MaxPool",
				Domain: "",
				Attribute: []*ir.AttributeProto{
					{
						Name: "kernel_shape",
						Type: 7,
						Ints: []int64{2, 2},
					},
					{
						Name: "strides",
						Type: 7,
						Ints: []int64{2, 2},
					},
					{
						Name: "pads",
						Type: 7,
						Ints: []int64{0, 0, 0, 0},
					},
					{
						Name: "auto_pad",
						Type: 3,
						S:    []byte{0x4e, 0x4f, 0x54, 0x53, 0x45, 0x54}, // NOTSET
					},
				},
				DocString: "",
			},
			{
				Input:  []string{"Pooling66_Output_0", "Parameter87"},
				Output: []string{"Convolution110_Output_0"},
				Name:   "Convolution110",
				OpType: "Conv",
				Domain: "",
				Attribute: []*ir.AttributeProto{
					{
						Name: "kernel_shape",
						Type: 7,
						Ints: []int64{5, 5},
					},
					{
						Name: "strides",
						Type: 7,
						Ints: []int64{1, 1},
					},
					{
						Name: "auto_pad",
						Type: 3,
						S:    []byte{0x53, 0x41, 0x4d, 0x45, 0x5f, 0x55, 0x50, 0x50, 0x45, 0x52}, // SAME_UPPER
					},
					{
						Name: "group",
						Type: 2,
						I:    1,
					},
					{
						Name: "dilations",
						Type: 7,
						Ints: []int64{1, 1},
					},
				},
				DocString: "",
			},
			{
				Input:     []string{"Convolution110_Output_0", "Parameter88"},
				Output:    []string{"Plus112_Output_0"},
				Name:      "Plus112",
				OpType:    "Add",
				Domain:    "",
				DocString: "",
			},
			{
				Input:     []string{"Plus112_Output_0"},
				Output:    []string{"ReLU114_Output_0"},
				Name:      "ReLU114",
				OpType:    "Relu",
				Domain:    "",
				DocString: "",
			},
			{
				Input:  []string{"ReLU114_Output_0"},
				Output: []string{"Pooling160_Output_0"},
				Name:   "Pooling160",
				OpType: "MaxPool",
				Domain: "",
				Attribute: []*ir.AttributeProto{
					{
						Name: "kernel_shape",
						Type: 7,
						Ints: []int64{3, 3},
					},
					{
						Name: "strides",
						Type: 7,
						Ints: []int64{3, 3},
					},
					{
						Name: "pads",
						Type: 7,
						Ints: []int64{0, 0, 0, 0},
					},
					{
						Name: "auto_pad",
						Type: 3,
						S:    []byte{0x4e, 0x4f, 0x54, 0x53, 0x45, 0x54}, // NOTSET
					},
				},
				DocString: "",
			},
			{
				Input:     []string{"Pooling160_Output_0", "Pooling160_Output_0_reshape0_shape"},
				Output:    []string{"Pooling160_Output_0_reshape0"},
				Name:      "Times212_reshape0",
				OpType:    "Reshape",
				Domain:    "",
				DocString: "",
			},
			{
				Input:     []string{"Pooling160_Output_0_reshape0", "Parameter193_reshape1"},
				Output:    []string{"Times212_Output_0"},
				Name:      "Times212",
				OpType:    "MatMul",
				Domain:    "",
				DocString: "",
			},
			{
				Input:     []string{"Times212_Output_0", "Parameter194"},
				Output:    []string{"Plus214_Output_0"},
				Name:      "Plus214",
				OpType:    "Add",
				Domain:    "",
				DocString: "",
			},
		},
		Name: "CNTKGraph",
		Initializer: []*ir.TensorProto{
			{
				Dims:      []int64{16, 4, 4, 10},
				DataType:  1,
				FloatData: parameter193,
				Name:      "Parameter193",
			},
			{
				Dims:      []int64{16, 8, 5, 5},
				DataType:  1,
				FloatData: parameter87,
				Name:      "Parameter87",
			},
			{
				Dims:      []int64{8, 1, 5, 5},
				DataType:  1,
				FloatData: parameter5,
				Name:      "Parameter5",
			},
			{
				Dims:      []int64{8, 1, 1},
				DataType:  1,
				FloatData: parameter6,
				Name:      "Parameter6",
			},
			{
				Dims:      []int64{16, 1, 1},
				DataType:  1,
				FloatData: parameter88,
				Name:      "Parameter88",
			},
			{
				Dims:      []int64{2},
				DataType:  7,
				Int64Data: []int64{1, 256},
				Name:      "Pooling160_Output_0_reshape0_shape",
			},
			{
				Dims:      []int64{2},
				DataType:  7,
				Int64Data: []int64{256, 10},
				Name:      "Parameter193_reshape1_shape",
			},
			{
				Dims:      []int64{1, 10},
				DataType:  1,
				FloatData: parameter194,
				Name:      "Parameter194",
			},
		},
		Input: []*ir.ValueInfoProto{
			{
				Name: "Input3",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 28},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 28},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "Parameter5",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 8},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 5},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 5},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "Parameter6",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 8},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "Parameter87",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 16},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 8},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 5},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 5},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "Parameter88",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 16},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "Pooling160_Output_0_reshape0_shape",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 7,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 2},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "Parameter193",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 16},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 4},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 4},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 10},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "Parameter193_reshape1_shape",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 7,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 2},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "Parameter194",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 10},
									},
								},
							},
						},
					},
				},
			},
		},
		Output: []*ir.ValueInfoProto{
			{
				Name: "Plus214_Output_0",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 10},
									},
								},
							},
						},
					},
				},
			},
		},
		ValueInfo: []*ir.ValueInfoProto{
			{
				Name: "Convolution28_Output_0",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 8},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 28},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 28},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "Plus30_Output_0",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 8},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 28},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 28},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "ReLU32_Output_0",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 8},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 28},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 28},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "Pooling66_Output_0",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 8},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 14},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 14},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "Convolution110_Output_0",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 16},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 14},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 14},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "Plus112_Output_0",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 16},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 14},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 14},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "ReLU114_Output_0",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 16},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 14},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 14},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "Pooling160_Output_0",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 16},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 4},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 4},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "Pooling160_Output_0_reshape0",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 256},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "Parameter193_reshape1",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 256},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 10},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "Times212_Output_0",
				Type: &ir.TypeProto{
					Value: &ir.TypeProto_TensorType{
						TensorType: &ir.TypeProto_Tensor{
							ElemType: 1,
							Shape: &ir.TensorShapeProto{
								Dim: []*ir.TensorShapeProto_Dimension{
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									{
										Value: &ir.TensorShapeProto_Dimension_DimValue{DimValue: 10},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}
