package main

import "fmt"

func main() {
	const x string = "Hello, world"
	/* If we did this, then we would get a compile error
	x = "Some other string" */
	fmt.Println(x)
}
