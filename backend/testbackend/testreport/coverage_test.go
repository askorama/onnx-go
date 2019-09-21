package testreport

import (
	"testing"

	"github.com/owulveryck/onnx-go/backend/testbackend"
)

func TestCoverage(t *testing.T) {
	tests := []*testbackend.TestCase{
		{
			Skipped: false,
			Tested:  true,
		},
		{
			Skipped: true,
			Tested:  true,
		},
		{
			Skipped: true,
			Tested:  true,
		},
	}
	val := Coverage(tests)
	expected := 100 - float64(2)*100/float64(3)
	if val != expected {
		t.Fatalf("expected %v, got %v", expected, val)
	}
}
