package backend

import "github.com/owulveryck/onnx-go"

// ComputationBackend is a backend that can run the graph
type ComputationBackend interface {
	onnx.Backend
	Run() error
}
