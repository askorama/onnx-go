package gorgonnx

import (
	"testing"

	"github.com/owulveryck/onnx-go"
	"github.com/stretchr/testify/assert"
	"gorgonia.org/tensor"
)

func TestReshape_Scalar(t *testing.T) {

	inputT := tensor.New(
		tensor.WithShape(2, 1, 1, 4),
		tensor.WithBacking([]float32{0, 1, 2, 3, 10000, 10001, 10002, 10003}),
	)
	inputShape := tensor.New(
		tensor.WithShape(1),
		tensor.WithBacking([]int64{-1}),
	)

	expectedOutput := tensor.New(
		tensor.WithShape(8),
		tensor.WithBacking([]float32{0, 1, 2, 3, 10000, 10001, 10002, 10003}),
	)
	g := NewGraph()
	input1 := g.NewNode()
	g.AddNode(input1)
	input2 := g.NewNode()
	g.AddNode(input2)
	output := g.NewNode()
	g.AddNode(output)
	g.SetWeightedEdge(g.NewWeightedEdge(output, input1, 0))
	input1.(*Node).SetTensor(inputT)
	g.SetWeightedEdge(g.NewWeightedEdge(output, input2, 0))
	input2.(*Node).SetTensor(inputShape)
	g.ApplyOperation(onnx.Operation{
		Name: "Reshape",
	}, output)

	err := g.Run()
	if err != nil {
		t.Fatal(err)
	}

	outputT := output.(*Node).GetTensor()
	assert.InDeltaSlice(t, expectedOutput.Data(), outputT.Data(), 1e-6, "the two tensors should be equal.")
}
