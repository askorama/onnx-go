package onnx

import "testing"

func TestUnmarshalAttributes(t *testing.T) {
	err := UnmarshalAttributes(nil, nil)
	if err == nil {
		t.Fail()
	}
	_, ok := err.(*InvalidUnmarshalError)
	if !ok {
		t.Fatal("Invalid error type")
	}
}
