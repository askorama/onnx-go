package onnx

// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"math"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/iterator"
	"gonum.org/v1/gonum/graph/simple"
)

const (
	self, absent = math.MaxFloat64, float64(-1)
)

// testWeightedDirectedGraph implements a generalized weighted directed graph.
type testWeightedDirectedGraph struct {
	nodes map[int64]*nodeTest
	from  map[int64]map[int64]graph.WeightedEdge
	to    map[int64]map[int64]graph.WeightedEdge
}

// Edge returns the edge from u to v if such an edge exists and nil otherwise.
// The node v must be directly reachable from u as defined by the From method.
func (g *testWeightedDirectedGraph) Edge(uid, vid int64) graph.Edge {
	return g.WeightedEdge(uid, vid)
}

// Edges returns all the edges in the graph.
func (g *testWeightedDirectedGraph) Edges() graph.Edges {
	var edges []graph.Edge
	for _, u := range g.nodes {
		for _, e := range g.from[u.ID()] {
			edges = append(edges, e)
		}
	}
	if len(edges) == 0 {
		return graph.Empty
	}
	return iterator.NewOrderedEdges(edges)
}

// From returns all nodes in g that can be reached directly from n.
func (g *testWeightedDirectedGraph) From(id int64) graph.Nodes {
	if _, ok := g.from[id]; !ok {
		return graph.Empty
	}

	from := make([]graph.Node, len(g.from[id]))
	i := 0
	for vid := range g.from[id] {
		from[i] = g.nodes[vid]
		i++
	}
	if len(from) == 0 {
		return graph.Empty
	}
	return iterator.NewOrderedNodes(from)
}

// HasEdgeBetween returns whether an edge exists between nodes x and y without
// considering direction.
func (g *testWeightedDirectedGraph) HasEdgeBetween(xid, yid int64) bool {
	if _, ok := g.from[xid][yid]; ok {
		return true
	}
	_, ok := g.from[yid][xid]
	return ok
}

// HasEdgeFromTo returns whether an edge exists in the graph from u to v.
func (g *testWeightedDirectedGraph) HasEdgeFromTo(uid, vid int64) bool {
	if _, ok := g.from[uid][vid]; !ok {
		return false
	}
	return true
}

// NewWeightedEdge returns a new weighted edge from the source to the destination node.
func (g *testWeightedDirectedGraph) NewWeightedEdge(from, to graph.Node, weight float64) graph.WeightedEdge {
	return &simple.WeightedEdge{F: from, T: to, W: weight}
}

// Node returns the node with the given ID if it exists in the graph,
// and nil otherwise.
func (g *testWeightedDirectedGraph) Node(id int64) graph.Node {
	return g.nodes[id]
}

// Nodes returns all the nodes in the graph.
func (g *testWeightedDirectedGraph) Nodes() graph.Nodes {
	if len(g.from) == 0 {
		return graph.Empty
	}
	nodes := make([]graph.Node, len(g.nodes))
	i := 0
	for _, n := range g.nodes {
		nodes[i] = n
		i++
	}
	return iterator.NewOrderedNodes(nodes)
}

// To returns all nodes in g that can reach directly to n.
func (g *testWeightedDirectedGraph) To(id int64) graph.Nodes {
	if _, ok := g.from[id]; !ok {
		return graph.Empty
	}

	to := make([]graph.Node, len(g.to[id]))
	i := 0
	for uid := range g.to[id] {
		to[i] = g.nodes[uid]
		i++
	}
	if len(to) == 0 {
		return graph.Empty
	}
	return iterator.NewOrderedNodes(to)
}

// Weight returns the weight for the edge between x and y if Edge(x, y) returns a non-nil Edge.
// If x and y are the same node or there is no joining edge between the two nodes the weight
// value returned is either the graph's absent or self value. Weight returns true if an edge
// exists between x and y or if x and y have the same ID, false otherwise.
func (g *testWeightedDirectedGraph) Weight(xid, yid int64) (w float64, ok bool) {
	if xid == yid {
		return self, true
	}
	if to, ok := g.from[xid]; ok {
		if e, ok := to[yid]; ok {
			return e.Weight(), true
		}
	}
	return absent, false
}

// WeightedEdge returns the weighted edge from u to v if such an edge exists and nil otherwise.
// The node v must be directly reachable from u as defined by the From method.
func (g *testWeightedDirectedGraph) WeightedEdge(uid, vid int64) graph.WeightedEdge {
	edge, ok := g.from[uid][vid]
	if !ok {
		return nil
	}
	return edge
}

// WeightedEdges returns all the weighted edges in the graph.
func (g *testWeightedDirectedGraph) WeightedEdges() graph.WeightedEdges {
	var edges []graph.WeightedEdge
	for _, u := range g.nodes {
		for _, e := range g.from[u.ID()] {
			edges = append(edges, e)
		}
	}
	if len(edges) == 0 {
		return graph.Empty
	}
	return iterator.NewOrderedWeightedEdges(edges)
}
