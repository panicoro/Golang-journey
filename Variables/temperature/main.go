package main

import "fmt"

func main() {
	fmt.Print("Enter your temperature in Celsius: ")
	var fahr float64
	fmt.Scanf("%f", &fahr)

	celsius := (fahr - 32) * 5 / 9

	fmt.Println("Your Fahrenheit temperature is:", celsius)
}
