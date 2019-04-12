# Backend

This directory contains (sample) backend implementation. 

Any backend implementing a computation method (symbolic or not) should fulfil the `ComputationBackend` interface:

[embedmd]:# (computation_backend.go /type ComputationBackend/ /}/)
```go
type ComputationBackend interface {
	onnx.Backend
	Run() error
}
```

Unit tests are exposed in the `testbackend` package.

