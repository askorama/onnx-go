package onnx

import (
	"github.com/owulveryck/onnx-go/internal/onnx/ir"
)

func toOperationAttributes(attrs []*ir.AttributeProto) (map[string]interface{}, error) {
	output := make(map[string]interface{}, len(attrs))
	for _, attr := range attrs {
		o, err := toOperationAttribute(attr)
		if err != nil {
			return nil, err
		}
		output[attr.GetName()] = o
	}
	return output, nil
}

func toOperationAttribute(attr *ir.AttributeProto) (interface{}, error) {
	switch attr.GetType() {
	case ir.AttributeProto_UNDEFINED:
		return struct{}{}, nil
	case ir.AttributeProto_FLOAT:
		return attr.GetF(), nil
	case ir.AttributeProto_INT:
		return attr.GetI(), nil
	case ir.AttributeProto_STRING:
		return string(attr.GetS()), nil
	case ir.AttributeProto_TENSOR:
		return attr.GetT().Tensor()
	case ir.AttributeProto_GRAPH:
		return nil, &ErrNotImplemented{
			AttributeName:  attr.GetName(),
			AttributeValue: attr,
			Message:        "ir.AttributeProto_GRAPH not handled yet",
		}
	case ir.AttributeProto_FLOATS:
		return attr.GetFloats(), nil
	case ir.AttributeProto_INTS:
		return attr.GetInts(), nil
	case ir.AttributeProto_STRINGS:
		s := attr.GetStrings()
		strings := make([]string, len(s))
		for i := 0; i < len(s); i++ {
			strings[i] = string(s[i])
		}
		return strings, nil
	case ir.AttributeProto_TENSORS:
		return nil, &ErrNotImplemented{
			AttributeName:  attr.GetName(),
			AttributeValue: attr,
			Message:        "ir.AttributeProto_TENSORS not handled yet",
		}
	case ir.AttributeProto_GRAPHS:
		return nil, &ErrNotImplemented{
			AttributeName:  attr.GetName(),
			AttributeValue: attr,
			Message:        "ir.AttributeProto_GRAPHS not handled yet",
		}
	default:
		return nil, &ErrNotImplemented{
			AttributeName:  attr.GetName(),
			AttributeValue: attr,
			Message:        "undefined attributeproto type",
		}
	}
}
