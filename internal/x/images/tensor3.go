package images

import (
	"errors"
	"image/color"

	"gorgonia.org/tensor"
	"gorgonia.org/tensor/native"
)

// dumb structure to avoid type asertion at runtime
type tensor3 struct {
	c   int
	h   int
	w   int
	f32 [][][]float32
	f64 [][][]float64
	i32 [][][]int32
	i64 [][][]int64
}

func toTensor3(t tensor.Tensor) (*tensor3, error) {
	if len(t.Shape()) != 4 {
		return nil, errors.New("TensorToImg: expected a 4D tensor (BCHW)")
	}
	if t.Shape()[0] != 1 {
		return nil, errors.New("batch >1 not implemented")
	}
	dense, ok := t.(*tensor.Dense)
	if !ok {
		return nil, errors.New("This function can only convert dense tensors")
	}
	originalShape := make([]int, 4)
	newShape := make([]int, 3)
	copy(originalShape, t.Shape())
	copy(newShape, t.Shape()[1:4])
	err := dense.Reshape(newShape...)
	if err != nil {
		return nil, err
	}
	defer func() {
		dense.Reshape(originalShape...)
	}()

	if f32, err := native.Tensor3F32(dense); err == nil {
		return &tensor3{
			c:   dense.Shape()[0],
			h:   dense.Shape()[1],
			w:   dense.Shape()[2],
			f32: f32,
		}, nil
	}
	if f64, err := native.Tensor3F64(dense); err == nil {
		return &tensor3{
			c:   dense.Shape()[0],
			h:   dense.Shape()[1],
			w:   dense.Shape()[2],
			f64: f64,
		}, nil
	}

	if i32, err := native.Tensor3I32(dense); err == nil {
		return &tensor3{
			c:   dense.Shape()[0],
			h:   dense.Shape()[1],
			w:   dense.Shape()[2],
			i32: i32,
		}, nil
	}

	if i64, err := native.Tensor3I64(dense); err == nil {
		return &tensor3{
			c:   dense.Shape()[0],
			h:   dense.Shape()[1],
			w:   dense.Shape()[2],
			i64: i64,
		}, nil
	}

	return nil, errors.New("cannot convert to tensor3")
}
func (t *tensor3) getUint8(c, h, w int) (uint8, error) {
	lc := t.c
	lh := t.h
	lw := t.w
	if c > lc || h > lh || w > lw {
		return 0, errors.New("request out of bound")
	}
	switch {
	case t.f32 != nil:
		return uint8(t.f32[c][h][w]), nil
	case t.f64 != nil:
		return uint8(t.f64[c][h][w]), nil
	case t.i32 != nil:
		return uint8(t.i32[c][h][w]), nil
	case t.i64 != nil:
		return uint8(t.i64[c][h][w]), nil
	}
	return 0, nil
}

func (t *tensor3) getUint16(c, h, w int) (uint16, error) {
	lc := t.c
	lh := t.h
	lw := t.w
	if c > lc || h > lh || w > lw {
		return 0, errors.New("request out of bound")
	}
	switch {
	case t.f32 != nil:
		return uint16(t.f32[c][h][w]), nil
	case t.f64 != nil:
		return uint16(t.f64[c][h][w]), nil
	case t.i32 != nil:
		return uint16(t.i32[c][h][w]), nil
	case t.i64 != nil:
		return uint16(t.i64[c][h][w]), nil
	}
	return 0, nil
}

func (t *tensor3) getColor(h, w int) (color.Color, error) {
	switch t.c {
	case 1:
		y, err := t.getUint8(0, h, w)
		return color.Gray{
			Y: y,
		}, err
	case 3:
		r, err := t.getUint8(0, h, w)
		if err != nil {
			return nil, err
		}
		g, err := t.getUint8(1, h, w)
		if err != nil {
			return nil, err
		}
		b, err := t.getUint8(2, h, w)
		if err != nil {
			return nil, err
		}
		return color.NRGBA{
			R: r,
			G: g,
			B: b,
			A: uint8(255),
		}, nil
	default:
		return nil, errors.New("unhandled number of channel")
	}
}
