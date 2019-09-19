package onnx

import "testing"

func TestContains(t *testing.T) {
	table := []string{"a", "b", "c"}
	ok := contains("a", table)
	if !ok {
		t.Fail()
	}
	ok = contains("z", table)
	if ok {
		t.Fail()
	}
	table = []string{"a", "a", "b", "c"}
	ok = contains("a", table)
	if !ok {
		t.Fail()
	}
}
