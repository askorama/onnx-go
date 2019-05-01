package images

import (
	"image"
	"image/png"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
	if grayImg.Bounds().Size().X != sampleGrayTensor.Shape[2] ||
		grayImg.Bounds().Size().Y != sampleGrayTensor.Shape[3] {
		t.Fatalf("Expected image size to be %v, but it's %v", sampleGrayTensor.Shape[2:], grayImg.Bounds().Size())
	}
	// Check size
	generatedT := tensor.New(tensor.WithShape(sampleGrayTensor.Shape...), tensor.Of(tensor.Float32))
	err = GrayToBCHW(grayImg, generatedT)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, sampleT.Shape(), generatedT.Shape())
	assert.Equal(t, sampleT.Data(), generatedT.Data())
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
