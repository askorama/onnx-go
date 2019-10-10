package ir

import (
	"fmt"

	"github.com/pkg/errors"
	"gorgonia.org/tensor"
)

// Dtype returns a compatible dtype from the source
func (t TensorProto_DataType) Dtype() (tensor.Dtype, error) {
	switch t {
	case TensorProto_FLOAT:
		// As defined in the spec, a float is a 32 floating-point value
		return tensor.Float32, nil
	case TensorProto_UINT8:
		return tensor.Uint8, nil
	case TensorProto_INT8:
		return tensor.Int8, nil
	case TensorProto_UINT16:
		return tensor.Uint16, nil
	case TensorProto_INT16:
		return tensor.Int16, nil
	case TensorProto_INT32:
		return tensor.Int32, nil
	case TensorProto_INT64:
		return tensor.Int64, nil
	case TensorProto_STRING:
		return tensor.String, nil
	case TensorProto_BOOL:
		return tensor.Bool, nil
		// Advanced types
	case TensorProto_DOUBLE:
		// BUG(): a double type is replaced by a Float64 on all plateforms.
		return tensor.Float64, nil
	case TensorProto_UINT32:
		return tensor.Uint32, nil
	case TensorProto_UINT64:
		return tensor.Uint64, nil
	case TensorProto_COMPLEX64:
		return tensor.Complex64, nil
	case TensorProto_COMPLEX128:
		return tensor.Complex128, nil
	}
	return wrapUnknownTensorDtype(t)
}

func wrapUnknownTensorDtype(t TensorProto_DataType) (tensor.Dtype, error) {
	switch t {
	case TensorProto_FLOAT16:
		return tensor.Dtype{}, errors.Wrapf(ErrNotYetImplemented, "type: %v", t)
	}
	return tensor.Dtype{}, fmt.Errorf("Unknown input type: %v", t)
}
