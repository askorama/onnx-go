package onnx

import (
	"reflect"

	"github.com/gogo/protobuf/proto"
	pb "github.com/owulveryck/onnx-go/internal/pb-onnx"
	"github.com/pkg/errors"
	"gonum.org/v1/gonum/graph"
	"gorgonia.org/tensor"
)

// Model is a wrapper around a computation graph.
// Input and Output are containing the ID of the corresponding nodes.
type Model struct {
	backend  Backend
	dbByName map[string]graph.Node
	Input    []int64
	Output   []int64
}

// NewModel with dst as backend.
// dst should be a non-nil pointer.
func NewModel(dst Backend) *Model {
	return &Model{
		backend: dst,
	}
}

// Decode the onnx binary data into the model
func (m *Model) Decode(data []byte) error {
	pbModel := &pb.ModelProto{}
	err := proto.Unmarshal(data, pbModel)
	if err != nil {
		return err
	}
	return m.decode(pbModel)
}

// GetNodeByName is a utility method that returns a node of the computation graph
func (m *Model) GetNodeByName(name string) (graph.Node, bool) {
	n, ok := m.dbByName[name]
	return n, ok
}

func (m *Model) processValue(io *pb.ValueInfoProto) (graph.Node, error) {
	var opts []tensor.ConsOpt
	dst := m.backend
	n := dst.NewNode()
	if _, ok := n.(Namer); ok {
		n.(Namer).SetName(io.Name)
	}
	dst.AddNode(n)
	m.dbByName[io.Name] = n
	if _, ok := n.(DataCarrier); !ok {
		return n, nil
	}
	ttype := io.Type.GetTensorType()
	if ttype.GetShape() != nil {
		var shape []int
		for i := range ttype.Shape.Dim {
			_, ok := ttype.Shape.Dim[i].GetValue().(*pb.TensorShapeProto_Dimension_DimValue)
			if ok {
				shape = append(shape, int(ttype.Shape.Dim[i].GetDimValue()))
			}
		}
		opts = append(opts, tensor.WithShape(shape...))
	}
	dtype, err := pb.TensorProto_DataType(ttype.GetElemType()).Dtype()
	if err != nil {
		return n, err
	}
	opts = append(opts, tensor.Of(dtype))
	t := tensor.New(opts...)
	err = n.(DataCarrier).SetTensor(t)
	if err != nil {
		return n, err
	}

	return n, nil
}

func (m *Model) decode(model *pb.ModelProto) error {
	rv := reflect.ValueOf(m.backend)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(m.backend)}
	}

	m.Input = make([]int64, len(model.Graph.Input))
	m.Output = make([]int64, len(model.Graph.Output))
	m.dbByName = make(map[string]graph.Node, len(model.Graph.Output)+len(model.Graph.Input))
	dst := m.backend
	// Well...
	for i, io := range model.Graph.Input {
		n, err := m.processValue(io)
		if err != nil {
			return err
		}
		m.Input[i] = n.ID()
	}
	for _, io := range model.Graph.ValueInfo {
		_, err := m.processValue(io)
		if err != nil {
			return err
		}
	}
	for i, io := range model.Graph.Output {
		n, err := m.processValue(io)
		if err != nil {
			return err
		}
		m.Output[i] = n.ID()
	}
	for _, tensorProto := range model.Graph.GetInitializer() {
		name := tensorProto.GetName()
		if name == "" {
			return errors.New("initializer should have a name")
		}
		n, ok := m.dbByName[name]
		if !ok {
			return errors.New("invalid model: initializer has not been defined in input, output or value")
		}
		// Remove it from the input
		// find the ID
		for i := 0; i < len(m.Input); i++ {
			if m.Input[i] == n.ID() {
				m.Input = append(m.Input[:i], m.Input[i+1:]...)
			}
		}
		if _, ok := n.(DataCarrier); !ok {
			continue
		}
		t, err := tensorProto.Tensor()
		if err != nil {
			return err
		}
		err = n.(DataCarrier).SetTensor(t)
		if err != nil {
			return err
		}
	}
	for _, node := range model.Graph.Node {
		for _, output := range node.Output {
			var ok bool
			var no graph.Node
			if no, ok = m.dbByName[output]; !ok {
				no = dst.NewNode()
				if _, ok := no.(Namer); ok {
					no.(Namer).SetName(output)
				}
				dst.AddNode(no)
				m.dbByName[output] = no
			}
			// input should be ordered for non-commutatives operations
			for i, input := range node.Input {
				var ni graph.Node
				var ok bool
				if ni, ok = m.dbByName[input]; !ok {
					ni = dst.NewNode()
					if _, ok := ni.(Namer); ok {
						ni.(Namer).SetName(input)
					}
					dst.AddNode(ni)
					m.dbByName[input] = ni
				}
				e := dst.NewWeightedEdge(no, ni, float64(i))
				dst.SetWeightedEdge(e)
			}
			// The graph can apply operations
			err := dst.ApplyOperation(Operation{
				node.OpType,
				toOperationAttributes(node.GetAttribute()),
			}, no)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
