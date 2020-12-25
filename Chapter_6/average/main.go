package main

import "fmt"

// An example of a function that can return multiple values
func f() (int, int) {
	return 5, 6
}

// Notice that in this function we named the return name
// This creates a new variable and we don't need to redeclare it again.
func average(someArray []float64) (total float64) {
	total = 0.0
	for _, v := range someArray {
		total += v
	}
	total /= float64(len(someArray))
	// we only use the return statement
	return
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	x, y := f()
	fmt.Println(x, y)

}
