package main

import "fmt"

func main() {
	var x string = "hello"
	var y string = "world"
	fmt.Println("They are the same:", x == y)
	y = "hello"
	fmt.Println("They are the same:", x == y)

}
