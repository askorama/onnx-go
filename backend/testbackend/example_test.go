package testbackend_test

import (
	"testing"

	"github.com/owulveryck/onnx-go/backend"
	"github.com/owulveryck/onnx-go/backend/testbackend"
	// Register the tests from onnx
	// _ "github.com/owulveryck/onnx-go/backend/testbackend/onnx"
)

func Example() {
	// The backend you want to test
	var backend backend.ComputationBackend
	// Replace this line in your code by:
	// func TestConv(t i*testing.T) {
	testConv := func(t *testing.T) {
		for _, tc := range testbackend.GetOpTypeTests("Conv") {
			tc := tc // capture range variable
			t.Run(tc().GetInfo(), tc().RunTest(backend, false))
		}
	}
	_ = testConv
}
