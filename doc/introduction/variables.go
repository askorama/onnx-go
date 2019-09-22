package main

import (
	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
)

var (
	backend    *gorgonnx.Graph
	model      *onnx.Model
	mnistTable = []string{
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}
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
	models = map[string]modelDemo{
		"mnist": {
			height: 28,
			width:  28,
			table:  mnistTable,
		},
		"emotion": {
			height:         64,
			width:          64,
			table:          emotionTable,
			postProcessing: softmax,
		},
	}
)

type modelDemo struct {
	height         int
	width          int
	table          []string
	postProcessing func([]float32) []float32
}
