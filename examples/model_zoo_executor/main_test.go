package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
	"github.com/stretchr/testify/assert"
)

var (
	model, input, output []byte
)

func TestMain(m *testing.M) {
	dir, ok := os.LookupEnv("MODELDIR")
	if !ok {
		log.Println("$MODELDIR undefined")
		os.Exit(0)
	}
	modelFile := filepath.Join(dir, "model.onnx")
	var err error
	model, err = ioutil.ReadFile(modelFile)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	inputFile := filepath.Join(dir, "test_data_set_0/input_0.ir")
	input, err = ioutil.ReadFile(inputFile)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	outputFile := filepath.Join(dir, "test_data_set_0/output_0.ir")
	output, err = ioutil.ReadFile(outputFile)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func testUnmarshal(tb testing.TB) (backend.ComputationBackend, *onnx.Model) {
	// Create a backend receiver
	engine := gorgonnx.NewGraph()
	// Create a model and set the execution backend
	m := onnx.NewModel(engine)
	// Decode it into the model
	err := m.UnmarshalBinary(model)
	if err != nil {
		tb.Fatal(err)
	}
	return engine, m
}

func testSetInput(tb testing.TB, m *onnx.Model) {
	inputT, err := onnx.NewTensor(input)
	if err != nil {
		tb.Fatal(err)
	}
	m.SetInput(0, inputT)
}

func testRun(tb testing.TB, engine backend.ComputationBackend) {
	err := engine.Run()
	if err != nil {
		tb.Fatal(err)
	}
}

func TestModel(t *testing.T) {
	var engine backend.ComputationBackend
	var m *onnx.Model
	t.Run("Unmarshal", func(t *testing.T) {
		engine, m = testUnmarshal(t)
	})
	t.Run("Set input", func(t *testing.T) {
		testSetInput(t, m)
	})
	t.Run("Run", func(t *testing.T) {
		testRun(t, engine)
	})
	t.Run("Validate output", func(t *testing.T) {
		computedOutputT, err := m.GetOutputTensors()
		if err != nil {
			t.Fatal(err)
		}
		outputT, err := onnx.NewTensor(output)
		if err != nil {
			t.Fatal(err)
		}
		assert.InDeltaSlice(t, outputT.Data(), computedOutputT[0].Data(), 5e-3, "the two tensors should be equal.")
	})
}

func BenchmarkRun(b *testing.B) {
	engine, m := testUnmarshal(b)
	testSetInput(b, m)
	for n := 0; n < b.N; n++ {
		testRun(b, engine)
	}
}
