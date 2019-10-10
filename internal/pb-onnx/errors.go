package pb

import (
	"fmt"

	"github.com/pkg/errors"
)

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
	AttributeName  string
	AttributeValue interface{}
	Message        string
}

func (e *ErrNotImplemented) Error() string {
	if e.AttributeName != "" {
		return fmt.Sprintf("onnx: operator %v. Implementation error for attribute %v (%#v): %v",
			e.Operator,
			e.AttributeName,
			e.AttributeValue,
			e.Message)
	}
	return "onnx: operator " + e.Operator + " not implemented"
}
