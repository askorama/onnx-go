package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/png"
	"strings"
)

const img8 = `iVBORw0KGgoAAAANSUhEUgAAABwAAAAcCAAAAABXZoBIAAABUklEQVQoz32SyytFURjFf3uf7ZEZIwzupUQiJWKAkjDxP5iZKUUZGhiTMjI1MaGUoqTIxDOZeZcMuI/BTYnccs8ywH2cc7JGu37f2nt9qw3/yBSdrTHg+1Fj3u+cDTuNzRFvrjPe0Y0Nei2MnkiSNIYXenlRSu6sLGxpMwANblf+dCXQrvOSkOAxqEQPeOV06aIAf7PluDqrMLkcFpUkAZGgsybrjKOLu+JlfnSibYD+jCZwwU3iKe33VUxJawQCgUfLo5SU9sIMHNVzT5KSI+EnwcGNUpfScLAhwLGsw2qWdBC2Guy9uqHqRb15q83DWFP6mvKPA2L5SAXYwN27/aKWz3xJBRjngTK/bej1FD8AoZ5nYJz1tFMIxkiSnZ9ltbj6v6F99bIhTUasabjV9rEyA1EFecxIeuqgrPT3/Km1MXv25r6Ikv25gGintfJLg34DFzx7cr/0CSgAAAAASUVORK5CYII=`

var (
	im image.Image
	pi *image.Paletted
)

func main() {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(img8))
	im, _, _ = image.Decode(reader)
	console := flag.Bool("c", false, "console output")
	flag.Parse()
	if *console {
		outputConsole()
	} else {
		outputValues()
	}
}

func outputConsole() {
	const width = 28
	pi = image.NewPaletted(im.Bounds(), []color.Color{
		color.Gray{Y: 255},
		color.Gray{Y: 160},
		color.Gray{Y: 70},
		color.Gray{Y: 35},
		color.Gray{Y: 0},
	})

	draw.FloydSteinberg.Draw(pi, im.Bounds(), im, image.ZP)
	shade := []string{" ", "░", "▒", "▓", "█"}
	for i, p := range pi.Pix {
		fmt.Print(shade[p])
		if (i+1)%width == 0 {
			fmt.Print("\n")
		}
	}
}

func outputValues() {
	for x := 0; x < im.Bounds().Dx(); x++ {
		for y := 0; y < im.Bounds().Dy(); y++ {
			fmt.Printf("%3d ", im.At(y, x).(color.Gray).Y)
		}
		fmt.Printf("\n")
	}
}
