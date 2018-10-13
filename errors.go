package onnx

import "github.com/pkg/errors"

var (
	// ErrNotYetImplemented ...
	ErrNotYetImplemented = errors.New("Not Yet Implemented")
	// ErrNoDataFound ...
	ErrNoDataFound = errors.New("No data found")
	// ErrCorruptedData ...
	ErrCorruptedData = errors.New("Unable to decode data")
)
