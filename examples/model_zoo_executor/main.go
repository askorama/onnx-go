package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
	"github.com/stretchr/testify/assert"
)

func main() {
	model := flag.String("model", "model.onnx", "path to the model file")
	input := flag.String("input", "test_data_set_0/input_0.pb", "path to the input file")
	output := flag.String("output", "test_data_set_0/output_0.pb", "path to the output file")
	h := flag.Bool("h", false, "help")
	flag.Parse()
	if *h {
		flag.Usage()
		os.Exit(0)
	}
	for _, f := range []string{*model, *input, *output} {
		if _, err := os.Stat(f); err != nil && os.IsNotExist(err) {
			log.Fatalf("%v does not exist", f)
		}
	}
	// Create a backend receiver
	backend := gorgonnx.NewGraph()
	// Create a model and set the execution backend
	m := onnx.NewModel(backend)

	// read the onnx model
	b, err := ioutil.ReadFile(*model)
	if err != nil {
		log.Fatal(err)
	}
	// Decode it into the model
	err = m.UnmarshalBinary(b)
	if err != nil {
		log.Fatal(err)
	}
	// Set the first input, the number depends of the model
	// TODO
	b, err = ioutil.ReadFile(*input)
	if err != nil {
		log.Fatal(err)
	}
	inputT, err := onnx.NewTensor(b)
	if err != nil {
		log.Fatal(err)
	}
	m.SetInput(0, inputT)
	err = backend.Run()
	if err != nil {
		log.Fatal(err)
	}
	b, err = ioutil.ReadFile(*output)
	if err != nil {
		log.Fatal(err)
	}
	outputT, err := onnx.NewTensor(b)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(&testingT{}, outputT, m.Output[0])
}

type testingT struct{}

func (t *testingT) Errorf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
