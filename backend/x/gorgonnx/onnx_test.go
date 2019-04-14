package gorgonnx

import (
	"testing"

	"github.com/owulveryck/onnx-go/backend/testbackend"
	_ "github.com/owulveryck/onnx-go/backend/testbackend/onnx"
)

// TestONNX run the onnx's backend tests against all registered operatos
func TestONNX(t *testing.T) {
	for optype := range operators {
		optype := optype
		for _, tc := range testbackend.GetOpTypeTests(optype) {
			tc := tc() // capture range variable
			t.Run(tc.GetInfo(), tc.RunTest(NewGraph(), true))
		}
	}
}
