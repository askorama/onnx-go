package onnx

import (
	"testing"

	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"github.com/stretchr/testify/assert"
)

type testGraph struct {
	onnx     *pb.ModelProto
	expected *testBackend
	err      error
}

var tests = []testGraph{
	testGraph{
		onnx:     &pb.ModelProto{},
		expected: &testBackend{},
		err:      errGraphIsNil,
	},
}

func TestDecodeProto(t *testing.T) {
	m := NewModel(newTestBackend())
	for _, test := range tests {
		err := m.decodeProto(test.onnx)
		assert.Equal(t, test.err, err)
		graphEqual(t, test.expected, m.backend)
	}
}

func graphEqual(t *testing.T, src, dst Backend) {
	// TODO compare the graphs
}
