package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			//Image point (px, py) represents complex value z.
			img.Set(px, py, newtonMethod(z))
		}
	}
	f, _ := os.Create("./output.png")
	defer f.Close()
	// save image
	png.Encode(f, img) //NOTE: ignoring errors
}

func newtonMethod(z complex64) color.Color {
	const iterations = 255
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = z*z*z*z - 1
		if v == 0 {
			return color.Gray{n}
		}
		z -= (v / (4 * (z * z * z)))
	}
	return color.Black
}
