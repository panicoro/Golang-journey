package main

import "fmt"

func graetest(args ...int) int {
	var max int
	for i, v := range args {
		if i == 0 || v > max {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(graetest(3, 5, 19, 10, -3))
}
