package onnx_test

import (
	"testing"

	onnx "github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/simple"
)

func TestUnmarshal(t *testing.T) {
	graph := simple.NewSimpleGraph()
	//graph := simple.NewDirectedGraph()
	err := onnx.Unmarshal(sigmoidNeuronONNX, graph)
	if err != nil {
		t.Fatal(err)
	}

}
