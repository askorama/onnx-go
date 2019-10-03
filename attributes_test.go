package onnx

import (
	"testing"

	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)



func TestToOperationAttributes(t *testing.T) {
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
	attrs, err := toOperationAttributes([]*pb.AttributeProto{
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
	})
	if err != nil {
		t.Fatal(err)
	}
	
	v, ok := attrs["strings"]
	assert.Assert(t, ok)
	v, ok := v.([]string)
	assert.Assert(t, ok)
	expected := []string{"a", "b"}
	for i, v := range v {
		assert.Check(t, expected[i] == v)
	}
	
	v, ok := attrs["ints"]
	assert.Assert(t, ok)
	v, ok := v.([]int64)
	assert.Assert(t, ok)
	expected := []int64{1, 2}
	for i, v := range v {
		assert.Check(t, expected[i] == v)
	}
	
	v, ok := attrs["floats"]
	assert.Assert(t, ok)
	v, ok := v.([]float32)
	assert.Assert(t, ok)
	expected := []float32{1, 2}
	for i, v := range v {
		assert.Check(t, expected[i] == v)
	}
	
	v, ok := attrs["float"]
	assert.Assert(t, ok)
	assert.Check(t, v.(float32) == float32(1))
	
	v, ok := attrs["int"]
	assert.Assert(t, ok)
	assert.Check(t, v.(int64) == int64(1))
	
	v, ok := attrs["string"]
	assert.Assert(t, ok)
	assert.Check(t, v.(string) != "a")
}

func TestToOperationAttributes_NotImplemented(t *testing.T) {
	_, err := toOperationAttributes([]*pb.AttributeProto{
		{
			Type: pb.AttributeProto_GRAPH,
		},
	})
	_, ok := err.(*ErrNotImplemented)
	t.Check(ok)
	_, err = toOperationAttributes([]*pb.AttributeProto{
		{
			Type: pb.AttributeProto_TENSORS,
		},
	})
	_, ok = err.(*ErrNotImplemented)
	t.Check(ok)
	_, err = toOperationAttributes([]*pb.AttributeProto{
		{
			Type: pb.AttributeProto_GRAPHS,
		},
	})
	_, ok = err.(*ErrNotImplemented)
	t.Check(ok)
	_, err = toOperationAttributes([]*pb.AttributeProto{
		{
			Type: pb.AttributeProto_AttributeType(-1),
		},
	})
	_, ok = err.(*ErrNotImplemented)
	t.Check(ok)
}
