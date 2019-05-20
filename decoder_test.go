package onnx

import (
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
		// A -> B
		onnxModelProto: &pb.ModelProto{
			Graph: &pb.GraphProto{
				Node: []*pb.NodeProto{
					&pb.NodeProto{
						Name:   "A",
						Output: []string{"B"},
					},
					&pb.NodeProto{
						Name:  "B",
						Input: []string{"A"},
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
			err := m.decodeProto(test.onnxModelProto)
			assert.Equal(t, test.err, err)
			graphEqual(t, test.expected, m.backend)
		})
	}
}

func graphEqual(t *testing.T, src graph.WeightedDirected, dst Backend) {
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
	for i := 0; itSrc.Next(); i++ {
		srcNode := itSrc.Node().(*nodeTest)
		var dstNode *nodeTest
		var ok bool
		if dstNode, ok = dstNodes[srcNode.name]; !ok {
			t.Fatalf("node %v not found in generated graph", srcNode.name)
		}
		assertNodeEqual(t, srcNode, dstNode)
		srcTo := src.To(srcNode.ID())
		dstTo := dst.To(dstNode.ID())
		if srcTo.Len() != dstTo.Len() {
			t.Fatalf("expected node %v has %v parents but %v only have %v", srcNode, srcTo.Len(), srcTo, dstTo.Len())
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
