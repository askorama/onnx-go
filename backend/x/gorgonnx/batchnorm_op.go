package gorgonnx

import (
	"fmt"
	"hash"
	"hash/fnv"
	"math"

	"github.com/chewxy/hm"
	"github.com/pkg/errors"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
	"gorgonia.org/tensor/native"
)

type batchNorm struct {
	scale, bias, mean, varN gorgonia.Value
	epsilon                 float32
}

func (b *batchNorm) Arity() int {
	return 1
}

func (b *batchNorm) Type() hm.Type {
	t := gorgonia.TensorType{Dims: 4, Of: hm.TypeVariable('a')}
	return hm.NewFnType(t, t)
}

func (b *batchNorm) InferShape(ns ...gorgonia.DimSizer) (tensor.Shape, error) {
	if len(ns) != b.Arity() {
		return nil, errors.New("wrong number of arguments for batchnorm")
	}

	return ns[0].(tensor.Shape).Clone(), nil
}

func (b *batchNorm) Do(values ...gorgonia.Value) (gorgonia.Value, error) {
	// xNorm = (x - meanN) / sqrt( varN + b.epsilon)
	// output = scaleN * xNorm + biasN
	if len(values) != b.Arity() {
		return nil, errors.New("wrong number of arguments for batchnorm_op")
	}
	x, ok := values[0].(*tensor.Dense)
	if !ok {
		return nil, errors.New("batchNorm_op only works on dense tensors")
	}

	if len(x.Shape()) != 4 {
		return nil, errors.New("batchNorm_op expects a BCHW tensor")
	}
	if x.Shape()[0] != 1 {
		return nil, errors.New("batchNorm_op expects a BCHW tensor with B=1")
	}
	if x.Dtype() != tensor.Float32 {
		panic("batchnorm_op: dense not []float32 not implemented")
	}
	if len(b.scale.Shape()) != 1 || len(b.bias.Shape()) != 1 ||
		len(b.mean.Shape()) != 1 || len(b.varN.Shape()) != 1 {
		return nil, errors.New("batchnorp_op only support 1 dim scale,bias,mean and var")
	}
	ch := x.Shape()[1]
	if b.scale.Shape()[0] != ch || b.bias.Shape()[0] != ch ||
		b.mean.Shape()[0] != ch || b.varN.Shape()[0] != ch {
		return nil, errors.New("batchnorp_op only support 1 dim scale,bias,mean and var")
	}
	// Reshape to CHW
	s := make([]int, len(x.Shape()))
	copy(s, x.Shape())
	err := x.Reshape(s[1:]...)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := x.Reshape(s...)
		if err != nil {
			panic(err)
		}
	}()
	// TODO: handle float64

	vals, err := native.Tensor3F32(x)
	if err != nil {
		return nil, err
	}
	// xNorm = (x - meanN) / sqrt( varN + b.epsilon)
	// output = scaleN * xNorm + biasN
	for c := 0; c < len(vals); c++ {
		mean := b.mean.Data().([]float32)[c]
		varV := b.varN.Data().([]float32)[c]
		scale := b.scale.Data().([]float32)[c]
		bias := b.bias.Data().([]float32)[c]
		for h := 0; h < len(vals[c]); h++ {
			for w := 0; w < len(vals[c][h]); w++ {
				x := vals[c][h][w]
				//outT3[c][h][w] = scale*((x-mean)/sqrtF32(varV+b.epsilon)) + bias
				vals[c][h][w] = scale*((x-mean)/sqrtF32(varV+b.epsilon)) + bias
			}
		}
	}
	return x, nil
}

func sqrtF32(v float32) float32 {
	return float32(math.Sqrt(float64(v)))
}

func (b *batchNorm) ReturnsPtr() bool {
	return true
}

func (b *batchNorm) CallsExtern() bool {
	return false
}

func (b *batchNorm) OverwritesInput() int {
	//return -1
	return 0
}

func (b *batchNorm) WriteHash(h hash.Hash) {
	fmt.Fprintf(h, "batchnorm-%1.1f", b.epsilon)
}

func (b *batchNorm) Hashcode() uint32 {
	h := fnv.New32a()
	b.WriteHash(h)
	return h.Sum32()
}

func (b *batchNorm) String() string {
	return fmt.Sprintf("batchnorm-%1.1f", b.epsilon)
}
