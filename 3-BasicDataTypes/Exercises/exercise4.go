//Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         //number of grid cells
	xyrange = 30.0        //axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 //angle of x, y axes (=30º)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin(30º), cos(30º)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	width, werr := strconv.Atoi(r.URL.Query().Get("width"))
	if werr != nil {
		width = 640
	}
	height, herr := strconv.Atoi(r.URL.Query().Get("height"))
	if herr != nil {
		height = 480
	}
	computeSurface(w, width, height)
}

func computeSurface(out io.Writer, width int, height int) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, aerr := corner(i+1, j, width, height)
			bx, by, bz, berr := corner(i, j, width, height)
			cx, cy, cz, cerr := corner(i, j+1, width, height)
			dx, dy, dz, derr := corner(i+1, j+1, width, height)
			if aerr != nil || berr != nil || cerr != nil || derr != nil {
				continue
			}
			height := computeColor(az, bz, cz, dz)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill=\"rgb(%g,0,0)\"/>\n", ax, ay, bx, by, cx, cy, dx, dy, height)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j, width, height int) (float64, float64, float64, error) {
	//Find point (x, y) at corner of cell (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//Compute surface height z.
	z := f(x, y)

	//Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy).
	xyscale := float64(width / 2 / xyrange)  //pixels per x or y unit
	zscale := float64(height) * float64(0.4) //pixels per z unit

	sx := float64(width/2) + (x-y)*cos30*xyscale
	sy := float64(height/2) + (x+y)*sin30*xyscale - z*zscale

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

func computeColor(a, b, c, d float64) float64 {
	return ((a + b + c + d) / 4) * 255
}
