// +build !wasm

package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"

	"github.com/disintegration/imaging"
	"github.com/vincent-petithory/dataurl"
)

func displayResult(r results) {
	fmt.Printf("%v / %2.2f%%\n", r[0].result, r[0].weight*100)
	fmt.Printf("%v / %2.2f%%\n", r[1].result, r[1].weight*100)
}

func processPicture(data string, height, width int) (*image.Gray, error) {
	dataURL, err := dataurl.DecodeString(data)
	if err != nil {
		return nil, err
	}
	if dataURL.ContentType() != "image/png" {
		return nil, errors.New("not a png image")
	}
	m, err := png.Decode(bytes.NewReader(dataURL.Data))
	if err != nil {
		return nil, err
	}
	if m.Bounds().Dx() != width && m.Bounds().Dy() != height {
		// resize
		m = imaging.Resize(m, height, width, imaging.Lanczos)
	}

	var imgGray *image.Gray
	var ok bool
	imgGray, ok = m.(*image.Gray)
	if !ok {
		// convert to gray
		gray := imaging.Grayscale(m)
		imgGray = image.NewGray(gray.Bounds())
		for i := 0; i < len(imgGray.Pix); i++ {
			imgGray.Pix[i] = gray.Pix[i*4]
		}
	}
	return imgGray, nil
}
