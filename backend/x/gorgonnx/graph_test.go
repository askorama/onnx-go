package gorgonnx

import (
	"testing"

	"github.com/owulveryck/onnx-go"
	"gorgonia.org/tensor"
)

func TestGraph_badnode(t *testing.T) {
	inputT := tensor.New(
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
	s := &stableSoftmax{
		axis: 1,
	}
	s.apply(g, output.(*Node))
	err := g.Run()
	if err == nil {
		t.Fatal("should raise an error because output is not a tensor nor an operation")
	}
}

func TestGetExprGraph(t *testing.T) {
	inputT := tensor.New(
		tensor.WithShape(2, 4),
		tensor.WithBacking([]float32{0, 1, 2, 3, 10000, 10001, 10002, 10003}),
	)
	inputT2 := tensor.New(
		tensor.WithShape(2, 4),
		tensor.WithBacking([]float32{0, 1, 2, 3, 10000, 10001, 10002, 10003}),
	)
	g := NewGraph()
	input := g.NewNode()
	g.AddNode(input)
	input2 := g.NewNode()
	g.AddNode(input2)
	output := g.NewNode()
	output.(*Node).operation = &onnx.Operation{
		Name: "Add",
	}
	g.AddNode(output)
	g.SetWeightedEdge(g.NewWeightedEdge(output, input, 0))
	g.SetWeightedEdge(g.NewWeightedEdge(output, input2, 1))
	input.(*Node).SetTensor(inputT)
	input2.(*Node).SetTensor(inputT2)
	var err error
	_, err = g.GetExprGraph()
	if err == nil {
		t.Fail()
	}

	g.exprgraph = nil
	g.ApplyOperation(onnx.Operation{
		Name: "Add",
	}, output.(*Node))
	expg, err := g.GetExprGraph()
	if err != nil {
		t.Fatal(err)
	}

	if len(expg.AllNodes()) != 3 {
		t.Log(expg)
		t.Fatalf("graph has %v node (expected %v)", len(expg.AllNodes()), 3)
	}
}
