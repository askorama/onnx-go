package onnx

import (
	"testing"

	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
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
		&pb.AttributeProto{
			Name:   "floats",
			Type:   pb.AttributeProto_FLOATS,
			Floats: []float32{1, 2},
		},
		&pb.AttributeProto{
			Name: "float",
			Type: pb.AttributeProto_FLOAT,
			F:    1,
		},
		&pb.AttributeProto{
			Name: "int",
			Type: pb.AttributeProto_INT,
			I:    1,
		},
		&pb.AttributeProto{
			Name: "ints",
			Type: pb.AttributeProto_INTS,
			Ints: []int64{1, 2},
		},
		&pb.AttributeProto{
			Name: "string",
			Type: pb.AttributeProto_STRING,
			S:    []byte("a"),
		},
		&pb.AttributeProto{
			Name:    "strings",
			Type:    pb.AttributeProto_STRINGS,
			Strings: [][]byte{[]byte("a"), []byte("b")},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if v, ok := attrs["strings"]; ok {
		if v, ok := v.([]string); ok {
			expected := []string{"a", "b"}
			for i, v := range v {
				if expected[i] != v {
					t.Fail()
				}
			}
		} else {
			t.Fail()
		}
	} else {
		t.Fail()
	}
	if v, ok := attrs["ints"]; ok {
		if v, ok := v.([]int64); ok {
			expected := []int64{1, 2}
			for i, v := range v {
				if expected[i] != v {
					t.Fail()
				}
			}
		} else {
			t.Fail()
		}
	} else {
		t.Fail()
	}
	if v, ok := attrs["floats"]; ok {
		if v, ok := v.([]float32); ok {
			expected := []float32{1, 2}
			for i, v := range v {
				if expected[i] != v {
					t.Fail()
				}
			}
		} else {
			t.Fail()
		}
	} else {
		t.Fail()
	}
	if v, ok := attrs["float"]; ok {
		if v.(float32) != float32(1) {
			t.Fail()
		}
	} else {
		t.Fail()
	}
	if v, ok := attrs["int"]; ok {
		if v.(int64) != int64(1) {
			t.Fatal("bad value for int")
		}
	} else {
		t.Fail()
	}
	if v, ok := attrs["string"]; ok {
		if v.(string) != "a" {
			t.Fatal("bad value for sting")
		}
	} else {
		t.Fail()
	}
}

func TestToOperationAttributes_NotImplemented(t *testing.T) {
	_, err := toOperationAttributes([]*pb.AttributeProto{
		&pb.AttributeProto{
			Type: pb.AttributeProto_GRAPH,
		},
	})
	_, ok := err.(*ErrNotImplemented)
	if !ok {
		t.Fail()
	}
	_, err = toOperationAttributes([]*pb.AttributeProto{
		&pb.AttributeProto{
			Type: pb.AttributeProto_TENSORS,
		},
	})
	_, ok = err.(*ErrNotImplemented)
	if !ok {
		t.Fail()
	}
	_, err = toOperationAttributes([]*pb.AttributeProto{
		&pb.AttributeProto{
			Type: pb.AttributeProto_GRAPHS,
		},
	})
	_, ok = err.(*ErrNotImplemented)
	if !ok {
		t.Fail()
	}
	_, err = toOperationAttributes([]*pb.AttributeProto{
		&pb.AttributeProto{
			Type: pb.AttributeProto_AttributeType(-1),
		},
	})
	_, ok = err.(*ErrNotImplemented)
	if !ok {
		t.Fail()
	}
}
