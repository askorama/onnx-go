package main

import (
	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
)

func init() {
	// Create a backend receiver
	backend = gorgonnx.NewGraph()
	// Create a model and set the execution backend
	model = onnx.NewModel(backend)
}

const (
	height = 64
	width  = 64
)

var (
	backend      *gorgonnx.Graph
	model        *onnx.Model
	emotionTable = []string{
		"neutral",
		"happiness",
		"surprise",
		"sadness",
		"anger",
		"disgust",
		"fear",
		"contempt",
	}
)
