package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	x := 3
	y := 7
	fmt.Println("Before the swap...")
	fmt.Println("\tThe value of x is", x)
	fmt.Println("\tThe value of y is", y)
	swap(&x, &y)
	fmt.Println("After the swap...")
	fmt.Println("\tThe value of x is", x)
	fmt.Println("\tThe value of y is", y)

}
