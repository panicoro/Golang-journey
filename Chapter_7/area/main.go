package main

import (
	"fmt"
	"math"
)

// Circle type represents the center of the circle in
// cartesian coordinates and its radius.
type Circle struct {
	x, y, r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

/* If we attempted to modify one of the fields inside of
the circleArea function, it would not modify the original variable.
Because of this, we would typically write the function using a
pointer to the Circle:

func circleArea(c *Circle) float64 {
	return math.Pi * c.r*c.r
}

And change main to use & before c:
c := Circle{0, 0, 5}
fmt.Println(circleArea(&c))
*/

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(circleArea(c))
}
