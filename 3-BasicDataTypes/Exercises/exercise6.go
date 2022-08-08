//Mandelbrot emits a PNG image of the Mandelbrot integral.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			leftIntensity := float64(getPixelValue(float64(px), float64(py)))
			topIntensity := float64(getPixelValue(float64(px), float64(py)))
			rightIntensity := float64(getPixelValue(float64(px), float64(py)))
			bottomIntensity := float64(getPixelValue(float64(px), float64(py)))
			averageIntensity := uint8((leftIntensity + topIntensity + rightIntensity + bottomIntensity) / 4)
			//Image point (px, py) represents complex value z.
			img.Set(px, py, color.Gray{averageIntensity})
		}
	}
	f, _ := os.Create("./output.png")
	defer f.Close()
	// save image
	png.Encode(f, img) //NOTE: ignoring errors
}

func getPixelValue(px, py float64) uint8 {
	x := px/width*(xmax-xmin) + xmin
	y := py/height*(ymax-ymin) + ymin
	z := complex(x, y)
	return mandelbrot(z)
}
func mandelbrot(z complex128) uint8 {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return 255 - contrast*n
		}
	}
	return 0
}
