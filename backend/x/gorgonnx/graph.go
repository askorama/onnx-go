package gorgonnx

import (
	"errors"

	"github.com/owulveryck/onnx-go"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gorgonia.org/gorgonia"

	xvm "gorgonia.org/gorgonia/x/vm"
	"gorgonia.org/tensor"
)

// Graph is the top structure that should be compatible with
//    backend.ComputationGraph
// It holds a gorgonia.ExprGraph that is populated on the first call to the
// Run() method
type Graph struct {
	g         *simple.WeightedDirectedGraph
	exprgraph *gorgonia.ExprGraph
	m         gorgonia.VM
	roots     []int64
	groups    [][]*Node // a reference of all the nodes that belongs to a group
}

// SetVM used by the backend
// A call to this method do not call the PopulateExprgraph method
// it is the responsibility of the caller to call it before
func (g *Graph) SetVM(vm gorgonia.VM) {
	g.m = vm
}

// GetExprGraph returns the gorgonia graph; if the graph is nil, it populates the graph before returing it
func (g *Graph) GetExprGraph() (*gorgonia.ExprGraph, error) {
	var err error
	if g.exprgraph == nil {
		err = g.PopulateExprgraph()
	}
	return g.exprgraph, err
}

// ApplyOperation to fulfill the onnx.Backend interface
func (g *Graph) ApplyOperation(o onnx.Operation, ns ...graph.Node) error {
	nodes := make([]*Node, len(ns))
	for i, n := range ns {
		n.(*Node).operation = &o
		nodes[i] = n.(*Node)
	}
	g.groups = append(g.groups, nodes)
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
	if g.m == nil {
		//g.m = gorgonia.NewTapeMachine(g.exprgraph)
		g.m = xvm.NewGoMachine(g.exprgraph)
	}

	err := g.m.RunAll()
	if err != nil {
		return err
	}
	// Now sets the output tensor
	for i := 0; i < len(g.roots); i++ {
		root := g.Node(g.roots[i]).(*Node)
		var ok bool
		if root.gorgoniaNode == nil {
			return errors.New("root node is nil")
		}
		root.t, ok = root.gorgoniaNode.Value().(tensor.Tensor)
		if !ok {
			return errors.New("root node is not a tensor")
		}
	}
	return nil
}

// PopulateExprgraph creates the underlynig graph by walking the current graph
func (g *Graph) PopulateExprgraph() error {
	g.exprgraph = gorgonia.NewGraph()
	// Find the root nodes
	// TODO make it more efficient
	g.roots = make([]int64, 0)
	it := g.g.Nodes()
	for it.Next() {
		n := it.Node()
		if g.g.To(n.ID()).Len() == 0 {
			g.roots = append(g.roots, n.ID())
		}
	}
	return g.populateExprgraph()
}
