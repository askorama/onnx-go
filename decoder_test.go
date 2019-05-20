package onnx

import (
	"sort"
	"testing"

	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/graph"
)

type testGraph struct {
	name           string
	onnxModelProto *pb.ModelProto
	expected       graph.WeightedDirected
	err            error
}

var tests = []testGraph{
	testGraph{
		name:           "nil_graph",
		onnxModelProto: &pb.ModelProto{},
		expected:       &testExpectedGraph{},
		err:            errGraphIsNil,
	},
	testGraph{
		name: "empty_graph",
		onnxModelProto: &pb.ModelProto{
			Graph: &pb.GraphProto{},
		},
		expected: &testExpectedGraph{},
		err:      errEmptyGraph,
	},
	testGraph{
		// A
		name: "graph with no input",
		onnxModelProto: &pb.ModelProto{
			Graph: &pb.GraphProto{
				Node: []*pb.NodeProto{
					&pb.NodeProto{
						Name: "A",
					},
				},
			},
		},
		expected: &testExpectedGraph{},
		err:      errGraphNoIO,
	},
	testGraph{
		// A
		name: "graph with empty input",
		onnxModelProto: &pb.ModelProto{
			Graph: &pb.GraphProto{
				Node: []*pb.NodeProto{
					&pb.NodeProto{
						Name: "A",
					},
				},
				Input: []*pb.ValueInfoProto{},
			},
		},
		expected: &testExpectedGraph{},
		err:      errGraphNoIO,
	},
	testGraph{
		name: "simple graph",
		// A is the Output
		// B is the Input
		// A -> B
		onnxModelProto: &pb.ModelProto{
			Graph: &pb.GraphProto{
				Node: []*pb.NodeProto{
					&pb.NodeProto{
						Name:   "output",
						Input:  []string{"B"},
						Output: []string{"A"},
					},
				},
				Output: []*pb.ValueInfoProto{
					&pb.ValueInfoProto{
						Name: "A",
					},
				},
				Input: []*pb.ValueInfoProto{
					&pb.ValueInfoProto{
						Name: "B",
					},
				},
			},
		},
		expected: newExpectedGraph([]edge{
			edge{
				from: &nodeTest{
					id:   0,
					name: "A",
				},
				to: &nodeTest{
					id:   1,
					name: "B",
				},
				weight: 0,
			},
		}),
		err: nil,
	},
}

func TestDecodeProto(t *testing.T) {
	m := NewModel(newTestBackend())
	for _, test := range tests {
		test := test // capture range variable
		t.Run(test.name, func(t *testing.T) {
			//t.Parallel()
			err := m.decodeProto(test.onnxModelProto)
			assert.Equal(t, test.err, err)
			assertGraphEqual(t, test.expected, m.backend)
		})
	}
}

func assertGraphEqual(t *testing.T, src graph.WeightedDirected, dst Backend) {
	itSrc := src.Nodes()
	itDst := dst.Nodes()
	if itSrc.Len() != itDst.Len() {
		t.Fatalf("graphs differs: expected %v node(s) but graph have %v node(s)", itSrc.Len(), itDst.Len())
	}
	dstNodes := make(map[string]*nodeTest, itDst.Len())
	for i := 0; itDst.Next(); i++ {
		n := itDst.Node().(*nodeTest)
		dstNodes[n.name] = n
	}
	for itSrc.Next() {
		srcNode := itSrc.Node().(*nodeTest)
		var dstNode *nodeTest
		var ok bool
		if dstNode, ok = dstNodes[srcNode.name]; !ok {
			t.Fatalf("node %v not found in generated graph", srcNode.name)
		}
		assertNodeEqual(t, srcNode, dstNode)
		fromSrcNode := src.From(srcNode.ID())
		fromDstNode := dst.From(dstNode.ID())
		if fromSrcNode.Len() != fromDstNode.Len() {
			t.Fatalf("expected node %v has %v child(ren) but %v have %v", srcNode, fromSrcNode.Len(), dstNode, fromDstNode.Len())
		}
		srcWeightedEdges := make([]graph.WeightedEdge, fromSrcNode.Len())
		dstWeightedEdges := make([]graph.WeightedEdge, fromDstNode.Len())
		for i := 0; fromSrcNode.Next(); i++ {
			srcNodeFrom := fromSrcNode.Node()
			srcWeightedEdges[i] = src.WeightedEdge(srcNode.ID(), srcNodeFrom.ID())
		}
		sort.Sort(weightedEdge(srcWeightedEdges))
		for i := 0; fromDstNode.Next(); i++ {
			dstNodeFrom := fromDstNode.Node()
			dstWeightedEdges[i] = dst.(weightedBackend).WeightedEdge(dstNode.ID(), dstNodeFrom.ID())
		}
		sort.Sort(weightedEdge(dstWeightedEdges))
		if len(dstWeightedEdges) != len(srcWeightedEdges) {
			t.Fatalf("not the same number of edges")
		}
		if len(srcWeightedEdges) == 0 {
			continue
		}
		// Compare the weights
		for i := 0; i < len(srcWeightedEdges); i++ {
			if srcWeightedEdges[i].Weight() != dstWeightedEdges[i].Weight() {
				t.Fatalf("Expected weight %v, got %v", srcWeightedEdges[i].Weight(), dstWeightedEdges[i].Weight())
			}
			assertNodeEqual(t, srcWeightedEdges[i].From().(*nodeTest), dstWeightedEdges[i].From().(*nodeTest))
			assertNodeEqual(t, srcWeightedEdges[i].To().(*nodeTest), dstWeightedEdges[i].To().(*nodeTest))
		}
	}
}

func assertNodeEqual(t *testing.T, a, b *nodeTest) {
	if a.opType != b.opType {
		t.Fatalf("nodes %v and %v differs", a, b)
	}
	if a.value != b.value {
		t.Fatalf("nodes %v and %v differs", a, b)
	}
	if a.name != b.name {
		t.Fatalf("nodes %v and %v differs", a, b)
	}

}

type weightedEdge []graph.WeightedEdge

func (e weightedEdge) Len() int           { return len(e) }
func (e weightedEdge) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e weightedEdge) Less(i, j int) bool { return e[i].Weight() < e[j].Weight() }

type weightedBackend interface {
	Backend
	// WeightedEdge returns the weighted edge from u to v
	// with IDs uid and vid if such an edge exists and
	// nil otherwise. The node v must be directly
	// reachable from u as defined by the From method.
	WeightedEdge(uid, vid int64) graph.WeightedEdge
}
