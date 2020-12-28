package main

import (
	"fmt"
	"math"
)

// Shape interface allows to group Circle and Rectangle structs.
type Shape interface {
	area() float64
}

// Circle type represents the center of the circle in
// cartesian coordinates and its radius.
type Circle struct {
	x, y, r float64
}

// Rectangle type represents the of the circle in
// cartesian coordinates and its radius.
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, x2, y1, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, c := range shapes {
		area += c.area()
	}
	return area
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.area())
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
	fmt.Println(totalArea(&c, &r))
}
