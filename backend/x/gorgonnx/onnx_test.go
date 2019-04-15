package gorgonnx

import (
	"fmt"
	"os"
	"testing"

	"github.com/owulveryck/onnx-go/backend/testbackend"
	_ "github.com/owulveryck/onnx-go/backend/testbackend/onnx"
)

type report struct {
	info    string
	failed  bool
	skipped bool
}

// TestONNX run the onnx's backend tests against all registered operatos
func TestONNX(t *testing.T) {
	for optype := range operators {
		optype := optype
		for _, tc := range testbackend.GetOpTypeTests(optype) {
			tc := *tc() // capture range variable
			t.Run(tc.GetInfo(), tc.RunTest(NewGraph(), true))
		}
	}
}

func TestONNXCoverage(t *testing.T) {
	if _, ok := os.LookupEnv("ONNX_COVERAGE"); !ok {
		t.SkipNow()
	}
	status := make([]report, 0)
	for _, tc := range testbackend.GetAllRegisteredTests() {
		tc := *tc() // capture range variable
		failed := t.Run(tc.GetInfo(), tc.RunTest(NewGraph(), true))
		status = append(status, report{tc.GetInfo(), failed, t.Skipped()})
	}
	for i := range status {
		status := status[i]
		fmt.Printf("|%-40v|%-10v|%-10v|\n", status.info, status.skipped, status.failed)

	}

}
