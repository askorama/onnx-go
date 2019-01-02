package main

import (
	"fmt"
	"math"

	onnx "github.com/owulveryck/onnx-go"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/simple"
	"gorgonia.org/tensor"
)

type simpleGraph struct {
	g *simple.WeightedDirectedGraph
}

type node struct {
	id          int64
	name        string
	description string
	value       tensor.Tensor
	opType      string
	attributes  []*onnx.Attribute
}

// ID to fulfil the graph.Node interface
func (n *node) ID() int64 {
	return n.id
}

// Attributes it a method to fulfil the encoding/dot package
func (n *node) Attributes() []encoding.Attribute {
	var value string
	value = fmt.Sprintf(`%v`, n.name)
	if n.opType != "" {
		value = fmt.Sprintf(`%v|{Op|%v}`, value, n.opType)
	}
	if n.value != nil {
		value = fmt.Sprintf(`%v|{shape|%v}|{type|%v}`, value, n.value.Shape(), n.value.Dtype())
	}
	value = fmt.Sprintf(`"%v"`, value)
	return []encoding.Attribute{
		encoding.Attribute{
			Key:   "shape",
			Value: "Mrecord",
		},
		encoding.Attribute{
			Key:   "label",
			Value: value,
		},
	}
}

// SetDoc to fulfil the Docuemnter interface
func (n *node) SetDoc(name string) {
	n.description = name
}

// SetName to fulfil the Namer interface
func (n *node) SetName(desc string) {
	n.name = desc
}

// GetName to fulfil the Namer interface
func (n *node) GetName() string {
	return n.name
}

// SetDoc to fulfil the Namer interface
func (n *node) SetDescription(desc string) {
	n.description = desc
}

// GetDoc to fulfil the Namer interface
func (n *node) GetDescription() string {
	return n.description
}

func (n *node) SetValue(t tensor.Tensor) error {
	n.value = t
	return nil
}

func (n *node) GetValue() tensor.Tensor {
	return n.value
}

func (n *node) SetOpType(op string) {
	n.opType = op
}

func (n *node) SetOpAttributes(attrs []*onnx.Attribute) error {
	n.attributes = attrs
	return nil
}

type attributer []encoding.Attribute

func (a attributer) Attributes() []encoding.Attribute { return a }

func (g *simpleGraph) DOTAttributers() (graph, node, edge encoding.Attributer) {
	graphAttributes := attributer{
		encoding.Attribute{
			Key:   "label",
			Value: "name",
		},
		encoding.Attribute{
			Key:   "rankdir",
			Value: `"LR"`,
		},
	}
	nodeAttributes := attributer{
		encoding.Attribute{
			Key:   "style",
			Value: `"rounded,filled"`,
		},
		encoding.Attribute{
			Key:   "shape",
			Value: "Mrecord",
		},
		encoding.Attribute{
			Key:   "fillcolor",
			Value: "white",
		},
	}
	return graphAttributes, nodeAttributes, attributer{}
}

// SetDoc to fulfil the Documenter interface

func newSimpleGraph() *simpleGraph {
	return &simpleGraph{
		g: simple.NewWeightedDirectedGraph(math.MaxFloat64, -1),
	}
}

func (g *simpleGraph) SetWeightedEdge(e graph.WeightedEdge) {
	g.g.SetWeightedEdge(e)

}
func (g *simpleGraph) NewWeightedEdge(from, to graph.Node, w float64) graph.WeightedEdge {
	return g.g.NewWeightedEdge(from, to, w)

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
