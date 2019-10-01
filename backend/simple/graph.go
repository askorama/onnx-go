package simple

import (
	"fmt"
	"math"

	"github.com/owulveryck/onnx-go"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/simple"
	"gorgonia.org/tensor"
)

// Graph is a simple directed graph
type Graph struct {
	g *simple.WeightedDirectedGraph
}

// Node of the graph
type Node struct {
	id          int64
	name        string
	description string
	value       tensor.Tensor
	opType      string
	//attributes  []*pb.Attribute
}

// ID to fulfil the graph.Node interface
func (n *Node) ID() int64 {
	return n.id
}

// Attributes it a method to fulfil the encoding/dot package
func (n *Node) Attributes() []encoding.Attribute {
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
		{
			Key:   "shape",
			Value: "Mrecord",
		},
		{
			Key:   "label",
			Value: value,
		},
	}
}

// SetName to fulfil the Namer interface
func (n *Node) SetName(desc string) {
	n.name = desc
}

// GetName to fulfil the Namer interface
func (n *Node) GetName() string {
	return n.name
}

// SetDescription to fulfil the Namer interface
func (n *Node) SetDescription(desc string) {
	n.description = desc
}

// GetDescription to fulfil the Namer interface
func (n *Node) GetDescription() string {
	return n.description
}

// ApplyTensor to fulfil the TensorCarrier interface
func (n *Node) ApplyTensor(t tensor.Tensor) error {
	n.value = t
	return nil
}

type attributer []encoding.Attribute

func (a attributer) Attributes() []encoding.Attribute { return a }

// DOTAttributers is used for representation in graphviz
func (g *Graph) DOTAttributers() (graph, node, edge encoding.Attributer) {
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

// NewSimpleGraph ...
func NewSimpleGraph() *Graph {
	return &Graph{
		g: simple.NewWeightedDirectedGraph(math.MaxFloat64, -1),
	}
}

// SetWeightedEdge adds a weighted edge from one node to another. If the nodes do not exist, they are added
// and are set to the nodes of the edge otherwise.
// It will panic if the IDs of the e.From and e.To are equal.
func (g *Graph) SetWeightedEdge(e graph.WeightedEdge) {
	g.g.SetWeightedEdge(e)

}

// NewWeightedEdge returns a new weighted edge from the source to the destination node.
func (g *Graph) NewWeightedEdge(from, to graph.Node, w float64) graph.WeightedEdge {
	return g.g.NewWeightedEdge(from, to, w)

}

// AddNode adds n to the graph. It panics if the added node ID matches an existing node ID.
func (g *Graph) AddNode(n graph.Node) {
	g.g.AddNode(n)

}

// NewNode returns a new unique Node to be added to g. The Node's ID does
// not become valid in g until the Node is added to g.
func (g *Graph) NewNode() graph.Node {
	n := g.g.NewNode()
	return &Node{
		id: n.ID(),
	}
}

// Node returns the node with the given ID if it exists in the graph,
// and nil otherwise.
func (g *Graph) Node(id int64) graph.Node {
	return g.g.Node(id)
}

// Nodes returns all the nodes in the graph.
func (g *Graph) Nodes() graph.Nodes {
	return g.g.Nodes()
}

// From returns all nodes in g that can be reached directly from n.
func (g *Graph) From(id int64) graph.Nodes {
	return g.g.From(id)
}

// HasEdgeBetween returns whether an edge exists between nodes x and y without
// considering direction.
func (g *Graph) HasEdgeBetween(xid, yid int64) bool {
	return g.g.HasEdgeBetween(xid, yid)
}

// Edge returns the edge from u to v if such an edge exists and nil otherwise.
// The node v must be directly reachable from u as defined by the From method.
func (g *Graph) Edge(uid, vid int64) graph.Edge {
	return g.g.Edge(uid, vid)
}

// HasEdgeFromTo returns whether an edge exists in the graph from u to v.
func (g *Graph) HasEdgeFromTo(uid, vid int64) bool {
	return g.g.HasEdgeFromTo(uid, vid)
}

// To returns all nodes in g that can reach directly to n.
func (g *Graph) To(id int64) graph.Nodes {
	return g.g.To(id)
}

// Apply operation to fulfil the OperationCarrier interface
func (g *Graph) ApplyOperation(_ onnx.Operation, _ ...graph.Node) error {
	return nil
}
