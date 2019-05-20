package onnx

import (
	"testing"

	"gorgonia.org/tensor"
)

func TestSetInput_nil_model(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

		}
	}()
	m := new(Model)
	tens := tensor.New()
	err := m.SetInput(0, tens)
	t.Fatal("should have paniced but have passed with error", err)
}
