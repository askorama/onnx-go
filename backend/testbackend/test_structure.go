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

// Register a test
func Register(optype, testTitle string, constructor func() *TestCase) {
	allOpTypes[optype] = append(allOpTypes[optype], constructor)
	allTests[testTitle] = constructor
}

// allOpTypes returns all the tests for a given OpType
var allOpTypes = map[string][]func() *TestCase{}

// allTests holds a reference of the test regarding their name
var allTests = map[string]func() *TestCase{}

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
	OpType         string
	Title          string
	ModelB         []byte
	Input          []tensor.Tensor
	ExpectedOutput []tensor.Tensor
	Tested         bool // true if the test has be executed
	Skipped        bool // true if the test has been executed and skipped
	Failed         bool // true if the test failed
}

// GetInfo ...
func (tc *TestCase) GetInfo() string {
	return tc.Title
}

type testWrapper struct {
	tc *TestCase
	t  *testing.T
}

func (tw testWrapper) Errorf(format string, args ...interface{}) {
	tw.tc.Failed = true
	tw.t.Errorf(format, args...)

}

// RunTest Returns a function to be executed against the ComputationBackend.
// The return function should be executed via a call to testing.Run(...)
// If parallel is true, a t.Parallel() is added at the beginning of the test
func (tc *TestCase) RunTest(b backend.ComputationBackend, parallel bool) func(t *testing.T) {
	return func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatal(r)
			}
		}()
		tc.Tested = true
		if parallel {
			t.Parallel()
		}
		m := onnx.NewModel(b)
		err := m.UnmarshalBinary(tc.ModelB)
		if err != nil {
			if _, ok := err.(*onnx.ErrNotImplemented); ok {
				tc.Skipped = true
				t.Skip(err)
			}
			tc.Failed = true
			t.Fatal(err)
		}
		for i := range tc.Input {
			err := m.SetInput(i, tc.Input[i])
			if err != nil {
				tc.Failed = true
				t.Fatal(err)
			}
		}

		err = b.Run()
		if err != nil {
			if _, ok := err.(*onnx.ErrNotImplemented); ok {
				tc.Skipped = true
				t.Skip(err)
			}
			tc.Failed = true
			t.Fatal(err)
		}
		output, err := m.GetOutputTensors()
		if err != nil {
			tc.Failed = true
			t.Fatal(err)
		}

		if len(output) != len(tc.ExpectedOutput) {
			tc.Failed = true
			t.Fatalf("expected %v output, got %v", len(tc.ExpectedOutput), len(output))
		}
		tw := testWrapper{tc, t}
		for i := range output {
			if len(tc.ExpectedOutput[i].Shape()) != len(output[i].Shape()) {
				t.Fatalf("the two tensors doesn't have the same dimension, expected %v, got %v", tc.ExpectedOutput[i].Shape(), output[i].Shape())
			}
			for j, v := range tc.ExpectedOutput[i].Shape() {
				if v != output[i].Shape()[j] {
					t.Fatalf("the two tensors doesn't have the same dimension, expected %v, got %v", tc.ExpectedOutput[i].Shape(), output[i].Shape())
				}
			}
			assert.InDeltaSlice(tw, tc.ExpectedOutput[i].Data(), output[i].Data(), 1e-6, "the two tensors should be equal.")
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
