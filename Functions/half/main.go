package main

import "fmt"

func half(x int) (int, bool) {
	x, y := x/2, x%2 == 0
	return x, y
}

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}
