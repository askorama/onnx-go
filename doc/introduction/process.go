package main

import (
	"image"

	"github.com/owulveryck/onnx-go/internal/x/images"
	"gorgonia.org/tensor"
)

// START_PROCESS OMIT
func process(img *image.Gray, height, width int, resultTable []string) ([]float32, error) {
	inputT := tensor.New(tensor.WithShape(1, 1, height, width), tensor.Of(tensor.Float32))
	err := images.GrayToBCHW(img, inputT)
	if err != nil { // OMIT
		return nil, err // OMIT
	} // OMIT
	model.SetInput(0, inputT)
	err = backend.Run() // HL
	if err != nil {     // OMIT
		return nil, err // OMIT
	} // OMIT
	computedOutputT, err := model.GetOutputTensors()
	if err != nil { // OMIT
		return nil, err // OMIT
	} // OMIT
	return computedOutputT[0].Data().([]float32), nil
}

// END_PROCESS OMIT
