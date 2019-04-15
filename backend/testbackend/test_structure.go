package testbackend

import (
	"io"
	"regexp"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend"
	"github.com/owulveryck/onnx-go/internal/pb-onnx"
	"github.com/stretchr/testify/assert"
	"gorgonia.org/tensor"
)

func init() {
	allOpTypes = make(map[string][]func() *TestCase, 0)
	allTests = make(map[string]func() *TestCase, 0)
}

// Register a test
func Register(optype, testTitle string, constructor func() *TestCase) {
	allOpTypes[optype] = append(allOpTypes[optype], constructor)
	allTests[testTitle] = constructor
}

// allOpTypes returns all the tests for a given OpType
var allOpTypes map[string][]func() *TestCase

// allTests holds a reference of the test regarding their name
var allTests map[string]func() *TestCase

// GetAllRegisteredTests ...
func GetAllRegisteredTests() []func() *TestCase {
	output := make([]func() *TestCase, 0)
	for _, v := range allTests {
		output = append(output, v)
	}
	return output

}

// FindAllTestsMatching tests matching the regexp
func FindAllTestsMatching(re *regexp.Regexp) []func() *TestCase {
	output := make([]func() *TestCase, 0)
	for k, v := range allTests {
		if re.MatchString(k) {
			output = append(output, v)
		}
	}
	return output
}

// GetOpTypeTests returns all the tests of the OpType passed as argument
func GetOpTypeTests(optype string) []func() *TestCase {
	return allOpTypes[optype]
}

// TestCase describes an integration test
type TestCase struct {
	Title          string
	ModelB         []byte
	Input          []tensor.Tensor
	ExpectedOutput []tensor.Tensor
}

// GetInfo ...
func (tc *TestCase) GetInfo() string {
	return tc.Title
}

// RunTest Returns a function to be executed against the ComputationBackend.
// The return function should be executed via a call to testing.Run(...)
// If parallel is true, a t.Parallel() is added at the begining of the test
func (tc *TestCase) RunTest(b backend.ComputationBackend, parallel bool) func(t *testing.T) {
	return func(t *testing.T) {
		if parallel {
			t.Parallel()
		}
		m := onnx.NewModel(b)
		err := m.UnmarshalBinary(tc.ModelB)
		if err != nil {
			if _, ok := err.(*onnx.ErrNotImplemented); ok {
				t.Skip(err)
			}
			t.Fatal(err)
		}
		for i := range tc.Input {
			err := m.SetInput(i, tc.Input[i])
			if err != nil {
				t.Fatal(err)
			}
		}

		err = b.Run()
		if err != nil {
			if _, ok := err.(*onnx.ErrNotImplemented); ok {
				t.Skip(err)
			}
			t.Fatal(err)
		}
		output, err := m.GetOutputTensors()
		if err != nil {
			t.Fatal(err)
		}

		if len(output) != len(tc.ExpectedOutput) {
			t.Fatalf("expected %v output, got %v", len(tc.ExpectedOutput), len(output))
		}
		for i := range output {
			assert.InDeltaSlice(t, tc.ExpectedOutput[i].Data(), output[i].Data(), 1e-6, "the two tensors should be equal.")
		}

	}
}

// Dump a raw version of the onnx data decoded in the protobuf structure.
// Useful for debugging
func (tc *TestCase) Dump(w io.Writer) error {
	model := new(pb.ModelProto)
	err := model.XXX_Unmarshal(tc.ModelB)
	if err != nil {
		return err
	}
	scs := spew.ConfigState{
		Indent:                  "\t",
		DisablePointerAddresses: true,
		DisableCapacities:       true,
	}

	scs.Fdump(w, model)
	return nil
}
