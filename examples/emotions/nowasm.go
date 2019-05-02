// +build !wasm

package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	modelFile string
	input     string
	h         bool
)

func init() {
	flag.StringVar(&modelFile, "model", "model.onnx", "path to the model file")
	flag.StringVar(&input, "input", "file.png", "path to the input file")
	flag.BoolVar(&h, "h", false, "help")
	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(0)
	}
	if _, err := os.Stat(modelFile); err != nil && os.IsNotExist(err) {
		log.Fatalf("%v does not exist", modelFile)
	}
	if _, err := os.Stat(input); err != nil && input != "-" && os.IsNotExist(err) {
		log.Fatalf("%v does not exist", input)
	}

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getModel() ([]byte, error) {
	// read the onnx model
	return ioutil.ReadFile(modelFile)
}

func getImage() (*image.Gray, error) {
	var inputStream io.Reader
	if input != "-" {
		imgContent, err := os.Open(input)
		if err != nil {
			return nil, err
		}
		defer imgContent.Close()
		inputStream = imgContent
	} else {
		inputStream = os.Stdin
	}
	img, err := png.Decode(inputStream)
	if err != nil {
		return nil, err
	}
	imgGray, ok := img.(*image.Gray)
	if !ok {
		return nil, errors.New("Please give a gray image as input")
	}
	return imgGray, nil
}

func displayResult(e emotions) {
	fmt.Printf("%v / %2.2f%%\n", e[0].emotion, e[0].weight*100)
	fmt.Printf("%v / %2.2f%%\n", e[1].emotion, e[1].weight*100)
}

func run() error {
	// Decode it into the model
	b, err := getModel()
	if err != nil {
		return err
	}
	err = model.UnmarshalBinary(b)
	if err != nil {
		return err
	}
	// Set the first input, the number depends of the model
	img, err := getImage()
	if err != nil {
		return err
	}
	result, err := process(img)
	if err != nil {
		return err
	}
	displayResult(result)
	return nil
}
