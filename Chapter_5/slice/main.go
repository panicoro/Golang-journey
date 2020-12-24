package main

import "fmt"

func main() {

	/* A slice is a segment of an array
	They are indexable and have a length, but
	anlike arrays this lenght is allowed to change.
	*/

	// Is this case x has been created with length zero.
	var x []float64
	fmt.Println(len(x))

	// This creates a slice that is associated with an underlying
	// float64 array of length 5.
	y := make([]float64, 5)
	fmt.Println(len(y))

	// 10 represents the capacity of the underlying array that the
	// slice points to:
	z := make([]float64, 5, 10)
	fmt.Println(z)

	// [low:high] -> low is the index of where to start the slice and high
	// is the index of where to end it (but not including the index itself)
	arr := [5]float64{1, 2, 3, 4, 5}
	k := arr[0:5]
	fmt.Println(k)

}
