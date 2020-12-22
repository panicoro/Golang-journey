package main

import "fmt"

// This variable is global, we can use it
// everywhere inside our code
var x string = "Hello, World (I'm global)"

func main() {

	var x string = "Hello, World (but not the same)"
	fmt.Println(x)
	f()
}

func f() {
	fmt.Println(x)
}
