package gorgonnx

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/owulveryck/onnx-go"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/traverse"
	"gorgonia.org/gorgonia"
)

// walk the graph from node "node"
func (g *Graph) walk(node int64) error {
	// n contains an ordered list of the nodes of the graph
	nodes := make([]int64, 0, g.Nodes().Len())
	// Walk the graph
	bf := traverse.BreadthFirst{
		Visit: func(v graph.Node) {
			if len(nodes) == 0 || nodes[len(nodes)-1] != v.ID() {
				nodes = append(nodes, v.ID())
			}
		},
	}

	bf.Walk(g, g.Node(node), nil)
	if len(nodes) == 0 {
		return errors.New("unable to compute node, empty path")
	}
	// for each node, if nil, and if hold an operation, add the graph
	for i := len(nodes) - 1; i >= 0; i-- {
		n := g.g.Node(nodes[i]).(*Node)
		if n.t == nil && n.operation == nil {
			return fmt.Errorf("node %v is not a tensor nor an operation", n)
		}
		if n.t != nil && n.gorgoniaNode == nil && n.operation == nil {
			n.gorgoniaNode = gorgonia.NodeFromAny(g.exprgraph, n.t, gorgonia.WithName(uuid.New().String()))
		}
		if n.operation != nil {
			var err error
			err = g.applyOperation(n)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// applyOperation creates a new node on the exprgraph
func (g *Graph) applyOperation(n *Node) error {
	// Is this node already in the ExprGraph?
	if n.gorgoniaNode != nil {
		return fmt.Errorf("unsupported case: node is already in the exprgraph")
	}
	var op operator
	var opC func() operator
	var ok bool
	if opC, ok = operators[n.operation.Name]; !ok {
		return &onnx.ErrNotImplemented{
			Operator: n.operation.Name,
		}
	}
	op = opC()
	err := op.init(*n.operation)
	if err != nil {
		return err
	}
	return op.apply(g, n)
}
