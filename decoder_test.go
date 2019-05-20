package onnx

import (
	"testing"

	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/graph"
)

type testGraph struct {
	onnxModelProto *pb.ModelProto
	expected       graph.WeightedDirected
	err            error
}

var tests = []testGraph{
	testGraph{
		onnxModelProto: &pb.ModelProto{},
		expected:       &testWeightedDirectedGraph{},
		err:            errGraphIsNil,
	},
	testGraph{
		onnxModelProto: &pb.ModelProto{
			Graph: &pb.GraphProto{},
		},
		expected: &testWeightedDirectedGraph{},
		err:      nil,
	},
	testGraph{
		// A -> B
		onnxModelProto: &pb.ModelProto{
			Graph: &pb.GraphProto{
				Node: []*pb.NodeProto{
					&pb.NodeProto{
						Name: "A",
					},
					&pb.NodeProto{
						Name:  "B",
						Input: []string{"A"},
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
			},
		}),
		err: nil,
	},
}

func TestDecodeProto(t *testing.T) {
	m := NewModel(newTestBackend())
	for _, test := range tests {
		err := m.decodeProto(test.onnxModelProto)
		assert.Equal(t, test.err, err)
		graphEqual(t, test.expected, m.backend)
	}
}

func graphEqual(t *testing.T, src graph.WeightedDirected, dst Backend) {
	itSrc := src.Nodes()
	itDst := dst.Nodes()
	t.Logf("src has %v nodes and Dst has %v nodes", itSrc.Len(), itDst.Len())
	if itSrc.Len() != itDst.Len() {
		t.Fatalf("graphs differs: src has %v nodes and Dst has %v nodes", itSrc.Len(), itDst.Len())
	}
	// TODO compare the graphs
}

type edge struct {
	from *nodeTest
	to   *nodeTest
}

func newExpectedGraph(e []edge) *testWeightedDirectedGraph {
	return &testWeightedDirectedGraph{
		//TODO
	}
}
