package main

import (
	"image"

	"github.com/owulveryck/onnx-go/internal/x/images"
	"gorgonia.org/tensor"
)

func process(img *image.Gray, height, width int, resultTable []string) ([]float32, error) {
	inputT := tensor.New(tensor.WithShape(1, 1, height, width), tensor.Of(tensor.Float32))
	err := images.GrayToBCHW(img, inputT)
	if err != nil {
		return nil, err
	}
	model.SetInput(0, inputT)
	err = backend.Run()
	if err != nil {
		return nil, err
	}
	computedOutputT, err := model.GetOutputTensors()
	if err != nil {
		return nil, err
	}
	return computedOutputT[0].Data().([]float32), nil
	//	return classify(softmax(computedOutputT[0].Data().([]float32)), resultTable), nil
}
