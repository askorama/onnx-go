package onnx

import (
	"testing"

	"github.com/owulveryck/onnx-go/internal/onnx/ir"
	"github.com/stretchr/testify/assert"
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

func GetTestPBAttributeProto() []*ir.AttributeProto {
	return []*ir.AttributeProto{
		{
			Name:   "floats",
			Type:   ir.AttributeProto_FLOATS,
			Floats: []float32{1, 2},
		},
		{
			Name: "float",
			Type: ir.AttributeProto_FLOAT,
			F:    1,
		},
		{
			Name: "int",
			Type: ir.AttributeProto_INT,
			I:    1,
		},
		{
			Name: "ints",
			Type: ir.AttributeProto_INTS,
			Ints: []int64{1, 2},
		},
		{
			Name: "string",
			Type: ir.AttributeProto_STRING,
			S:    []byte("a"),
		},
		{
			Name:    "strings",
			Type:    ir.AttributeProto_STRINGS,
			Strings: [][]byte{[]byte("a"), []byte("b")},
		},
	}
}

func TestToOperationAttributes_Strings(t *testing.T) {
	attrs, err := toOperationAttributes(GetTestPBAttributeProto())
	assert.NoError(t, err)

	v, ok := attrs["strings"]
	assert.True(t, ok)
	value, ok := v.([]string)
	assert.True(t, ok)
	expected := []string{"a", "b"}
	for i, v := range value {
		assert.Equal(t, expected[i], v)
	}
}

func TestToOperationAttributes_Ints(t *testing.T) {
	attrs, err := toOperationAttributes(GetTestPBAttributeProto())
	assert.NoError(t, err)

	v, ok := attrs["ints"]
	assert.True(t, ok)
	value, ok := v.([]int64)
	assert.True(t, ok)
	expected := []int64{1, 2}
	for i, v := range value {
		assert.Equal(t, expected[i], v)
	}
}

func TestToOperationAttributes_Floats(t *testing.T) {
	attrs, err := toOperationAttributes(GetTestPBAttributeProto())
	assert.NoError(t, err)

	v, ok := attrs["floats"]
	assert.True(t, ok)
	value, ok := v.([]float32)
	assert.True(t, ok)
	expected := []float32{1, 2}
	for i, v := range value {
		assert.Equal(t, expected[i], v)
	}
}

func TestToOperationAttributes_String(t *testing.T) {
	attrs, err := toOperationAttributes(GetTestPBAttributeProto())
	assert.NoError(t, err)

	v, ok := attrs["string"]
	assert.True(t, ok)
	assert.Equal(t, v.(string), "a")
}

func TestToOperationAttributes_Int(t *testing.T) {
	attrs, err := toOperationAttributes(GetTestPBAttributeProto())
	assert.NoError(t, err)

	v, ok := attrs["int"]
	assert.True(t, ok)
	assert.Equal(t, v.(int64), int64(1))
}

func TestToOperationAttributes_Float(t *testing.T) {
	attrs, err := toOperationAttributes(GetTestPBAttributeProto())
	assert.NoError(t, err)

	v, ok := attrs["float"]
	assert.True(t, ok)
	assert.Equal(t, v.(float32), float32(1))
}

func TestToOperationAttributes_NotImplemented(t *testing.T) {
	_, err := toOperationAttributes([]*ir.AttributeProto{
		{
			Type: ir.AttributeProto_GRAPH,
		},
	})
	_, ok := err.(*ErrNotImplemented)
	assert.True(t, ok)
	_, err = toOperationAttributes([]*ir.AttributeProto{
		{
			Type: ir.AttributeProto_TENSORS,
		},
	})
	_, ok = err.(*ErrNotImplemented)
	assert.True(t, ok)
	_, err = toOperationAttributes([]*ir.AttributeProto{
		{
			Type: ir.AttributeProto_GRAPHS,
		},
	})
	_, ok = err.(*ErrNotImplemented)
	assert.True(t, ok)
	_, err = toOperationAttributes([]*ir.AttributeProto{
		{
			Type: ir.AttributeProto_AttributeType(-1),
		},
	})
	_, ok = err.(*ErrNotImplemented)
	assert.True(t, ok)
}

func TestToOperationAttributes_Undefined(t *testing.T) {
	attrs, err := toOperationAttributes([]*ir.AttributeProto{
		nil,
	})
	assert.NoError(t, err)

	v, ok := attrs[""]
	assert.True(t, ok)
	expected := struct{}{}
	assert.Equal(t, expected, v)
}
