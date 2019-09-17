package gorgonnx

import (
	"testing"

	"github.com/owulveryck/onnx-go"
	"github.com/stretchr/testify/assert"
	"gorgonia.org/tensor"
)

func TestSqueeze_NoAxis(t *testing.T) {

	inputT := tensor.New(
		tensor.WithShape(2, 1, 1, 4),
		tensor.WithBacking([]float32{0, 1, 2, 3, 10000, 10001, 10002, 10003}),
	)
	expectedOutput := tensor.New(
		tensor.WithShape(2, 4),
		tensor.WithBacking([]float32{0, 1, 2, 3, 10000, 10001, 10002, 10003}),
	)
	g := NewGraph()
	input := g.NewNode()
	g.AddNode(input)
	output := g.NewNode()
	g.AddNode(output)
	g.SetWeightedEdge(g.NewWeightedEdge(output, input, 0))
	input.(*Node).SetTensor(inputT)
	g.ApplyOperation(onnx.Operation{
		Name:       "Squeeze",
		Attributes: nil,
	}, output)

	err := g.Run()
	if err != nil {
		t.Fatal(err)
	}

	outputT := output.(*Node).GetTensor()
	assert.InDeltaSlice(t, expectedOutput.Data(), outputT.Data(), 1e-6, "the two tensors should be equal.")
}

func TestSqueeze_Scalar(t *testing.T) {

	inputT := tensor.New(
		tensor.WithShape(1, 1, 1),
		tensor.WithBacking([]float32{3.14}),
	)
	expectedOutput := tensor.New(
		tensor.WithShape(1),
		tensor.WithBacking([]float32{3.14}),
	)
	g := NewGraph()
	input := g.NewNode()
	g.AddNode(input)
	output := g.NewNode()
	g.AddNode(output)
	g.SetWeightedEdge(g.NewWeightedEdge(output, input, 0))
	input.(*Node).SetTensor(inputT)
	g.ApplyOperation(onnx.Operation{
		Name: "Squeeze",
		Attributes: map[string]interface{}{
			"axes": nil,
		},
	}, output)

	err := g.Run()
	if err != nil {
		t.Fatal(err)
	}

	outputT := output.(*Node).GetTensor()
	assert.InDelta(t, expectedOutput.Data(), outputT.Data(), 1e-6, "the two tensors should be equal.")
}

func TestSqueeze_Axis(t *testing.T) {

	inputT := tensor.New(
		tensor.WithShape(1, 2, 1, 4, 1),
		tensor.WithBacking([]float32{0, 1, 2, 3, 10000, 10001, 10002, 10003}),
	)
	expectedOutput := tensor.New(
		tensor.WithShape(1, 2, 4),
		tensor.WithBacking([]float32{0, 1, 2, 3, 10000, 10001, 10002, 10003}),
	)
	g := NewGraph()
	input := g.NewNode()
	g.AddNode(input)
	output := g.NewNode()
	g.AddNode(output)
	g.SetWeightedEdge(g.NewWeightedEdge(output, input, 0))
	input.(*Node).SetTensor(inputT)
	g.ApplyOperation(onnx.Operation{
		Name: "Squeeze",
		Attributes: map[string]interface{}{
			"axes": []int64{4, 0},
		},
	}, output)

	err := g.Run()
	if err != nil {
		t.Fatal(err)
	}

	outputT := output.(*Node).GetTensor()
	assert.InDeltaSlice(t, expectedOutput.Data(), outputT.Data(), 1e-6, "the two tensors should be equal.")
}

func TestSqueeze_Fail(t *testing.T) {

	inputT := tensor.New(
		tensor.WithShape(2, 1, 1, 4),
		tensor.WithBacking([]float32{0, 1, 2, 3, 10000, 10001, 10002, 10003}),
	)
	g := NewGraph()
	input := g.NewNode()
	g.AddNode(input)
	output := g.NewNode()
	g.AddNode(output)
	g.SetWeightedEdge(g.NewWeightedEdge(output, input, 0))
	input.(*Node).SetTensor(inputT)
	g.ApplyOperation(onnx.Operation{
		Name: "Squeeze",
		Attributes: map[string]interface{}{
			"axes": []int64{0},
		},
	}, output)

	err := g.Run()
	assert.EqualError(t, err, "Unable to squeeze an axis whose shape entry is not 1 (got 2 instead)")
}
