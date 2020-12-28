package main

import "fmt"

// Circle type represents the center of the circle in
// cartesian coordinates and its radius.
type Circle struct {
	x, y, r float64
}

func main() {

	// This two options create new structs with default zeros values
	var C Circle
	// This second returns a pointer to the struct (*Circle)
	otherC := new(Circle)

	fmt.Println(C.x, C.y, C.r)
	fmt.Println(otherC.x, otherC.y, otherC.r)

	initializedC := Circle{x: 0, y: 0, r: 5}
	// Another way is leave off the fields name if we know the order they were defined
	// initializedC := Circle{0, 0, 5}
	fmt.Println(initializedC.x, initializedC.y, initializedC.r)
	// If you want a pointer to the struct, use &
	// c := &Circle{0, 0, 5}
}
