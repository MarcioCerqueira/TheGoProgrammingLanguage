//Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	width, height = 600, 320            //canvas size in pixels
	cells         = 100                 //number of grid cells
	xyrange       = 30.0                //axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange //pixels per x or y unit
	zscale        = height * 0.4        //pixels per z unit
	angle         = math.Pi / 6         //angle of x, y axes (=30º)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin(30º), cos(30º)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := -100; i < cells; i++ {
		for j := -100; j < cells; j++ {
			ax, ay, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Printf("</svg>")
}

func corner(i, j int) (float64, float64, error) {
	//Find point (x, y) at corner of cell (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//Compute surface height z.
	z := f(x, y)

	//Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	if math.IsInf(z, 0) {
		return sx, sy, errors.New("Inf value was found")
	} else if math.IsNaN(z) {
		return sx, sy, errors.New("NaN value was found")
	} else {
		return sx, sy, nil
	}
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
