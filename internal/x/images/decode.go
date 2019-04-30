package images

import (
	"image"

	"gorgonia.org/tensor"
)

// TensorToImg turn a CHW tensor into an image (BCHW with B=1)
func TensorToImg(tensor.Tensor) (image.Image, error) {
	return nil, nil
}
