package images

import (
	"errors"
	"image"
	"image/color"

	"gorgonia.org/tensor"
	"gorgonia.org/tensor/native"
)

// ImgToBCHW convert an image to a BCHW tensor
func ImgToBCHW(i image.Image, dt tensor.Dtype) (tensor.Tensor, error) {
	w := i.Bounds().Dx()
	h := i.Bounds().Dy()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
		}
	}
	return nil, nil
}

// TensorToImg turn a CHW tensor into an image (BCHW with B=1)
func TensorToImg(t tensor.Tensor) (image.Image, error) {
	type img interface {
		image.Image
		Set(x, y int, c color.Color)
	}
	var output img
	s := t.Shape()
	c, h, w := s[0], s[1], s[2]
	var rect = image.Rect(0, 0, w, h)
	t3, err := toTensor3(t)
	if err != nil {
		return nil, err
	}
	switch c {
	case 1:
		output = image.NewGray16(rect)
	case 3:
		output = image.NewRGBA(rect)
	default:
		return nil, errors.New("unhandled image encoding")
	}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			color, err := t3.getColor(y, x)
			if err != nil {
				return nil, err
			}
			output.Set(x, y, color)
		}
	}
	return output, nil
}

func toTensor3(t tensor.Tensor) (*tensor3, error) {
	if len(t.Shape()) != 3 {
		return nil, errors.New("TensorToImg: expected a 3D tensor (CHW)")
	}
	dense, ok := t.(*tensor.Dense)
	if !ok {
		return nil, errors.New("This function can only convert dense tensors")
	}
	if f32, err := native.Tensor3F32(dense); err == nil {
		return &tensor3{
			c:   t.Shape()[0],
			h:   t.Shape()[1],
			w:   t.Shape()[2],
			f32: f32,
		}, nil
	}
	if f64, err := native.Tensor3F64(dense); err == nil {
		return &tensor3{
			c:   t.Shape()[0],
			h:   t.Shape()[1],
			w:   t.Shape()[2],
			f64: f64,
		}, nil
	}

	if i32, err := native.Tensor3I32(dense); err == nil {
		return &tensor3{
			c:   t.Shape()[0],
			h:   t.Shape()[1],
			w:   t.Shape()[2],
			i32: i32,
		}, nil
	}

	if i64, err := native.Tensor3I64(dense); err == nil {
		return &tensor3{
			c:   t.Shape()[0],
			h:   t.Shape()[1],
			w:   t.Shape()[2],
			i64: i64,
		}, nil
	}

	return nil, errors.New("cannot convert to tensor3")
}

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
		y, err := t.getUint16(0, h, w)
		return color.Gray16{
			Y: y,
		}, err
	case 3:
		r, err := t.getUint16(0, h, w)
		if err != nil {
			return nil, err
		}
		g, err := t.getUint16(1, h, w)
		if err != nil {
			return nil, err
		}
		b, err := t.getUint16(2, h, w)
		if err != nil {
			return nil, err
		}
		return color.RGBA64{
			R: r,
			G: g,
			B: b,
			A: uint16(255),
		}, nil
	default:
		return nil, errors.New("unhandled number of channel")
	}
}
