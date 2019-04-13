package gorgonnx

import (
	"testing"

	"github.com/owulveryck/onnx-go/backend/testbackend"
	_ "github.com/owulveryck/onnx-go/backend/testbackend/onnx"
)

func TestAdd(t *testing.T) {
	for _, tc := range testbackend.GetOpTypeTests("Add") {
		tc := tc // capture range variable
		t.Run(tc().GetInfo(), tc().RunTest(NewGraph(), false))
	}
}
