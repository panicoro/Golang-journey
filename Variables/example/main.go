package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := 2 * input

	fmt.Println("The double of your number:", output)
}
