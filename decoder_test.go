package onnx

import (
	"testing"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/simple"
	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"github.com/stretchr/testify/assert"
)

type testsGraph struct {
	onnx     *pb.ModelProto
	expected *simple.Graph
	err      error
}

var tests = []testGraph{}

func TestDecodeProto(t *testing.T) {
	m := NewModel(simple.NewSimpleGraph())
	for _, test := range tests {
		err := m.decodeProto(test.onnx)
		assert.Equal(t, test.err, err)
		graphEqual(t, test.expected, m.backend)
	}
}

func graphEqual(t *testing.T, src, dst onnx.Backend) {
	// TODO compare teh graphs
}
