package onnx

import (
	fmt "fmt"

	"gorgonia.org/tensor"
)

// SetInput assign a tensor to the i-th input of the graph
func (m *Model) SetInput(i int, t tensor.Tensor) error {
	if i >= len(m.Input) {
		return fmt.Errorf("error, trying to set input #%v, but model has only #%v input", i, len(m.Input))
	}
	// Get the corresponding node
	n := m.backend.Node(int64(i))
	if n == nil {
		return fmt.Errorf("cannot set input for node %v, node is nil", i)
	}

	if _, ok := n.(DataCarrier); !ok {
		return fmt.Errorf("cannot set input because node is not a DataCarrier")
	}
	return n.(DataCarrier).SetTensor(t)
}

// GetOutputTensors of the graph
func (m *Model) GetOutputTensors() ([]tensor.Tensor, error) {
	output := make([]tensor.Tensor, len(m.Output))
	for i := range m.Output {
		n := m.backend.Node(int64(m.Output[i]))
		if n == nil {
			return nil, fmt.Errorf("cannot get output for node %v, node is nil", i)
		}
		if _, ok := n.(DataCarrier); !ok {
			return nil, fmt.Errorf("cannot set output because node is not a DataCarrier")
		}
		output[i] = n.(DataCarrier).GetTensor()
	}
	return output, nil
}
