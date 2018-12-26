package pb

import (
	"reflect"
)

const (
	// AttrTagName ...
	AttrTagName = "attributeName"
	// RequiredTagName ...
	RequiredTagName = "required"
)

// UnmarshalAttributes reads the array of attributes and stores the result in the struct pointed to by v. If v is nil or not a pointer to a struct, Unmarshal returns an InvalidUnmarshalError.
// The structure pointed by v can only be flat and composed of one of the following types:
//    * string
//    * []string
//    * int64
//    * []int64
//    * float32
//    * []float32
//    * tensor.Tensor
//
// The values are associated thanks to the `onnx` tag fields and `required` tag if needed.
// Warning: any attribute not present in the v structure is silently discarded
func UnmarshalAttributes(attrs []*AttributeProto, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}

	rvi := reflect.Indirect(rv)
	if rvi.Kind() != reflect.Struct {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}
	attrsInventory := make(map[string]*AttributeProto, len(attrs))
	for _, attr := range attrs {
		attrsInventory[*attr.Name] = attr
	}
	for i := 0; i < rvi.NumField(); i++ {
		onnxTag, ok := rvi.Type().Field(i).Tag.Lookup(AttrTagName)
		if ok {
			required := false
			req, ok := rvi.Type().Field(i).Tag.Lookup(RequiredTagName)
			if ok && req == "true" {
				required = true

			}
			attr, ok := attrsInventory[onnxTag]
			if !ok {
				if required {
					return &RequiredTagNotFound{
						TagName: onnxTag,
					}
				}
				continue
			}
			// Process the attribute
			// Check if the attribute type match the type of the field
			switch *attr.Type {
			case AttributeProto_UNDEFINED:
				return &UnmarshalTypeError{
					Type: reflect.TypeOf(rvi.Field(i)),
				}
			case AttributeProto_FLOAT:
				kind := rvi.Field(i).Kind()
				if kind != reflect.Float64 {
					return &UnmarshalTypeError{
						Type: reflect.TypeOf(rvi.Field(i)),
					}
				}
				rvi.Field(i).SetFloat(float64(*attr.F))
			case AttributeProto_INT:
				if rvi.Field(i).Kind() != reflect.Int64 {
					return &UnmarshalTypeError{
						Type: reflect.TypeOf(rvi.Field(i)),
					}
				}
				rvi.Field(i).SetInt(*attr.I)
			case AttributeProto_STRING:
				if rvi.Field(i).Kind() != reflect.String {
					return &UnmarshalTypeError{
						Type: reflect.TypeOf(rvi.Field(i)),
					}
				}
				rvi.Field(i).SetString(string(attr.S))
			case AttributeProto_TENSOR:
				if rvi.Field(i).Kind() != reflect.Interface {
					return &UnmarshalTypeError{
						Type: reflect.TypeOf(rvi.Field(i)),
					}
				}
				t, err := attr.T.Tensor()
				if err != nil {
					return err
				}
				rvi.Field(i).Set(reflect.ValueOf(t))
			case AttributeProto_GRAPH:
				return &UnmarshalTypeError{
					Type: reflect.TypeOf(rvi.Field(i)),
				}
			case AttributeProto_FLOATS:
				if rvi.Field(i).Kind() != reflect.Slice {
					return &UnmarshalTypeError{
						Type: reflect.TypeOf(rvi.Field(i)),
					}
				}
				/*
					if rvi.Field(i).Elem().Kind() != reflect.Float32 {
						return &UnmarshalTypeError{}
					}
				*/
				rvi.Field(i).Set(reflect.ValueOf(attr.Floats))
			case AttributeProto_INTS:
				if rvi.Field(i).Kind() != reflect.Slice {
					return &UnmarshalTypeError{
						Type: reflect.TypeOf(rvi.Field(i)),
					}
				}
				/*
					if rvi.Field(i).Kind() != reflect.Int64 {
						return &UnmarshalTypeError{}
					}
				*/
				rvi.Field(i).Set(reflect.ValueOf(attr.Ints))
			case AttributeProto_STRINGS:
				if rvi.Field(i).Kind() != reflect.Slice {
					return &UnmarshalTypeError{
						Type: reflect.TypeOf(rvi.Field(i)),
					}
				}
				/*
					if rvi.Field(i).Elem().Kind() != reflect.String {
						return &UnmarshalTypeError{}
					}
				*/
			case AttributeProto_TENSORS:
				return &UnmarshalTypeError{}
			case AttributeProto_GRAPHS:
				return &UnmarshalTypeError{}
			default:
				return &UnmarshalTypeError{}
			}
		}
	}
	return nil
}

// A RequiredTagNotFound error is raised if a TagName is marked as required and is not found in the attribute list
type RequiredTagNotFound struct {
	TagName string
}

func (e *RequiredTagNotFound) Error() string {
	return "onnx: tag " + e.TagName + " is required but not found in the attribute list"
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
