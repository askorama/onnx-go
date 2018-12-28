package onnx

import (
	"fmt"

	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"gonum.org/v1/gonum/graph"
)

// Attribute ...
type Attribute struct {
	Name  string
	Value interface{}
}

// NodeOp ...
type NodeOp interface {
	SetName(string)
	SetDoc(string)
	SetType(string)
	SetAttributes([]Attribute) error
}

// Unmarshal onnx encoded model proto data into a graph builder
func Unmarshal(data []byte, dst graph.DirectedBuilder) error {
	model := new(pb.ModelProto)
	err := model.Unmarshal(data)
	if err != nil {
		return err
	}
	db := make(map[string]graph.Node)
	for _, input := range model.Graph.Input {
		n := dst.NewNode()
		db[*input.Name] = n
		dst.AddNode(n)
	}
	for _, output := range model.Graph.Output {
		n := dst.NewNode()
		db[*output.Name] = n
		dst.AddNode(n)
	}
	for _, node := range model.Graph.Node {
		for _, input := range node.Input {
			for _, output := range node.Output {
				var ni, no graph.Node
				var ok bool
				if ni, ok = db[input]; !ok {
					return fmt.Errorf("Node %v not defined in the input", input)
				}
				if no, ok = db[output]; !ok {
					return fmt.Errorf("Node %v not defined in the input", output)
				}
				e := dst.NewEdge(ni, no)
				dst.SetEdge(e)
			}
		}
	}

	return err
}
