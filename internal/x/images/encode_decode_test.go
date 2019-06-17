package images

import (
	"image"
	"image/color"
	"math"
	"testing"

	"gorgonia.org/tensor"
)

type circle struct {
	X, Y, R float64
}

func (c *circle) brightness(x, y float64) uint8 {
	var dx, dy float64 = c.X - x, c.Y - y
	d := math.Sqrt(dx*dx+dy*dy) / c.R
	if d > 1 {
		return 0
	}
	return 255
}

func TestEncodeDecode_RGBA(t *testing.T) {
	var w, h int = 280, 240
	var hw, hh float64 = float64(w / 2), float64(h / 2)
	r := 40.0
	θ := 2 * math.Pi / 3
	cr := &circle{hw - r*math.Sin(0), hh - r*math.Cos(0), 60}
	cg := &circle{hw - r*math.Sin(θ), hh - r*math.Cos(θ), 60}
	cb := &circle{hw - r*math.Sin(-θ), hh - r*math.Cos(-θ), 60}

	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := color.RGBA{
				cr.brightness(float64(x), float64(y)),
				cg.brightness(float64(x), float64(y)),
				cb.brightness(float64(x), float64(y)),
				255,
			}
			img.Set(x, y, c)
		}
	}

	dense := tensor.New(tensor.Of(tensor.Float32), tensor.WithShape(1, 3, img.Bounds().Max.Y, img.Bounds().Max.X))
	err := ImageToBCHW(img, dense)
	if err != nil {
		t.Fatal(err)
	}
	output, err := TensorToImg(dense)
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, img, output)

}

func assertEqual(t *testing.T, src, dst image.Image) {
	if src.Bounds() != dst.Bounds() {
		t.Fatalf("image bounds not equal: %+v, %+v", src.Bounds(), dst.Bounds())
	}
	for i := src.Bounds().Min.X; i < src.Bounds().Max.X; i++ {
		for j := src.Bounds().Min.Y; j < src.Bounds().Max.Y; j++ {
			a, b, c, d := src.At(i, j).RGBA()
			e, f, g, h := dst.At(i, j).RGBA()

			if a != e || b != f || c != g || d != h {
				t.Fatalf("image not equal: %+v, %+v", src.At(i, j), dst.At(i, j))
			}
		}
	}
}
