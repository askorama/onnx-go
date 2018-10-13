package onnx

import (
	"reflect"
)

// UnmarshalAttributes reads the array of attributes and stores the result in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.
func UnmarshalAttributes(attrs []AttributeProto, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}
	return nil
}

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
// (The argument to Unmarshal must be a non-nil pointer.)
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "onnx: UnmarshalAttributes(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "onnx: UnmarshalAttributes(non-pointer " + e.Type.String() + ")"
	}
	return "onnx: UnmarshalAttributes(nil " + e.Type.String() + ")"
}

// An UnmarshalTypeError describes an attribute value that was
// not appropriate for a value of a specific Go type.
type UnmarshalTypeError struct {
	Value  string       // description of the ONNX attribute value - "bool", "array", "number -5"
	Type   reflect.Type // type of Go value it could not be assigned to
	Struct string       // name of the struct type containing the field
	Field  string       // name of the field holding the Go value
}

func (e *UnmarshalTypeError) Error() string {
	if e.Struct != "" || e.Field != "" {
		return "onnx: cannot unmarshal " + e.Value + " into Go struct field " + e.Struct + "." + e.Field + " of type " + e.Type.String()
	}
	return "onnx: cannot unmarshal " + e.Value + " into Go value of type " + e.Type.String()
}
