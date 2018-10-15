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

// ErrNotImplemented is returned for any operator or attribute
type ErrNotImplemented struct {
	Operator       string
	AttributeName  *string
	AttributeValue interface{}
	Message        string
}

func (e *ErrNotImplemented) Error() string {
	return "onnx: operator " + e.Operator + " not implemented"
}
