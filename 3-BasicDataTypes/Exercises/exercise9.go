package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	width, werr := strconv.Atoi(r.URL.Query().Get("width"))
	if werr != nil {
		width = 1024
	}
	height, herr := strconv.Atoi(r.URL.Query().Get("height"))
	if herr != nil {
		height = 1024
	}
	rootFinding(w, width, height)
}

func rootFinding(out io.Writer, width int, height int) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			//Image point (px, py) represents complex value z.
			img.Set(px, py, newtonMethod(z))
		}
	}
	png.Encode(out, img)
}

func newtonMethod(z complex128) color.Color {
	const iterations = 255
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = z*z*z*z - 1
		if v == 0 {
			return color.Gray{n}
		}
		z -= (v / (4 * (z * z * z)))
	}
	return color.Black
}
