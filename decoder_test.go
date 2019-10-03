package onnx

import (
	"fmt"
	"sort"
	"testing"

	"github.com/gogo/protobuf/proto"
	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/graph"
	"gorgonia.org/tensor"
)

type testGraph struct {
	name           string
	onnxModelProto *pb.ModelProto
	expected       graph.WeightedDirected
	err            error
}

var tests = []testGraph{
	{
		name:           "nil_graph",
		onnxModelProto: nil,
		expected:       &testExpectedGraph{},
		err:            errModelIsNil,
	},
	{
		name:           "empty model",
		onnxModelProto: &pb.ModelProto{},
		expected:       &testExpectedGraph{},
		err:            errGraphIsNil,
	},
	{
		name: "empty_graph",
		onnxModelProto: &pb.ModelProto{
			Graph: &pb.GraphProto{},
		},
		expected: &testExpectedGraph{},
		err:      errEmptyGraph,
	},
	{
		// A
		name: "graph with no input",
		onnxModelProto: &pb.ModelProto{
			Graph: &pb.GraphProto{
				Node: []*pb.NodeProto{
					{
						Name: "A",
					},
				},
			},
		},
		expected: &testExpectedGraph{},
		err:      errGraphNoIO,
	},
	{
		// A
		name: "graph with empty input",
		onnxModelProto: &pb.ModelProto{
			Graph: &pb.GraphProto{
				Node: []*pb.NodeProto{
					{
						Name: "A",
					},
				},
				Input: []*pb.ValueInfoProto{},
			},
		},
		expected: &testExpectedGraph{},
		err:      errGraphNoIO,
	},
	{
		name: "initializer with no input",
		// A is the Output
		// B is the Input
		// A -> B
		onnxModelProto: &pb.ModelProto{
			Graph: &pb.GraphProto{
				Node: []*pb.NodeProto{
					{
						Name:   "output",
						Input:  []string{"B"},
						Output: []string{"A"},
					},
				},
				Output: []*pb.ValueInfoProto{
					{
						Name: "A",
					},
				},
				Initializer: []*pb.TensorProto{
					{
						Name:      "B",
						DataType:  pb.TensorProto_DataType_value["FLOAT"],
						FloatData: []float32{0},
					},
				},
			},
		},
		expected: newExpectedGraph([]edge{
			{
				from: &nodeTest{
					id:   0,
					name: "A",
				},
				to: &nodeTest{
					id:    1,
					name:  "B",
					value: tensor.New(tensor.Of(tensor.Float32), tensor.WithBacking([]float32{0})),
				},
				weight: 0,
			},
		}),
		err: nil,
	},
	{
		name: "simple graph",
		// A is the Output
		// B is the Input
		// A -> B
		onnxModelProto: &pb.ModelProto{
			Graph: &pb.GraphProto{
				Node: []*pb.NodeProto{
					{
						Name:   "output",
						Input:  []string{"B"},
						Output: []string{"A"},
					},
				},
				Output: []*pb.ValueInfoProto{
					{
						Name: "A",
					},
				},
				Input: []*pb.ValueInfoProto{
					{
						Name: "B",
					},
				},
			},
		},
		expected: newExpectedGraph([]edge{
			{
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

func TestDecodeProto_badBackend(t *testing.T) {
	m := NewModel(nil)
	err := m.decodeProto(nil)
	_, ok := err.(*InvalidUnmarshalError)
	assert.True(t, ok, fmt.Sprintf("Expected an InvalidUnmarshalError, but got %#v", err))
}

func TestDecodeProto(t *testing.T) {
	for _, test := range tests {
		m := NewModel(newTestBackend())
		test := test // capture range variable
		t.Run(test.name, func(t *testing.T) {
			//t.Parallel()
			err := m.decodeProto(test.onnxModelProto)
			assert.Equal(t, test.err, err)
			assertGraphEqual(t, test.expected, m.backend)
		})
	}
}

func TestUnmarshalBinary(t *testing.T) {
	m := NewModel(newTestBackend())
	b := []byte{0}
	err := m.UnmarshalBinary(b)
	assert.Error(t, err, "Expected an error")
	model := &pb.ModelProto{}
	b, err = proto.Marshal(model)
	assert.NoError(t, err)
	err = m.UnmarshalBinary(b)
	assert.Equal(t, err, errGraphIsNil, fmt.Sprintf("bad error, expected errGraphIsNil, got %v", err))
}

func TestProcessValue(t *testing.T) {
	m := NewModel(newTestBackend())
	_, err := m.processValue(nil)
	assert.Error(t, err)
	io := &pb.ValueInfoProto{
		Name: "name",
		Type: &pb.TypeProto{
			Value: &pb.TypeProto_TensorType{
				TensorType: &pb.TypeProto_Tensor{
					ElemType: pb.AttributeProto_AttributeType_value["FLOAT"],
					Shape: &pb.TensorShapeProto{
						Dim: []*pb.TensorShapeProto_Dimension{
							{
								Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 1},
							},
							{
								Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 2},
							},
							{
								Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 3},
							},
							{
								Value: &pb.TensorShapeProto_Dimension_DimValue{DimValue: 4},
							},
						},
					},
				},
			},
		},
	}
	n, err := m.processValue(io)
	assert.NoError(t, err)

	expectedNode := &nodeTest{
		name:  "name",
		value: tensor.New(tensor.WithShape(1, 2, 3, 4), tensor.Of(tensor.Float32)),
	}
	assertNodeEqual(t, expectedNode, n.(*nodeTest))

}

func assertGraphEqual(t *testing.T, src graph.WeightedDirected, dst Backend) {
	if src == nil && dst == nil {
		return
	}
	assert.NotNil(t, src)
	assert.NotNil(t, dst)
	itSrc := src.Nodes()
	itDst := dst.Nodes()
	assert.Equal(t, itSrc.Len(), itDst.Len(),
		fmt.Sprintf("graphs differs: expected %v node(s) but graph have %v node(s)", itSrc.Len(), itDst.Len()))
	dstNodes := make(map[string]*nodeTest, itDst.Len())
	for i := 0; itDst.Next(); i++ {
		n := itDst.Node().(*nodeTest)
		dstNodes[n.name] = n
	}
	for itSrc.Next() {
		srcNode := itSrc.Node().(*nodeTest)
		var dstNode *nodeTest
		var ok bool
		dstNode, ok = dstNodes[srcNode.name]
		assert.True(t, ok, fmt.Sprintf("node %v not found in generated graph", srcNode.name))
		assertNodeEqual(t, srcNode, dstNode)
		fromSrcNode := src.From(srcNode.ID())
		fromDstNode := dst.From(dstNode.ID())
		assert.Equal(t, fromSrcNode.Len(), fromDstNode.Len(),
			fmt.Sprintf("expected node %v has %v child(ren) but %v have %v",
				srcNode, fromSrcNode.Len(), dstNode, fromDstNode.Len()))
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
		assert.Equal(t, len(dstWeightedEdges), len(srcWeightedEdges), "not the same number of edges")
		if len(srcWeightedEdges) == 0 {
			continue
		}
		// Compare the weights
		for i := 0; i < len(srcWeightedEdges); i++ {
			assert.Equal(t, srcWeightedEdges[i].Weight(), dstWeightedEdges[i].Weight(),
				fmt.Sprintf("Expected weight %v, got %v",
					srcWeightedEdges[i].Weight(), dstWeightedEdges[i].Weight()))
			assertNodeEqual(t, srcWeightedEdges[i].From().(*nodeTest), dstWeightedEdges[i].From().(*nodeTest))
			assertNodeEqual(t, srcWeightedEdges[i].To().(*nodeTest), dstWeightedEdges[i].To().(*nodeTest))
		}
	}
}

func assertNodeEqual(t *testing.T, a, b *nodeTest) {
	assert.Equal(t, a.opType, b.opType, fmt.Sprintf("nodes %v and %v differs", a, b))
	if a.value != nil && b.value != nil {
		_, err := tensor.ElEq(a.value, b.value)
		assert.NoError(t, err)
	}
	assert.NotNil(t, a.value, fmt.Sprintf("nodes %v and %v differs", a, b))
	assert.NotNil(t, b.value, fmt.Sprintf("nodes %v and %v differs", a, b))
	assert.Equal(t, a.name, b.name, fmt.Sprintf("nodes %v and %v differs", a, b))

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
