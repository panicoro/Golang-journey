package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2st")
}

func main() {
	// defer schedules a function call to be run after
	// the function completes...
	defer second()
	first()
}
