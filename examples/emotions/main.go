package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"io"
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
	if _, err := os.Stat(*model); err != nil && os.IsNotExist(err) {
		log.Fatalf("%v does not exist", *model)
	}
	if _, err := os.Stat(*input); err != nil && *input != "-" && os.IsNotExist(err) {
		log.Fatalf("%v does not exist", *input)
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
	fmt.Println("Finished loading ONNX Model")
	// Set the first input, the number depends of the model
	// TODO
	inputStream := createInputStream(input)
	fmt.Println(("0 - image work - finished creating input stream"))
	img, err := png.Decode(inputStream)
	fmt.Println(("1 - image work - PNG decoded"))
	if err != nil {
		fmt.Println(("LOG ERROR"))
		log.Fatal(err)
	}

	//------
	/*
	// Working with grayscale image, e.g. convert to png
    f, err := os.Create("surprise-gray.png")
    if err != nil {
        // handle error
        log.Fatal(err)
    }

	// Convert image to grayscale
    grayImg := image.NewGray(img.Bounds())
    for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
        for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
            grayImg.Set(x, y, img.At(x, y))
        }
    }

	if err := png.Encode(f, grayImg); err != nil {
        log.Fatal(err)
    }

	f.Close()
	*/
	//-------

	imgGray, ok := img.(*image.Gray)
	fmt.Println(("2 - image work - finished image Gray"))
	if !ok {
		log.Fatal("Please give a gray image as input")
	}
	fmt.Println(("3 - image work"))
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
	fmt.Println(("4 - get output tensors"))
	result := classify(softmax(computedOutputT[0].Data().([]float32)))
	fmt.Println(("5 - classify"))
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

func createInputStream(input *string) io.Reader {
	var inputStream io.Reader
	if *input != "-" {
		imgContent, err := os.Open(*input)
		if err != nil {
			log.Fatal(err)
		}
		//defer imgContent.Close()
		inputStream = imgContent
	} else {
		inputStream = os.Stdin
	}
	return inputStream
}

type emotions []emotion
type emotion struct {
	emotion string
	weight  float32
}

func (e emotions) Len() int           { return len(e) }
func (e emotions) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e emotions) Less(i, j int) bool { return e[i].weight < e[j].weight }
