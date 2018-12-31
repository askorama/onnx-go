package onnx

import (
	"github.com/gogo/protobuf/proto"
	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"gonum.org/v1/gonum/graph"
	"gorgonia.org/tensor"
)

// Attribute ...
type Attribute struct {
	Key   string
	Value interface{}
}

// Node ...
type Node interface {
	graph.Node
	SetName(string)
}

// Tensor ...
type Tensor interface {
	Node
	SetValue(t tensor.Tensor) error
	GetValue() tensor.Tensor
}

// Op is a node that represents an operation
type Op interface {
	Node
	SetOpType(string)
	SetOpAttributes([]*Attribute) error
}

// Unmarshal onnx encoded model proto data into a graph builder
func Unmarshal(data []byte, dst graph.DirectedBuilder) error {
	model := &pb.ModelProto{}
	err := proto.Unmarshal(data, model)
	if err != nil {
		return err
	}
	return unmarshal(model, dst)
}

func unmarshal(model *pb.ModelProto, dst graph.DirectedBuilder) error {
	db := make(map[string]graph.Node)
	for _, io := range append(model.Graph.Output, model.Graph.Input...) {
		n := dst.NewNode()
		db[io.Name] = n
		if _, ok := n.(Node); ok {
			n.(Node).SetName(io.Name)
		}
		if _, ok := n.(Tensor); ok {
			ttype := io.Type.GetTensorType()
			shape := make([]int, len(ttype.Shape.Dim))
			for i, d := range ttype.Shape.Dim {
				shape[i] = int(d.GetDimValue())
			}
			// TODO: generate the type
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
		n := dst.NewNode()
		db[node.Name] = n
		if _, ok := n.(Node); ok {
			n.(Node).SetName(node.Name)
		}
		if _, ok := n.(Op); ok {
			n.(Op).SetOpType(node.OpType)
			attrs := make([]*Attribute, len(node.Attribute))
			for i, a := range node.Attribute {
				attrs[i] = &Attribute{
					Key: a.Name,
					//Value: a.GetValue(),
				}
				err := n.(Op).SetOpAttributes(attrs)
				if err != nil {
					return err
				}
			}
			dst.AddNode(n)
		}
		for _, output := range node.Output {
			var ok bool
			var no graph.Node
			if no, ok = db[output]; !ok {
				return &ErrInvalidModel{
					NodeNotDefined: output,
				}
			}
			e := dst.NewEdge(n, no)
			dst.SetEdge(e)
		}
		// input should be ordered for non-commutatives operations
		for _, input := range node.Input {
			var ni graph.Node
			var ok bool
			if ni, ok = db[input]; !ok {
				return &ErrInvalidModel{
					NodeNotDefined: input,
				}
			}
			e := dst.NewEdge(ni, n)
			dst.SetEdge(e)
		}
	}
	return nil
}
