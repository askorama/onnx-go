package gorgonnx

import (
	"testing"

	"github.com/owulveryck/onnx-go/backend/testbackend"
	_ "github.com/owulveryck/onnx-go/backend/testbackend/onnx"
)

func TestCos(t *testing.T) {
	for _, tc := range testbackend.GetOpTypeTests("Cos") {
		tc := tc // capture range variable
		t.Run(tc().GetInfo(), tc().RunTest(NewGraph(), false))
	}
}
