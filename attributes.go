package onnx

import (
	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
)

func toOperationAttributes(attrs []*pb.AttributeProto) (map[string]interface{}, error) {
	output := make(map[string]interface{}, len(attrs))
	for _, attr := range attrs {
		switch attr.GetType() {
		case pb.AttributeProto_UNDEFINED:
			output[attr.Name] = struct{}{}
		case pb.AttributeProto_FLOAT:
			output[attr.Name] = attr.GetF()
		case pb.AttributeProto_INT:
			output[attr.Name] = attr.GetI()
		case pb.AttributeProto_STRING:
			output[attr.Name] = string(attr.GetS())
		case pb.AttributeProto_TENSOR:
			t, err := attr.GetT().Tensor()
			if err != nil {
				return nil, err
			}
			output[attr.Name] = t
		case pb.AttributeProto_GRAPH:
			return nil, &ErrNotImplemented{
				AttributeName:  attr.Name,
				AttributeValue: attr,
				Message:        "pb.AttributeProto_GRAPH not handled yet",
			}
		case pb.AttributeProto_FLOATS:
			output[attr.Name] = attr.GetFloats()
		case pb.AttributeProto_INTS:
			output[attr.Name] = attr.GetInts()
		case pb.AttributeProto_STRINGS:
			s := attr.GetStrings()
			strings := make([]string, len(s))
			for i := 0; i < len(s); i++ {
				strings[i] = string(s[i])
			}
			output[attr.Name] = strings
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
	return output, nil
}
