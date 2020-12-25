package main

import "fmt"

func fibonacci(n uint) uint {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	fmt.Println(fibonacci(10))
}
