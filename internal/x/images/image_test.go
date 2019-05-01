package images

import (
	"image"
	"image/png"
	"os"
	"testing"

	"gorgonia.org/tensor"
)

func TestEncodeDecode(t *testing.T) {
	sampleT := tensor.New(tensor.WithShape(sampleGrayTensor.Shape...), tensor.WithBacking(sampleGrayTensor.Data))
	decodedImg, err := TensorToImg(sampleT)
	if err != nil {
		t.Fatal(err)
	}
	grayImg, ok := decodedImg.(*image.Gray)
	if !ok {
		t.Fail()
	}
	// Check size
	generatedT := tensor.New(tensor.WithShape(sampleGrayTensor.Shape...), tensor.Of(tensor.Float32))
	err = GrayToBCHW(grayImg, generatedT)
	if err != nil {
		t.Fatal(err)
	}
}

func savePic(img image.Image) error {
	f, err := os.Create("image.png")
	if err != nil {
		return err
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
