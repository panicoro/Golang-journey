package main

import (
	"fmt"
	"math"
)

// Shape interface allows to group Circle and Rectangle structs.
type Shape interface {
	area() float64
	perimeter() float64
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

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, c := range shapes {
		area += c.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, c := range shapes {
		perimeter += c.perimeter()
	}
	return perimeter
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println("This is the circle area:", c.area())
	fmt.Println("This is the perimeter circle :", c.perimeter())
	fmt.Println()
	r := Rectangle{0, 0, 10, 10}
	fmt.Println("This is the rectangle area:", r.area())
	fmt.Println("This is the perimeter rectangle:", r.perimeter())
	fmt.Println()
	fmt.Println("The total area is:", totalArea(&c, &r))
	fmt.Println("The total perimeter is:", totalPerimeter(&c, &r))
}
