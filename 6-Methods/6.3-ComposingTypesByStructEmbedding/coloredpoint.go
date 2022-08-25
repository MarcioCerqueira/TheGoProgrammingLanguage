package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

//traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	var p = ColoredPoint{Point{1, 1}, color.RGBA{255, 0, 0, 255}}
	var q = ColoredPoint{Point{5, 4}, color.RGBA{0, 0, 255, 255}}
	fmt.Println(p.Distance(q.Point))
}
