package onnx

import pb "github.com/owulveryck/onnx-go/internal/pb-onnx"

func pInt64(v int64) int64 {
	return v
}

func pString(v string) string {
	return v
}

func pAttributeProtoAttributeType(i int32) pb.AttributeProto_AttributeType {
	v := pb.AttributeProto_AttributeType(i)
	return v
}

func pTensorProtoDataType(i int32) int32 {
	//v := pb.TensorProto_DataType(i)
	return i
}

var mnist = &pb.ModelProto{
	IrVersion: pInt64(3),
	OpsetImport: []*pb.OperatorSetIdProto{
		&pb.OperatorSetIdProto{
			Domain:  pString(""),
			Version: pInt64(7),
		},
	},
	ProducerName:    pString("CNTK"),
	ProducerVersion: pString("2.5.1"),
	Domain:          pString("ai.cntk"),
	ModelVersion:    pInt64(1),
	Graph: &pb.GraphProto{
		Node: []*pb.NodeProto{
			&pb.NodeProto{
				Input:     []string{"Parameter193", "Parameter193_reshape1_shape"},
				Output:    []string{"Parameter193_reshape1"},
				Name:      pString("Times212_reshape1"),
				OpType:    pString("Reshape"),
				Domain:    pString(""),
				DocString: pString(""),
			},
			&pb.NodeProto{
				Input:  []string{"Input3", "Parameter5"},
				Output: []string{"Convolution28_Output_0"},
				Name:   pString("Convolution28"),
				OpType: pString("Conv"),
				Domain: pString(""),
				Attribute: []*pb.AttributeProto{
					&pb.AttributeProto{
						Name: pString("kernel_shape"),
						Type: pAttributeProtoAttributeType(7),
						Ints: []int64{5, 5},
					},
					&pb.AttributeProto{
						Name: pString("strides"),
						Type: pAttributeProtoAttributeType(7),
						Ints: []int64{1, 1},
					},
					&pb.AttributeProto{
						Name: pString("auto_pad"),
						Type: pAttributeProtoAttributeType(3),
						S:    []byte{0x53, 0x41, 0x4d, 0x45, 0x5f, 0x55, 0x50, 0x50, 0x45, 0x52}, // SAME_UPPER
					},
					&pb.AttributeProto{
						Name: pString("group"),
						Type: pAttributeProtoAttributeType(2),
						I:    pInt64(1),
					},
					&pb.AttributeProto{
						Name: pString("dilations"),
						Type: pAttributeProtoAttributeType(7),
						Ints: []int64{1, 1},
					},
				},
				DocString: pString(""),
			},
			&pb.NodeProto{
				Input:     []string{"Convolution28_Output_0", "Parameter6"},
				Output:    []string{"Plus30_Output_0"},
				Name:      pString("Plus30"),
				OpType:    pString("Add"),
				Domain:    pString(""),
				DocString: pString(""),
			},
			&pb.NodeProto{
				Input:     []string{"Plus30_Output_0"},
				Output:    []string{"ReLU32_Output_0"},
				Name:      pString("ReLU32"),
				OpType:    pString("Relu"),
				Domain:    pString(""),
				DocString: pString(""),
			},
			&pb.NodeProto{
				Input:  []string{"ReLU32_Output_0"},
				Output: []string{"Pooling66_Output_0"},
				Name:   pString("Pooling66"),
				OpType: pString("MaxPool"),
				Domain: pString(""),
				Attribute: []*pb.AttributeProto{
					&pb.AttributeProto{
						Name: pString("kernel_shape"),
						Type: pAttributeProtoAttributeType(7),
						Ints: []int64{2, 2},
					},
					&pb.AttributeProto{
						Name: pString("strides"),
						Type: pAttributeProtoAttributeType(7),
						Ints: []int64{2, 2},
					},
					&pb.AttributeProto{
						Name: pString("pads"),
						Type: pAttributeProtoAttributeType(7),
						Ints: []int64{0, 0, 0, 0},
					},
					&pb.AttributeProto{
						Name: pString("auto_pad"),
						Type: pAttributeProtoAttributeType(3),
						S:    []byte{0x4e, 0x4f, 0x54, 0x53, 0x45, 0x54}, // NOTSET
					},
				},
				DocString: pString(""),
			},
			&pb.NodeProto{
				Input:  []string{"Pooling66_Output_0", "Parameter87"},
				Output: []string{"Convolution110_Output_0"},
				Name:   pString("Convolution110"),
				OpType: pString("Conv"),
				Domain: pString(""),
				Attribute: []*pb.AttributeProto{
					&pb.AttributeProto{
						Name: pString("kernel_shape"),
						Type: pAttributeProtoAttributeType(7),
						Ints: []int64{5, 5},
					},
					&pb.AttributeProto{
						Name: pString("strides"),
						Type: pAttributeProtoAttributeType(7),
						Ints: []int64{1, 1},
					},
					&pb.AttributeProto{
						Name: pString("auto_pad"),
						Type: pAttributeProtoAttributeType(3),
						S:    []byte{0x53, 0x41, 0x4d, 0x45, 0x5f, 0x55, 0x50, 0x50, 0x45, 0x52}, // SAME_UPPER
					},
					&pb.AttributeProto{
						Name: pString("group"),
						Type: pAttributeProtoAttributeType(2),
						I:    pInt64(1),
					},
					&pb.AttributeProto{
						Name: pString("dilations"),
						Type: pAttributeProtoAttributeType(7),
						Ints: []int64{1, 1},
					},
				},
				DocString: pString(""),
			},
			&pb.NodeProto{
				Input:     []string{"Convolution110_Output_0", "Parameter88"},
				Output:    []string{"Plus112_Output_0"},
				Name:      pString("Plus112"),
				OpType:    pString("Add"),
				Domain:    pString(""),
				DocString: pString(""),
			},
			&pb.NodeProto{
				Input:     []string{"Plus112_Output_0"},
				Output:    []string{"ReLU114_Output_0"},
				Name:      pString("ReLU114"),
				OpType:    pString("Relu"),
				Domain:    pString(""),
				DocString: pString(""),
			},
			&pb.NodeProto{
				Input:  []string{"ReLU114_Output_0"},
				Output: []string{"Pooling160_Output_0"},
				Name:   pString("Pooling160"),
				OpType: pString("MaxPool"),
				Domain: pString(""),
				Attribute: []*pb.AttributeProto{
					&pb.AttributeProto{
						Name: pString("kernel_shape"),
						Type: pAttributeProtoAttributeType(7),
						Ints: []int64{3, 3},
					},
					&pb.AttributeProto{
						Name: pString("strides"),
						Type: pAttributeProtoAttributeType(7),
						Ints: []int64{3, 3},
					},
					&pb.AttributeProto{
						Name: pString("pads"),
						Type: pAttributeProtoAttributeType(7),
						Ints: []int64{0, 0, 0, 0},
					},
					&pb.AttributeProto{
						Name: pString("auto_pad"),
						Type: pAttributeProtoAttributeType(3),
						S:    []byte{0x4e, 0x4f, 0x54, 0x53, 0x45, 0x54}, // NOTSET
					},
				},
				DocString: pString(""),
			},
			&pb.NodeProto{
				Input:     []string{"Pooling160_Output_0", "Pooling160_Output_0_reshape0_shape"},
				Output:    []string{"Pooling160_Output_0_reshape0"},
				Name:      pString("Times212_reshape0"),
				OpType:    pString("Reshape"),
				Domain:    pString(""),
				DocString: pString(""),
			},
			&pb.NodeProto{
				Input:     []string{"Pooling160_Output_0_reshape0", "Parameter193_reshape1"},
				Output:    []string{"Times212_Output_0"},
				Name:      pString("Times212"),
				OpType:    pString("MatMul"),
				Domain:    pString(""),
				DocString: pString(""),
			},
			&pb.NodeProto{
				Input:     []string{"Times212_Output_0", "Parameter194"},
				Output:    []string{"Plus214_Output_0"},
				Name:      pString("Plus214"),
				OpType:    pString("Add"),
				Domain:    pString(""),
				DocString: pString(""),
			},
		},
		Name: pString("CNTKGraph"),
		Initializer: []*pb.TensorProto{
			&pb.TensorProto{
				Dims:      []int64{16, 4, 4, 10},
				DataType:  pTensorProtoDataType(1),
				FloatData: parameter193,
				Name:      pString("Parameter193"),
			},
			&pb.TensorProto{
				Dims:      []int64{16, 8, 5, 5},
				DataType:  pTensorProtoDataType(1),
				FloatData: parameter87,
				Name:      pString("Parameter87"),
			},
			&pb.TensorProto{
				Dims:      []int64{8, 1, 5, 5},
				DataType:  pTensorProtoDataType(1),
				FloatData: parameter5,
				Name:      pString("Parameter5"),
			},
			&pb.TensorProto{
				Dims:      []int64{8, 1, 1},
				DataType:  pTensorProtoDataType(1),
				FloatData: parameter6,
				Name:      pString("Parameter6"),
			},
			&pb.TensorProto{
				Dims:      []int64{16, 1, 1},
				DataType:  pTensorProtoDataType(1),
				FloatData: parameter88,
				Name:      pString("Parameter88"),
			},
			&pb.TensorProto{
				Dims:      []int64{2},
				DataType:  pTensorProtoDataType(7),
				Int64Data: []int64{1, 256},
				Name:      pString("Pooling160_Output_0_reshape0_shape"),
			},
			&pb.TensorProto{
				Dims:      []int64{2},
				DataType:  pTensorProtoDataType(7),
				Int64Data: []int64{256, 10},
				Name:      pString("Parameter193_reshape1_shape"),
			},
			&pb.TensorProto{
				Dims:      []int64{1, 10},
				DataType:  pTensorProtoDataType(1),
				FloatData: parameter194,
				Name:      pString("Parameter194"),
			},
		},
		Input: []*pb.ValueInfoProto{
			&pb.ValueInfoProto{
				Name: pString("Input3"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 28},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 28},
									},
								},
							},
						},
					},
				},
			},
			&pb.ValueInfoProto{
				Name: pString("Parameter5"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 8},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 5},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 5},
									},
								},
							},
						},
					},
				},
			},
			&pb.ValueInfoProto{
				Name: pString("Parameter6"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 8},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
								},
							},
						},
					},
				},
			},
			&pb.ValueInfoProto{
				Name: pString("Parameter87"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 16},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 8},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 5},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 5},
									},
								},
							},
						},
					},
				},
			},
			&pb.ValueInfoProto{
				Name: pString("Parameter88"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 16},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
								},
							},
						},
					},
				},
			},
			&pb.ValueInfoProto{
				Name: pString("Pooling160_Output_0_reshape0_shape"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(7),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 2},
									},
								},
							},
						},
					},
				},
			},
			&pb.ValueInfoProto{
				Name: pString("Parameter193"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 16},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 4},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 4},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 10},
									},
								},
							},
						},
					},
				},
			},
			&pb.ValueInfoProto{
				Name: pString("Parameter193_reshape1_shape"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(7),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 2},
									},
								},
							},
						},
					},
				},
			},
			&pb.ValueInfoProto{
				Name: pString("Parameter194"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 10},
									},
								},
							},
						},
					},
				},
			},
		},
		Output: []*pb.ValueInfoProto{
			&pb.ValueInfoProto{
				Name: pString("Plus214_Output_0"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 10},
									},
								},
							},
						},
					},
				},
			},
		},
		ValueInfo: []*pb.ValueInfoProto{
			&pb.ValueInfoProto{
				Name: pString("Convolution28_Output_0"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 8},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 28},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 28},
									},
								},
							},
						},
					},
				},
			},
			&pb.ValueInfoProto{
				Name: pString("Plus30_Output_0"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 8},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 28},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 28},
									},
								},
							},
						},
					},
				},
			},
			&pb.ValueInfoProto{
				Name: pString("ReLU32_Output_0"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 8},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 28},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 28},
									},
								},
							},
						},
					},
				},
			},
			&pb.ValueInfoProto{
				Name: pString("Pooling66_Output_0"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 8},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 14},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 14},
									},
								},
							},
						},
					},
				},
			},
			&pb.ValueInfoProto{
				Name: pString("Convolution110_Output_0"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 16},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 14},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 14},
									},
								},
							},
						},
					},
				},
			},
			&pb.ValueInfoProto{
				Name: pString("Plus112_Output_0"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 16},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 14},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 14},
									},
								},
							},
						},
					},
				},
			},
			&pb.ValueInfoProto{
				Name: pString("ReLU114_Output_0"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 16},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 14},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 14},
									},
								},
							},
						},
					},
				},
			},
			&pb.ValueInfoProto{
				Name: pString("Pooling160_Output_0"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 16},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 4},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 4},
									},
								},
							},
						},
					},
				},
			},
			&pb.ValueInfoProto{
				Name: pString("Pooling160_Output_0_reshape0"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 256},
									},
								},
							},
						},
					},
				},
			},
			&pb.ValueInfoProto{
				Name: pString("Parameter193_reshape1"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 256},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 10},
									},
								},
							},
						},
					},
				},
			},
			&pb.ValueInfoProto{
				Name: pString("Times212_Output_0"),
				Type: &pb.TypeProto{
					Value: &pb.TypeProto_TensorType{
						TensorType: &pb.TypeProto_Tensor{
							ElemType: pTensorProtoDataType(1),
							Shape: &pb.TensorShapeProto{
								Dim: []*pb.TensorShapeProto_Dimension{
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
									},
									&pb.TensorShapeProto_Dimension{
										Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 10},
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
