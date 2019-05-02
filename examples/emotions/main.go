package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"time"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
	"github.com/owulveryck/onnx-go/internal/x/images"
	"gorgonia.org/tensor"
)

const (
	height = 64
	width  = 64
)

var emotionTable = []string{
	"neutral",
	"happiness",
	"surprise",
	"sadness",
	"anger",
	"disgust",
	"fear",
	"contempt",
}

func main() {
	model := flag.String("model", "model.onnx", "path to the model file")
	input := flag.String("input", "file.png", "path to the input file")
	h := flag.Bool("h", false, "help")
	flag.Parse()
	if *h {
		flag.Usage()
		os.Exit(0)
	}
	for _, f := range []string{*model, *input} {
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
	inputFile, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	img, err := png.Decode(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	imgGray, ok := img.(*image.Gray)
	if !ok {
		log.Fatal("Please give a gray image as input")
	}
	inputT := tensor.New(tensor.WithShape(1, 1, height, width), tensor.Of(tensor.Float32))
	err = images.GrayToBCHW(imgGray, inputT)
	if err != nil {
		log.Fatal(err)
	}
	m.SetInput(0, inputT)
	start := time.Now()
	err = backend.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Computation time: %v\n", time.Since(start))
	computedOutputT, err := m.GetOutputTensors()
	if err != nil {
		log.Fatal(err)
	}
	result := classify(softmax(computedOutputT[0].Data().([]float32)))
	fmt.Printf("%v / %2.2f%%\n", result[0].emotion, result[0].weight*100)
	fmt.Printf("%v / %2.2f%%\n", result[1].emotion, result[1].weight*100)
}

type testingT struct{}

func (t *testingT) Errorf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func softmax(input []float32) []float32 {
	var sumExp float64
	output := make([]float32, len(input))
	for i := 0; i < len(input); i++ {
		sumExp += math.Exp(float64(input[i]))
	}
	for i := 0; i < len(input); i++ {
		output[i] = float32(math.Exp(float64(input[i]))) / float32(sumExp)
	}
	return output
}

func classify(input []float32) emotions {
	result := make(emotions, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = emotion{
			emotion: emotionTable[i],
			weight:  input[i],
		}
	}
	sort.Sort(sort.Reverse(result))
	return result
}

type emotions []emotion
type emotion struct {
	emotion string
	weight  float32
}

func (e emotions) Len() int           { return len(e) }
func (e emotions) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e emotions) Less(i, j int) bool { return e[i].weight < e[j].weight }
