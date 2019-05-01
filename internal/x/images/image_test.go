package images

import (
	"image/png"
	"os"
	"testing"

	"gorgonia.org/tensor"
)

func TestEncodeDecode(t *testing.T) {
	happyTensor := tensor.New(tensor.WithShape(1, 64, 64), tensor.WithBacking(happyFace))
	img, err := TensorToImg(happyTensor)
	if err != nil {
		t.Fatal(err)
	}
	f, err := os.Create("image.png")
	if err != nil {
		t.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		t.Fatal(err)
	}

	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

}
