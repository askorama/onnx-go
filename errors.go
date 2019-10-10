package onnx

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

var (
	errModelIsNil = errors.New("Model is nil")
	errGraphIsNil = errors.New("Graph is nil")
	errGraphNoIO  = errors.New("Graph have no input or output")
	errEmptyGraph = errors.New("Graph is empty")
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
	return "onnx: operator " + e.Operator + " not implemented (" + e.Message + ")"
}

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
// (The argument to Unmarshal must be a non-nil pointer.)
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "onnx: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "onnx: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "onnx: Unmarshal(nil " + e.Type.String() + ")"
}
