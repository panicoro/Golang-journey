package main

import "fmt"

func main() {
	fmt.Print("Enter your distance in feet: ")
	var feetDistance float64
	fmt.Scanf("%f", &feetDistance)

	meterDistance := feetDistance * 0.3048

	fmt.Println("Your distance in meters is:", meterDistance)
}
