package gorgonnx

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/iterator"
)

// getOrderedChildren returns the children nodes of the current node
func getOrderedChildren(g graph.WeightedDirected, n *Node) []*Node {
	// Get all the edges that reach the node n
	children := g.From(n.ID())
	// Now get the edges
	edges := make([]graph.WeightedEdge, children.Len())
	for i := 0; children.Next(); i++ {
		edges[i] = g.WeightedEdge(n.ID(), children.Node().ID())
	}

	orderWeightedEdges := iterator.NewOrderedWeightedEdges(edges)
	nodes := make([]*Node, orderWeightedEdges.Len())
	for i := 0; orderWeightedEdges.Next(); i++ {
		nodes[i] = orderWeightedEdges.WeightedEdge().To().(*Node)
	}
	return nodes
}
