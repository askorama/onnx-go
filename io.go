package onnx

import "gorgonia.org/tensor"

// SetInput assign a tensor to the i-th input of the graph
func (m *Model) SetInput(i int, t tensor.Tensor) error {
	//TODO
	return nil
}

// GetOutputs of the graph
func (m *Model) GetOutput() []tensor.Tensor {
	// TODO
	return nil
}
