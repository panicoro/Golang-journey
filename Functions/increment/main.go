package main

import "fmt"

func main() {
	x := 0
	/* When you create a local function like this, it also has
	/ access to other local variables
	A function like this together with the nonlocal variables it references 
	is known as a closure.
	In this case, increment and the variable x form the closure.*/
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
