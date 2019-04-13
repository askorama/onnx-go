package gorgonnx

import (
	"github.com/owulveryck/onnx-go"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gorgonia.org/gorgonia"
)

// Graph is the top structure that should be compatible with
//    backend.ComputationGraph
// It holds a gorgonia.ExprGraph that is populated on the first call to the
// Run() method
type Graph struct {
	g         *simple.WeightedDirectedGraph
	exprgraph *gorgonia.ExprGraph
}

// ApplyOperation to fulfill the onnx.Backend interface
func (g *Graph) ApplyOperation(o onnx.Operation, n graph.Node) error {
	n.(*Node).operation = &o
	return nil
}

// Run the graph. It populate the underlying exprgraph if the graph is nil
func (g *Graph) Run() error {
	if g.exprgraph == nil {
		err := g.PopulateExprgraph()
		if err != nil {
			return err
		}
	}
	t := gorgonia.NewTapeMachine(g.exprgraph)
	err := t.RunAll()
	if err != nil {
		return err
	}
	// Now sets the output tensor
	// TODO
	return nil
}

// PopulateExprgraph creates the underlynig graph by walking the current graph
func (g *Graph) PopulateExprgraph() error {
	g.exprgraph = gorgonia.NewGraph()
	// Find the root nodes
	// TODO make it more efficient
	root := make([]int64, 0)
	it := g.g.Nodes()
	for it.Next() {
		n := it.Node()
		if g.g.To(n.ID()).Len() == 0 {
			root = append(root, n.ID())
		}
	}
	if len(root) != 1 {
		return &onnx.ErrNotImplemented{}
	}

	return g.walk(root[0])
}
