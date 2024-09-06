package gorgonnx

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/chewxy/hm"
	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

func init() {
	register("Gather", newGather)
}

func newGather() operator {
	return &gather{}
}

type gather struct {
	axis int64

	dataShape    tensor.Shape
	indicesShape tensor.Shape
}

func (g *gather) Arity() int {
	return 2
}

func (g *gather) Type() hm.Type {
	a := hm.TypeVariable('a')
	c := hm.TypeVariable('c')
	dataType := gorgonia.TensorType{Dims: len(g.dataShape), Of: a}
	indicesType := gorgonia.TensorType{Dims: len(g.indicesShape), Of: c}
	retType := gorgonia.TensorType{Dims: len(g.dataShape) + len(g.indicesShape) - 1, Of: a}
	return hm.NewFnType(dataType, indicesType, retType)
}

func gatherInferShape(axis int64, dataShape, indicesShape tensor.Shape) tensor.Shape {
	shape := make(tensor.Shape, 0, len(dataShape)+len(indicesShape)-1)
	shape = append(shape, dataShape[:axis]...)
	shape = append(shape, indicesShape...)
	shape = append(shape, dataShape[axis+1:]...)
	return shape
}

func (g *gather) InferShape(inputs ...gorgonia.DimSizer) (tensor.Shape, error) {
	if inputs[0] == nil || inputs[1] == nil {
		return nil, errors.New("gather: infershape failed, nil shape")
	}
	return gatherInferShape(g.axis, inputs[0].(tensor.Shape), inputs[1].(tensor.Shape)), nil
}

func doGather[T float32 | float64](axis int64, data, indices gorgonia.Value) (gorgonia.Value, error) {
	vals, ok := data.Data().([]T)
	if !ok {
		return nil, errors.New(fmt.Sprintf("expected []%T, but cannot cast", T(0)))
	}
	indexVals, ok := indices.Data().([]int64)
	if !ok {
		return nil, errors.New(fmt.Sprintf("expected []int64, but cannot cast %T", indices.Data()))
	}
	retVal := tensor.NewDense(data.Dtype(), gatherInferShape(axis, data.Shape(), indices.Shape()))

	var totalStartIndices, totalMidIndices, totalEndIndices int64 = 1, 1, 1
	for _, dim := range data.Shape()[:axis] {
		totalStartIndices *= int64(dim)
	}
	for _, dim := range indices.Shape() {
		totalMidIndices *= int64(dim)
	}
	for _, dim := range data.Shape()[axis+1:] {
		totalEndIndices *= int64(dim)
	}
	axisIndices := int64(data.Shape()[axis])

	var i, j, k int64
	for i = 0; i < totalStartIndices; i++ {
		for j = 0; j < totalMidIndices; j++ {
			replaceIdx := indexVals[j]
			for k = 0; k < totalEndIndices; k++ {
				value := vals[(i*axisIndices+replaceIdx)*totalEndIndices+k]
				newIdx := (i*totalMidIndices+j)*totalEndIndices + k
				retVal.Set(int(newIdx), value)
			}
		}
	}

	return retVal, nil
}

func (g *gather) Do(inputs ...gorgonia.Value) (gorgonia.Value, error) {
	if len(inputs) != g.Arity() {
		return nil, errors.New("gather: wrong number of arguments")
	}
	data, ok := inputs[0].(*tensor.Dense)
	if !ok {
		return nil, errors.New("gather: only dense are supported")

	}
	indices, ok := inputs[1].(*tensor.Dense)
	if !ok {
		return nil, errors.New("gather: only dense are supported")

	}
	switch data.Dtype() {
	case tensor.Float64:
		return doGather[float64](g.axis, data, indices)
	case tensor.Float32:
		return doGather[float32](g.axis, data, indices)
	default:
		return nil, errors.New("gather Unsupported type")
	}
}

func (g *gather) ReturnsPtr() bool {
	return false
}

func (g *gather) CallsExtern() bool {
	return false
}

func (g *gather) OverwritesInput() int {
	return -1
}

func (g *gather) WriteHash(h hash.Hash) {
	if err := binary.Write(h, binary.LittleEndian, []byte(`gather`)); err != nil {
		panic(err)
	}
}

func (g *gather) Hashcode() uint32 {
	h := fnv.New32a()
	g.WriteHash(h)
	return h.Sum32()
}

func (g *gather) String() string {
	return "gather"
}

func (g *gather) apply(gg *Graph, ns ...*Node) error {
	n := ns[0]
	var err error
	children := getOrderedChildren(gg.g, n)
	if err := checkCondition(children, 2); err != nil {
		return err
	}
	data := children[0]
	indices := children[1]
	g.dataShape = data.gorgoniaNode.Shape()
	g.indicesShape = indices.gorgoniaNode.Shape()
	n.gorgoniaNode, err = gorgonia.ApplyOp(g, data.gorgoniaNode, indices.gorgoniaNode)
	if err != nil {
		return err
	}

	return nil
}

func (g *gather) init(o onnx.Operation) error {
	g.axis = 0
	axis, ok := o.Attributes["axis"]
	if ok {
		if g.axis, ok = axis.(int64); !ok {
			return errors.New("axis is not an int64")
		}
	}
	return nil
}
