package images

import (
	"testing"

	"gorgonia.org/tensor"
)

func TestEncodeDecode(t *testing.T) {
	happyTensor := tensor.New(tensor.WithShape(1, 64, 64), tensor.WithBacking(happyFace))
	_, err := TensorToImg(happyTensor)
	if err != nil {
		t.Fatal(err)
	}
}
