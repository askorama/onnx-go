package onnx

import (
	"github.com/gogo/protobuf/proto"
	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"github.com/pkg/errors"
	"gonum.org/v1/gonum/graph"
	"gorgonia.org/tensor"
)

// Unmarshal onnx encoded model proto data into a weigthed directed builder
// The weight of the edges represents the indicies of the children (therefore their order)
// first child has a weight of 0
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
	// Well...
	for _, io := range append(append(model.Graph.Input, model.Graph.ValueInfo...), model.Graph.Output...) {
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
			// The graph can apply operations
			if _, ok := dst.(OperationApplyer); ok {
				op, err := dst.(OperationApplyer).ONNXGetOperationFromName(node.OpType)
				if err != nil {
					return err
				}
				err = pb.UnmarshalAttributes(node.GetAttribute(), &op)
				operation, ok := op.(Operation)
				if !ok {
					return errors.New("Graph builder did not return an operation")
				}
				err = dst.(OperationApplyer).ONNXApply(operation.Constructor(), no)
				if err != nil {
					return err
				}

			}
		}
	}
	return nil
}

// OperationApplyer is any graph that can apply operations on its node
// regarding the structure of the graph
type OperationApplyer interface {
	// ONNXGetOperationFromName returns an interface that should be compatible with Operation
	// It is the responsibility of the implementor to do that
	ONNXGetOperationFromName(s string) (interface{}, error)
	ONNXApply(operation func(g graph.WeightedDirected, n graph.Node) (interface{}, error), n graph.Node) error
}

// Operation is an interface that should be fulfiled by any Operation
type Operation interface {
	// Constructor returns a function that itself returns an operator to be applied on node n.
	// The arguments of the operator are found thanks to the graph
	Constructor() func(g graph.WeightedDirected, n graph.Node) (interface{}, error)
}
