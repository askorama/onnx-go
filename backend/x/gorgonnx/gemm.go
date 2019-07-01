package gorgonnx

import (
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/chewxy/hm"
	"github.com/owulveryck/onnx-go"
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/blas32"
	"gonum.org/v1/gonum/blas/blas64"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

func init() {
	register("Gemm", newGemm)
}

func newGemm() operator {
	return &gemm{
		alpha:  1.0,
		beta:   1.0,
		transA: false,
		transB: false,
	}
}

type gemm struct {
	alpha  float32 // Scalar multiplier for the product of input tensors A * B. default is 1.0
	beta   float32 // Scalar multiplier for input tensor C. default is 1.0
	transA bool    // Whether A should be transposed
	transB bool    // Whether B should be transposed
}

func (o *gemm) Arity() int {
	return 3
}

func (o *gemm) Type() hm.Type {
	t := gorgonia.TensorType{Dims: 2, Of: hm.TypeVariable('a')}
	return hm.NewFnType(t, t, t, t)
}

func (o *gemm) InferShape(inputs ...gorgonia.DimSizer) (tensor.Shape, error) {
	// output is (m,n)
	m, err := inputs[0].DimSize(0)
	if err != nil {
		return nil, err
	}
	if o.transA {
		m, err = inputs[0].DimSize(1)
		if err != nil {
			return nil, err
		}
	}
	n, err := inputs[1].DimSize(1)
	if err != nil {
		return nil, err
	}
	if o.transB {
		n, err = inputs[1].DimSize(0)
		if err != nil {
			return nil, err
		}
	}
	return []int{m, n}, nil
}

func (o *gemm) Do(inputs ...gorgonia.Value) (gorgonia.Value, error) {
	var a tensor.Tensor
	var ok bool
	if a, ok = inputs[0].(tensor.Tensor); !ok {
		return nil, errors.New("gemm: not a tensor")
	}
	switch a.Dtype() {
	case gorgonia.Float32:
		return o.do32(inputs...)
	case gorgonia.Float64:
		return o.do64(inputs...)
	default:
		return nil, errors.New("gemm: type not handled")
	}
}

func (o *gemm) do32(inputs ...gorgonia.Value) (gorgonia.Value, error) {
	var a, b, c tensor.Tensor
	var ok bool
	if a, ok = inputs[0].(tensor.Tensor); !ok {
		return nil, errors.New("gemm: not a tensor")
	}
	if b, ok = inputs[1].(tensor.Tensor); !ok {
		return nil, errors.New("gemm: not a tensor")
	}
	if c, ok = inputs[2].(tensor.Tensor); !ok {
		return nil, errors.New("gemm: not a tensor")
	}
	s, err := o.InferShape(a.Shape(), b.Shape(), c.Shape())
	if err != nil {
		return nil, err
	}
	m := s[0]
	n := s[1]
	if c.DataSize() != m*n {
		backend := make([]float32, m*n)
		switch c.DataSize() {
		case 0:
			for i := 0; i < len(backend); i++ {
				backend[i] = c.Data().(float32)
			}
		case 1:
			for i := 0; i < len(backend); i++ {
				backend[i] = c.Data().([]float32)[0]
			}
		case n:
			for i := 0; i < m; i++ {
				copy(backend[i:i+n], c.Data().([]float32))
			}
		default:
			return nil, fmt.Errorf("gemm: C has %v elements, but %vx%v are expected", c.DataSize(), m, n)
		}
		c = tensor.New(tensor.WithShape(m, n), tensor.WithBacking(backend))
	}
	transA := blas.NoTrans
	transB := blas.NoTrans
	if o.transA {
		transA = blas.Trans
	}
	if o.transB {
		transB = blas.Trans
	}
	// do we need to broadcast?
	blas32.Gemm(transA, transB, o.alpha,
		blas32.General{
			Rows:   a.Shape()[0],
			Cols:   a.Shape()[1],
			Stride: a.Strides()[0],
			Data:   a.Data().([]float32),
		},
		blas32.General{
			Rows:   b.Shape()[0],
			Cols:   b.Shape()[1],
			Stride: b.Strides()[0],
			Data:   b.Data().([]float32),
		},
		o.beta,
		blas32.General{
			Rows:   c.Shape()[0],
			Cols:   c.Shape()[1],
			Stride: c.Strides()[0],
			Data:   c.Data().([]float32),
		})

	return c, nil
}
func (o *gemm) do64(inputs ...gorgonia.Value) (gorgonia.Value, error) {
	var a, b, c tensor.Tensor
	var ok bool
	if a, ok = inputs[0].(tensor.Tensor); !ok {
		return nil, errors.New("gemm: not a tensor")
	}
	if b, ok = inputs[1].(tensor.Tensor); !ok {
		return nil, errors.New("gemm: not a tensor")
	}
	if c, ok = inputs[2].(tensor.Tensor); !ok {
		return nil, errors.New("gemm: not a tensor")
	}
	s, err := o.InferShape(a.Shape(), b.Shape(), c.Shape())
	if err != nil {
		return nil, err
	}
	m := s[0]
	n := s[1]
	if c.DataSize() != m*n {
		backend := make([]float64, m*n)
		switch c.DataSize() {
		case 0:
			for i := 0; i < len(backend); i++ {
				backend[i] = c.Data().(float64)
			}
		case 1:
			for i := 0; i < len(backend); i++ {
				backend[i] = c.Data().([]float64)[0]
			}
		case n:
			for i := 0; i < m; i++ {
				copy(backend[i:i+n], c.Data().([]float64))
			}
		default:
			return nil, fmt.Errorf("gemm: C has %v elements, but %vx%v are expected", c.DataSize(), m, n)
		}
		c = tensor.New(tensor.WithShape(m, n), tensor.WithBacking(backend))
	}
	transA := blas.NoTrans
	transB := blas.NoTrans
	if o.transA {
		transA = blas.Trans
	}
	if o.transB {
		transB = blas.Trans
	}
	// do we need to broadcast?
	blas64.Gemm(transA, transB, float64(o.alpha),
		blas64.General{
			Rows:   a.Shape()[0],
			Cols:   a.Shape()[1],
			Stride: a.Strides()[0],
			Data:   a.Data().([]float64),
		},
		blas64.General{
			Rows:   b.Shape()[0],
			Cols:   b.Shape()[1],
			Stride: b.Strides()[0],
			Data:   b.Data().([]float64),
		},
		float64(o.beta),
		blas64.General{
			Rows:   c.Shape()[0],
			Cols:   c.Shape()[1],
			Stride: c.Strides()[0],
			Data:   c.Data().([]float64),
		})

	return c, nil
}

func (*gemm) ReturnsPtr() bool {
	return true
}

func (*gemm) CallsExtern() bool {
	return false
}

func (*gemm) OverwritesInput() int {
	return 2
}

func (o *gemm) WriteHash(h hash.Hash) {
	fmt.Fprintf(h, "gemm-%v-%v-%v-%v", o.transA, o.transB, o.alpha, o.beta)
}

func (o *gemm) Hashcode() uint32 {
	h := fnv.New32a()
	o.WriteHash(h)
	return h.Sum32()
}

func (o *gemm) String() string {
	return fmt.Sprintf("gemm-%v-%v-%v-%v", o.transA, o.transB, o.alpha, o.beta)
}

// Compute Y = alpha * A' * B' + beta * C, where
//  * input tensor A has shape (M, K) or (K, M),
//  * input tensor B has shape (K, N) or (N, K),
//  * input tensor C is broadcastable to shape (M, N),
//  * output tensor Y has shape (M, N).
// A will be transposed before doing the computation if attribute transA is non-zero,
// same for B and transB.
// This operator supports unidirectional broadcasting i
// (tensor C should be unidirectional broadcastable to tensor A * B);
//
// https://github.com/onnx/onnx/blob/master/docs/Operators.md#Gemm
func (o *gemm) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 3)
	if err != nil {
		return err
	}
	a := children[0].gorgoniaNode
	b := children[1].gorgoniaNode
	c := children[2].gorgoniaNode
	n.gorgoniaNode, err = gorgonia.ApplyOp(o, a, b, c)
	return err
}

func (o *gemm) init(op onnx.Operation) error {
	o.alpha = 1.0
	o.beta = 1.0
	if alpha, ok := op.Attributes["alpha"]; ok {
		if alpha, ok := alpha.(float32); ok {
			o.alpha = alpha
		} else {
			return errors.New("Gemm: alpha is not a float32")
		}
	}
	if beta, ok := op.Attributes["beta"]; ok {
		if beta, ok := beta.(float32); ok {
			o.beta = beta
		} else {
			return errors.New("Gemm: beta is not a float32")
		}
	}
	if transA, ok := op.Attributes["transA"]; ok {
		if transA, ok := transA.(int64); ok {
			if transA == 1 {
				o.transA = true
			}
		} else {
			return errors.New("Gemm: transA is not an int")
		}
	}
	if transB, ok := op.Attributes["transB"]; ok {
		if transB, ok := transB.(int64); ok {
			if transB == 1 {
				o.transB = true
			}
		} else {
			return errors.New("Gemm: transB is not an int")
		}
	}
	return nil
}
