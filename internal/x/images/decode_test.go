package images

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorgonia.org/tensor"
)

func generateGreyImageForTesting(t *testing.T) *image.Gray {
	sampleT := tensor.New(tensor.WithShape(sampleGrayTensor.Shape...), tensor.WithBacking(sampleGrayTensor.Data))
	decodedImg, err := TensorToImg(sampleT)
	assert.NoError(t, err)
	grayImg, ok := decodedImg.(*image.Gray)
	if !ok {
		t.Fail()
	}
	return grayImg
}

func Test_GrayToBCHW_NilReference(t *testing.T) {
	grayImg := generateGreyImageForTesting(t)
	err := GrayToBCHW(grayImg, nil)
	assert.EqualError(t, err, "cannot decode image into a non pointer or a nil receiver")
}

func Test_GrayToBCHW_Not4DimensionTensor(t *testing.T) {
	grayImg := generateGreyImageForTesting(t)
	not4DimensionTensor := tensor.New(tensor.Of(tensor.Float32), tensor.WithShape(1, 64, 64))
	err := GrayToBCHW(grayImg, not4DimensionTensor)
	assert.EqualError(t, err, "Expected a 4 dimension tensor, but receiver has only 3")
}

func Test_GrayToBCHW_BigBatchSizeTensor(t *testing.T) {
	grayImg := generateGreyImageForTesting(t)
	bigBatchSizeTensor := tensor.New(tensor.Of(tensor.Float32), tensor.WithShape(2, 1, 64, 64))
	err := GrayToBCHW(grayImg, bigBatchSizeTensor)
	assert.EqualError(t, err, "only batch size of one is supported")
}

func Test_GrayToBCHW_SmallerSizeTensor(t *testing.T) {
	grayImg := generateGreyImageForTesting(t)
	smallerSizeTensor := tensor.New(tensor.Of(tensor.Float32), tensor.WithShape(1, 1, 1, 1))
	err := GrayToBCHW(grayImg, smallerSizeTensor)
	assert.EqualError(t, err, "cannot fit image into tensor; image is 64*64 but tensor is 1*1")
}

func Test_GrayToBCHW_CowardGrayScale(t *testing.T) {
	grayImg := generateGreyImageForTesting(t)
	smallSizeTensor := tensor.New(tensor.Of(tensor.Float32), tensor.WithShape(1, 2, 64, 64))
	err := GrayToBCHW(grayImg, smallSizeTensor)
	assert.EqualError(t, err, "Cowardly refusing to insert a gray scale into a tensor with more than one channel")
}

func Test_GrayToBCHW_UsingNotSupportedTypeTensor(t *testing.T) {
	grayImg := generateGreyImageForTesting(t)
	notSupportedTypeTensor := tensor.New(tensor.Of(tensor.Int32), tensor.WithShape(1, 1, 64, 64))
	err := GrayToBCHW(grayImg, notSupportedTypeTensor)
	assert.EqualError(t, err, "int32 not handled yet")
}

func Test_ImageToBCHW_NilReference(t *testing.T) {
	img := generateImageForTesting()
	err := ImageToBCHW(img, nil)
	assert.EqualError(t, err, "cannot decode image into a non pointer or a nil receiver")
}

func Test_ImageToBCHW_Not4DimensionTensor(t *testing.T) {
	img := generateImageForTesting()
	not4DimensionTensor := tensor.New(tensor.Of(tensor.Float32), tensor.WithShape(1, 240, 280))
	err := ImageToBCHW(img, not4DimensionTensor)
	assert.EqualError(t, err, "Expected a 4 dimension tensor, but receiver has only 3")
}

func Test_ImageToBCHW_BigBatchSizeTensor(t *testing.T) {
	img := generateImageForTesting()
	bigBatchSizeTensor := tensor.New(tensor.Of(tensor.Float32), tensor.WithShape(2, 1, 240, 280))
	err := ImageToBCHW(img, bigBatchSizeTensor)
	assert.EqualError(t, err, "only batch size of one is supported")
}

func Test_ImageToBCHW_SmallerSizeTensor(t *testing.T) {
	img := generateImageForTesting()
	smallerSizeTensor := tensor.New(tensor.Of(tensor.Float32), tensor.WithShape(1, 1, 1, 1))
	err := ImageToBCHW(img, smallerSizeTensor)
	assert.EqualError(t, err, "cannot fit image into tensor; image is 240*280 but tensor is 1*1")
}

func Test_ImageToBCHW_UsingNotSupportedTypeTensor(t *testing.T) {
	img := generateImageForTesting()
	notSupportedTypeTensor := tensor.New(tensor.Of(tensor.Int32), tensor.WithShape(1, 1, 240, 280))
	err := ImageToBCHW(img, notSupportedTypeTensor)
	assert.EqualError(t, err, "int32 not handled yet")
}

func Test_TensorToImg_NotBCHW(t *testing.T) {
	sampleT := tensor.New(tensor.WithShape(1, 64, 64), tensor.WithBacking(sampleGrayTensor.Data))
	_, err := TensorToImg(sampleT)
	assert.EqualError(t, err, "expected a BCHW")
}
