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
	var tests []func() *testbackend.TestCase
	if testing.Short() {
		for optype := range operators {
			tests = append(tests, testbackend.GetOpTypeTests(optype)...)
		}
	} else {
		tests = testbackend.GetAllRegisteredTests()
	}
	status := runner(t, tests)
	output, ok := os.LookupEnv("ONNX_COVERAGE_FILE")
	if ok {
		fmt.Println(output)
		f, err := os.Create(output)
		if err != nil {
			t.Log("Cannot open output file", err)
			return
		}
		defer f.Close()
		fmt.Fprintf(f, "|%-45v|%-10v|%-10v|\n", "Test", "Skipped", "Failed")
		fmt.Fprintf(f, "|---------------------------------------------|----------|----------|\n")
		for _, status := range status {
			fmt.Fprintf(f, "|%-45v|%-10v|%-10v|\n", status.info, status.skipped, status.failed)
		}
	}

}

func runner(t *testing.T, tests []func() *testbackend.TestCase) []report {
	t.Helper()
	status := make([]report, 0)
	for _, tc := range tests {
		tc := *tc() // capture range variable
		failed := t.Run(tc.GetInfo(), tc.RunTest(NewGraph(), false))
		status = append(status, report{tc.GetInfo(), failed, t.Skipped()})
	}
	return status
}
