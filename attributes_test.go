package onnx

import (
	"testing"
	"testify"

	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
)

/*
	AttributeProto_UNDEFINED AttributeProto_AttributeType = 0
	AttributeProto_FLOAT     AttributeProto_AttributeType = 1
	AttributeProto_INT       AttributeProto_AttributeType = 2
	AttributeProto_STRING    AttributeProto_AttributeType = 3
	AttributeProto_TENSOR    AttributeProto_AttributeType = 4
	AttributeProto_GRAPH     AttributeProto_AttributeType = 5
	AttributeProto_FLOATS    AttributeProto_AttributeType = 6
	AttributeProto_INTS      AttributeProto_AttributeType = 7
	AttributeProto_STRINGS   AttributeProto_AttributeType = 8
	AttributeProto_TENSORS   AttributeProto_AttributeType = 9
	AttributeProto_GRAPHS    AttributeProto_AttributeType = 10
*/
func GetTestPBAttributeProto() []*pb.AttributeProto {
	return []*pb.AttributeProto{
		{
			Name:   "floats",
			Type:   pb.AttributeProto_FLOATS,
			Floats: []float32{1, 2},
		},
		{
			Name: "float",
			Type: pb.AttributeProto_FLOAT,
			F:    1,
		},
		{
			Name: "int",
			Type: pb.AttributeProto_INT,
			I:    1,
		},
		{
			Name: "ints",
			Type: pb.AttributeProto_INTS,
			Ints: []int64{1, 2},
		},
		{
			Name: "string",
			Type: pb.AttributeProto_STRING,
			S:    []byte("a"),
		},
		{
			Name:    "strings",
			Type:    pb.AttributeProto_STRINGS,
			Strings: [][]byte{[]byte("a"), []byte("b")},
		},
	}
}


func TestToOperationAttributes_Strings(t *testing.T) {
	attrs, err := toOperationAttributes(GetTestPBAttributeProto())
	assert.NoError(err)
	
	v, ok := attrs["strings"]
	assert.True(t, ok)
	value, ok := v.([]string)
	assert.True(t, ok)
	expected := []string{"a", "b"}
	for i, v := range value {
		assert.Equals(t, expected[i], v)
	}
}

func TestToOperationAttributes_Ints(t *testing.T) {
	attrs, err := toOperationAttributes(GetTestPBAttributeProto())
	assert.NoError(err)
	
	v, ok := attrs["ints"]
	assert.True(t, ok)
	value, ok := v.([]int64)
	assert.True(t, ok)
	expected := []int64{1, 2}
	for i, v := range value {
		assert.Equals(t, expected[i], v)
	}
}

func TestToOperationAttributes_Floats(t *testing.T) {	
	attrs, err := toOperationAttributes(GetTestPBAttributeProto())
	assert.NoError(err)
	
	v, ok := attrs["floats"]
	assert.True(t, ok)
	value, ok := v.([]float32)
	assert.True(t, ok)
	expected := []float32{1, 2}
	for i, v := range value {
		assert.Equals(t, expected[i], v)
	}
}

func TestToOperationAttributes_String(t *testing.T) {	
	attrs, err := toOperationAttributes(GetTestPBAttributeProto())
	assert.NoError(err)
	
	v, ok := attrs["string"]
	assert.True(t, ok)
	assert.Equal(t, v.(string), "a")
}

func TestToOperationAttributes_Int(t *testing.T) {	
	attrs, err := toOperationAttributes(GetTestPBAttributeProto())
	assert.NoError(err)

	v, ok := attrs["int"]
	assert.True(t, ok)
	assert.Equal(t, v.(int64), int64(1))
}

func TestToOperationAttributes_Float(t *testing.T) {
	attrs, err := toOperationAttributes(GetTestPBAttributeProto())
	assert.NoError(err)
	
	v, ok := attrs["float"]
	assert.True(t, ok)
	assert.Equal(t, v.(float32), float32(1))
}

func TestToOperationAttributes_NotImplemented(t *testing.T) {
	_, err := toOperationAttributes([]*pb.AttributeProto{
		{
			Type: pb.AttributeProto_GRAPH,
		},
	})
	_, ok := err.(*ErrNotImplemented)
	assert.True(t, ok)
	_, err = toOperationAttributes([]*pb.AttributeProto{
		{
			Type: pb.AttributeProto_TENSORS,
		},
	})
	_, ok = err.(*ErrNotImplemented)
	assert.True(t, ok)
	_, err = toOperationAttributes([]*pb.AttributeProto{
		{
			Type: pb.AttributeProto_GRAPHS,
		},
	})
	_, ok = err.(*ErrNotImplemented)
	assert.True(t, ok)
	_, err = toOperationAttributes([]*pb.AttributeProto{
		{
			Type: pb.AttributeProto_AttributeType(-1),
		},
	})
	_, ok = err.(*ErrNotImplemented)
	assert.True(t, ok)
}
