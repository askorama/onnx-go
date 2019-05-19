package onnx

import (
	"testing"

	"gorgonia.org/tensor"
)

func TestSetInput_nil_model(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log("panic expected")
		}
	}()
	m := new(Model)
	tens := tensor.New()
	err := m.SetInput(0, tens)
	t.Fatal("should have panicedm but have failed passed with error", err)
}
