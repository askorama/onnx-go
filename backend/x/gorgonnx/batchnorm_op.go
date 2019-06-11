package gorgonnx

import (
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/chewxy/hm"
	"github.com/pkg/errors"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
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
	panic("not implemented")
}

func (b *batchNorm) ReturnsPtr() bool {
	return true
}

func (b *batchNorm) CallsExtern() bool {
	return false
}

func (b *batchNorm) OverwritesInput() int {
	return -1
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
