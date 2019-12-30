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

func TestGetInputTensors(t *testing.T) {
	backend := newTestBackend()
	n1 := backend.NewNode()
	n2 := backend.NewNode()
	model := &Model{
		Input:   []int64{n1.ID(), n2.ID()},
		backend: backend,
	}
	input := model.GetInputTensors()
	t.Log(input)
	if len(input) != 2 {
		t.FailNow()
	}
	if input[0] != nil || input[1] != nil {
		t.Fail()
	}
}
