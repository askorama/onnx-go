package images

import (
	"errors"
	"image"
	"image/color"

	"gorgonia.org/tensor"
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
		output = image.NewGray(rect)
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
			//log.Printf("Setting %#v at %v,%v\n", color, x, y)
			output.Set(x, y, color)
		}
	}
	return output, nil
}
