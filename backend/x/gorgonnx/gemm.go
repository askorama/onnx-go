package gorgonnx

import (
	"errors"

	"github.com/owulveryck/onnx-go"
)

func init() {
	register("Gemm", &gemm{})
}

type gemm struct {
	alpha  float32 // Scalar multiplier for the product of input tensors A * B. default is 1.0
	beta   float32 // Scalar multiplier for input tensor C. default is 1.0
	transA bool    // Whether A should be transposed
	transB bool    // Whether B should be transposed
}

func (m *gemm) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 3)
	if err != nil {
		return err
	}
	return &onnx.ErrNotImplemented{
		Operator: "gemm",
	}
}

func (m *gemm) init(o onnx.Operation) error {
	m.alpha = 1.0
	m.beta = 1.0
	if alpha, ok := o.Attributes["alpha"]; ok {
		if alpha, ok := alpha.(float32); ok {
			m.alpha = alpha
		} else {
			return errors.New("Gemm: alpha is not a float32")
		}
	}
	if beta, ok := o.Attributes["beta"]; ok {
		if beta, ok := beta.(float32); ok {
			m.beta = beta
		} else {
			return errors.New("Gemm: beta is not a float32")
		}
	}
	if transA, ok := o.Attributes["transA"]; ok {
		if transA, ok := transA.(int64); ok {
			if transA == 1 {
				m.transA = true
			}
		} else {
			return errors.New("Gemm: transA is not an int")
		}
	}
	if transB, ok := o.Attributes["transB"]; ok {
		if transB, ok := transB.(int64); ok {
			if transB == 1 {
				m.transB = true
			}
		} else {
			return errors.New("Gemm: transB is not an int")
		}
	}
	return nil
}
