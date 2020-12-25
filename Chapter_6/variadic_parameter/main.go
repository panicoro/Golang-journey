package main

import "fmt"

// By using an ellipsis (...) before the type name of the last parameter,
// you can indicate that it takes zero or more of those parameters
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	// We can also pass a slice of ints by following
	// the slice with an ellipsis:
	xs := []int{1, 2, 3, 4, 5}
	fmt.Println(add(xs...))
	fmt.Println(add(1, 2, 3, 4, 5))
}
