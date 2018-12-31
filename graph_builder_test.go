package onnx

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/simple"
)

type simpleGraph struct {
	g *simple.DirectedGraph
}

type node struct {
	id          int64
	name        string
	description string
}

func (n *node) ID() int64 {
	return n.id
}

func (n *node) Attributes() []encoding.Attribute {
	return []encoding.Attribute{
		encoding.Attribute{
			Key:   "label",
			Value: fmt.Sprintf(`"%v"`, n.name),
		},
	}
}
func (n *node) SetDoc(name string) {
	n.description = name
}
func (n *node) SetName(desc string) {
	n.name = desc
}

func newSimpleGraph() *simpleGraph {
	return &simpleGraph{
		g: simple.NewDirectedGraph(),
	}
}

func (g *simpleGraph) SetEdge(e graph.Edge) {
	g.g.SetEdge(e)

}
func (g *simpleGraph) NewEdge(from, to graph.Node) graph.Edge {
	return g.g.NewEdge(from, to)

}
func (g *simpleGraph) AddNode(n graph.Node) {
	g.g.AddNode(n)

}
func (g *simpleGraph) NewNode() graph.Node {
	n := g.g.NewNode()
	return &node{
		id: n.ID(),
	}
}

func (g *simpleGraph) Node(id int64) graph.Node {
	return g.g.Node(id)
}
func (g *simpleGraph) Nodes() graph.Nodes {
	return g.g.Nodes()
}
func (g *simpleGraph) From(id int64) graph.Nodes {
	return g.g.From(id)
}
func (g *simpleGraph) HasEdgeBetween(xid, yid int64) bool {
	return g.g.HasEdgeBetween(xid, yid)
}
func (g *simpleGraph) Edge(uid, vid int64) graph.Edge {
	return g.g.Edge(uid, vid)
}
func (g *simpleGraph) HasEdgeFromTo(uid, vid int64) bool {
	return g.g.HasEdgeFromTo(uid, vid)
}
func (g *simpleGraph) To(id int64) graph.Nodes {
	return g.g.To(id)
}
