package onnx

import pb "github.com/owulveryck/onnx-go/internal/pb-onnx"

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
func UnmarshalAttributes(attrs []*pb.AttributeProto, v interface{}) error {
	return pb.UnmarshalAttributes(attrs, v)
}
