package images

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"reflect"

	"gorgonia.org/tensor"
)

// ImageToBCHW convert an image to a BCHW tensor
// this function returns an error if:
//
//   - dst is not a pointer
//   - dst's shape is not 4
//   - dst' second dimension is not 1
//   - dst's third dimension != i.Bounds().Dy()
//   - dst's fourth dimension != i.Bounds().Dx()
//   - dst's type is not float32 or float64 (temporary)
func ImageToBCHW(img image.Image, dst tensor.Tensor) error {
	// check if tensor is a pointer
	rv := reflect.ValueOf(dst)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("cannot decode image into a non pointer or a nil receiver")
	}
	// check if tensor is compatible with BCHW (4 dimensions)
	if len(dst.Shape()) != 4 {
		return fmt.Errorf("Expected a 4 dimension tensor, but receiver has only %v", len(dst.Shape()))
	}
	// Check the batch size
	if dst.Shape()[0] != 1 {
		return errors.New("only batch size of one is supported")
	}
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	if dst.Shape()[2] != h || dst.Shape()[3] != w {
		return fmt.Errorf("cannot fit image into tensor; image is %v*%v but tensor is %v*%v", h, w, dst.Shape()[2], dst.Shape()[3])
	}
	switch dst.Dtype() {
	case tensor.Float32:
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				r, g, b, a := img.At(x, y).RGBA()
				if a != 65535 {
					return errors.New("transparency not handled")
				}
				err := dst.SetAt(float32(uint8(r/0x101)), 0, 0, y, x)
				if err != nil {
					return err
				}
				err = dst.SetAt(float32(uint8(g/0x101)), 0, 1, y, x)
				if err != nil {
					return err
				}
				err = dst.SetAt(float32(uint8(b/0x101)), 0, 2, y, x)
				if err != nil {
					return err
				}
			}
		}
	default:
		return fmt.Errorf("%v not handled yet", dst.Dtype())
	}
	return nil

}

// GrayToBCHW convert an image to a BCHW tensor
// this function returns an error if:
//
//   - dst is not a pointer
//   - dst's shape is not 4
//   - dst' second dimension is not 1
//   - dst's third dimension != i.Bounds().Dy()
//   - dst's fourth dimension != i.Bounds().Dx()
//   - dst's type is not float32 or float64 (temporary)
func GrayToBCHW(img *image.Gray, dst tensor.Tensor) error {
	// check if tensor is a pointer
	rv := reflect.ValueOf(dst)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("cannot decode image into a non pointer or a nil receiver")
	}
	// check if tensor is compatible with BCHW (4 dimensions)
	if len(dst.Shape()) != 4 {
		return fmt.Errorf("Expected a 4 dimension tensor, but receiver has only %v", len(dst.Shape()))
	}
	// Check the batch size
	if dst.Shape()[0] != 1 {
		return errors.New("only batch size of one is supported")
	}
	if dst.Shape()[1] != 1 {
		return errors.New("Cowardly refusing to insert a gray scale into a tensor with more than one channel")
	}
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	if dst.Shape()[2] != h || dst.Shape()[3] != w {
		return fmt.Errorf("cannot fit image into tensor; image is %v*%v but tensor is %v*%v", h, w, dst.Shape()[2], dst.Shape()[3])
	}
	switch dst.Dtype() {
	case tensor.Float32:
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				color := img.GrayAt(x, y)
				err := dst.SetAt(float32(color.Y), 0, 0, y, x)
				if err != nil {
					return err
				}
			}
		}
	case tensor.Float64:
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				color := img.GrayAt(x, y)
				err := dst.SetAt(float64(color.Y), x, y)
				if err != nil {
					return err
				}
			}
		}
	default:
		return fmt.Errorf("%v not handled yet", dst.Dtype())
	}
	return nil
}

// TensorToImg turn a BCHW tensor into an image (BCHW with B=1)
func TensorToImg(t tensor.Tensor) (image.Image, error) {
	type img interface {
		image.Image
		Set(x, y int, c color.Color)
	}
	var output img
	if len(t.Shape()) != 4 {
		return nil, errors.New("expected a BCHW")
	}
	if t.Shape()[0] != 1 {
		return nil, errors.New("unhandled tensor with batch size > 1")
	}
	s := t.Shape()
	c, h, w := s[1], s[2], s[3]
	var rect = image.Rect(0, 0, w, h)
	t3, err := toTensor3(t)
	if err != nil {
		return nil, err
	}
	switch c {
	case 1:
		output = image.NewGray(rect)
	case 3:
		output = image.NewNRGBA(rect)
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
