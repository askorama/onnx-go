package onnx

import (
	"github.com/gogo/protobuf/proto"
	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"gonum.org/v1/gonum/graph"
	"gorgonia.org/tensor"
)

// Unmarshal onnx encoded model proto data into a  weigthed directe builder
func Unmarshal(data []byte, dst graph.DirectedWeightedBuilder) error {
	model := &pb.ModelProto{}
	err := proto.Unmarshal(data, model)
	if err != nil {
		return err
	}
	return unmarshal(model, dst)
}

func unmarshal(model *pb.ModelProto, dst graph.DirectedWeightedBuilder) error {
	db := make(map[string]graph.Node, len(model.Graph.Output)+len(model.Graph.Input))
	for _, io := range append(model.Graph.Input, model.Graph.Output...) {
		n := dst.NewNode()
		db[io.Name] = n
		if _, ok := n.(Namer); ok {
			n.(Namer).SetName(io.Name)
		}
		if _, ok := n.(Tensor); ok {
			ttype := io.Type.GetTensorType()
			shape := make([]int, len(ttype.Shape.Dim))
			for i, d := range ttype.Shape.Dim {
				shape[i] = int(d.GetDimValue())
			}
			dtype, err := pb.TensorProto_DataType(ttype.GetElemType()).Dtype()
			if err != nil {
				return err
			}
			t := tensor.New(tensor.WithShape(shape...), tensor.Of(dtype))
			n.(Tensor).SetValue(t)

		}
		dst.AddNode(n)
	}
	for _, node := range model.Graph.Node {
		for _, output := range node.Output {
			var ok bool
			var no graph.Node
			if no, ok = db[output]; !ok {
				return &ErrInvalidModel{
					NodeNotDefined: output,
				}
			}
			if _, ok := no.(Op); ok {
				no.(Op).SetOpType(node.OpType)
				attrs := make([]*Attribute, len(node.Attribute))
				for i, a := range node.Attribute {
					attrs[i] = &Attribute{
						Key: a.Name,
						//Value: a.GetValue(),
					}
					err := no.(Op).SetOpAttributes(attrs)
					if err != nil {
						return err
					}
				}
			}
			// input should be ordered for non-commutatives operations
			for i, input := range node.Input {
				var ni graph.Node
				var ok bool
				if ni, ok = db[input]; !ok {
					return &ErrInvalidModel{
						NodeNotDefined: input,
					}
				}
				e := dst.NewWeightedEdge(no, ni, float64(i))
				dst.SetWeightedEdge(e)
			}
		}
	}
	return nil
}
