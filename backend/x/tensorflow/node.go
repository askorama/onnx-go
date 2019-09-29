package tfrt

import (
	"errors"
	"fmt"

	"github.com/owulveryck/onnx-go"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"

	"gorgonia.org/tensor"
	"gorgonia.org/tensor/native"
)

// Node is compatible with graph.Node and onnx.DataCarrier
type Node struct {
	id        int64
	t         tensor.Tensor
	operation *onnx.Operation
	name      string
	// gorgoniaNode stores a pointer to the node of the exprgraph
	opSpec *tf.OpSpec
}

// ID to fulfill the graph.Node interface
func (n *Node) ID() int64 {
	return n.id
}

// SetTensor assign the tensor N to the underlying node
func (n *Node) SetTensor(t tensor.Tensor) error {
	n.t = t
	if n.opSpec != nil {
		if n.opSpec.Type != "Placeholder" {
			return errors.New("Cannot set a tensor to a non placeholder node")
		}
		if !t.IsNativelyAccessible() {
			return errors.New("Cannot set tensor, tensor is not natively accessible")
		}
		var err error
		var backend interface{}
		switch t.Dtype() {
		case tensor.Float32:
			switch len(t.Shape()) {
			case 1:
				backend, err = native.VectorF32(t.(*tensor.Dense))
				if err != nil {
					return err
				}
			case 2:
				backend, err = native.MatrixF32(t.(*tensor.Dense))
				if err != nil {
					return err
				}
			case 3:
				backend, err = native.Tensor3F32(t.(*tensor.Dense))
				if err != nil {
					return err
				}
			case 4:
				if t.Shape()[0] != 1 {
					return fmt.Errorf("Tensor with shape %v not yet handled by this backend", t.Shape())
				}
			default:
				return fmt.Errorf("Tensor with %v axis not yet handled by this backend", t.Dims())
				backend, err = native.Tensor3F32(t.(*tensor.Dense))
				if err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("Tensor of type %v not yet handled by this backend", t.Dtype())
		}
		tfTensor, err := tf.NewTensor(backend)
		if err != nil {
			return err
		}
		n.opSpec.Attrs["value"] = tfTensor
	}
	return nil
}

// GetTensor value from the node
func (n *Node) GetTensor() tensor.Tensor {
	return n.t
}

// GetName get the name of the node
func (n *Node) GetName() string {
	return n.name
}

// SetName set the name of the node
func (n *Node) SetName(name string) {
	n.name = name
}
