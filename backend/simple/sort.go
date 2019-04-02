package simple

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/iterator"
)

// GetOrderedChildren returns an iterator of children nodes ordered by the weighted edges
func GetOrderedChildren(g graph.WeightedDirected, n graph.Node) *iterator.OrderedNodes {
	// Get all the edges that reach the node n
	children := g.From(n.ID())
	// Now get the edges
	edges := make([]graph.WeightedEdge, children.Len())
	for i := 0; children.Next(); i++ {
		edges[i] = g.WeightedEdge(n.ID(), children.Node().ID())
	}

	orderWeightedEdges := iterator.NewOrderedWeightedEdges(edges)
	nodes := make([]graph.Node, children.Len())
	for i := 0; orderWeightedEdges.Next(); i++ {
		nodes[i] = orderWeightedEdges.WeightedEdge().To()
	}
	return iterator.NewOrderedNodes(nodes)
}
