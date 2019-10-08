package onnx

import (
	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
)

func toOperationAttributes(attrs []*pb.AttributeProto) (map[string]interface{}, error) {
	output := make(map[string]interface{}, len(attrs))
	for _, attr := range attrs {
		o, err := toOperationAttribute(attr)
		if err != nil {
			return nil, err
		}
		output[attr.Name] = o
	}
	return output, nil
}

func toOperationAttribute(attr *pb.AttributeProto) (interface{}, error) {
	switch attr.GetType() {
	case pb.AttributeProto_UNDEFINED:
		return struct{}{}, nil
	case pb.AttributeProto_FLOAT:
		return attr.GetF(), nil
	case pb.AttributeProto_INT:
		return attr.GetI(), nil
	case pb.AttributeProto_STRING:
		return string(attr.GetS()), nil
	case pb.AttributeProto_TENSOR:
		return attr.GetT().Tensor()
	case pb.AttributeProto_GRAPH:
		return nil, &ErrNotImplemented{
			AttributeName:  attr.Name,
			AttributeValue: attr,
			Message:        "pb.AttributeProto_GRAPH not handled yet",
		}
	case pb.AttributeProto_FLOATS:
		return attr.GetFloats(), nil
	case pb.AttributeProto_INTS:
		return attr.GetInts(), nil
	case pb.AttributeProto_STRINGS:
		s := attr.GetStrings()
		strings := make([]string, len(s))
		for i := 0; i < len(s); i++ {
			strings[i] = string(s[i])
		}
		return strings, nil
	case pb.AttributeProto_TENSORS:
		return nil, &ErrNotImplemented{
			AttributeName:  attr.Name,
			AttributeValue: attr,
			Message:        "pb.AttributeProto_TENSORS not handled yet",
		}
	case pb.AttributeProto_GRAPHS:
		return nil, &ErrNotImplemented{
			AttributeName:  attr.Name,
			AttributeValue: attr,
			Message:        "pb.AttributeProto_GRAPHS not handled yet",
		}
	default:
		return nil, &ErrNotImplemented{
			AttributeName:  attr.Name,
			AttributeValue: attr,
			Message:        "undefined attributeproto type",
		}
	}
}
