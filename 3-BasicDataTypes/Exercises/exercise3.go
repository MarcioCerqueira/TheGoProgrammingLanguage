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
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, bz, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, cz, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, dz, err := corner(i+1, j+1)
			if err != nil {
				continue
			}
			height := ((az + bz + cz + dz) / 4) * 255
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill=\"rgb(%g,0,0)\"/>\n", ax, ay, bx, by, cx, cy, dx, dy, height)
		}
	}
	fmt.Printf("</svg>")
}

func corner(i, j int) (float64, float64, float64, error) {
	//Find point (x, y) at corner of cell (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//Compute surface height z.
	z := f(x, y)

	//Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	if math.IsInf(z, 0) {
		return sx, sy, z, errors.New("Inf value was found")
	} else if math.IsNaN(z) {
		return sx, sy, z, errors.New("NaN value was found")
	} else {
		return sx, sy, z, nil
	}
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
