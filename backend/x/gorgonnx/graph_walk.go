package gorgonnx

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/owulveryck/onnx-go"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/traverse"
	"gorgonia.org/gorgonia"
)

// populateExprgraph by walking through the graph
func (g *Graph) populateExprgraph() error {
	reverseGraph := simple.NewDirectedGraph()
	// Walk the graph
	itN := g.Nodes()
	for itN.Next() {
		reverseGraph.AddNode(itN.Node())
	}
	itE := g.g.Edges()
	for itE.Next() {
		reverseGraph.SetEdge(itE.Edge().ReversedEdge())
	}
	roots := make([]int64, 0)
	it := reverseGraph.Nodes()
	for it.Next() {
		n := it.Node()
		if reverseGraph.To(n.ID()).Len() == 0 {
			roots = append(roots, n.ID())
		}
	}
	var errorApply error
	bf := traverse.BreadthFirst{
		Visit: func(nde graph.Node) {
			n := nde.(*Node)
			if n.t != nil && n.gorgoniaNode == nil && n.operation == nil {
				n.gorgoniaNode = gorgonia.NodeFromAny(g.exprgraph, n.t, gorgonia.WithName(uuid.New().String()))
			}
			if n.operation != nil {
				children := getOrderedChildren(g.g, n)
				for i := 0; i < len(children); i++ {
					if children[i].gorgoniaNode == nil {
						errorApply = fmt.Errorf("will not apply operation on %v because its %vth child is nil (%v)", n, i, children[i])
						return
					}
				}
				var err error
				err = g.applyOperation(n)
				if err != nil {
					errorApply = err
					//log.Printf("ERROR: Cannot apply operation on node %v (%v)", n, err)
					return
				}
			}

		},
	}
	for i := 0; i < len(roots); i++ {
		bf.Reset()
		errorApply = nil
		bf.Walk(reverseGraph, reverseGraph.Node(roots[i]), nil)
	}
	return errorApply
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
