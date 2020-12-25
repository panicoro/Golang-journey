package main

import "fmt"

func zero(xPtr *int) {
	/* Pointers reference a location in memory
	where a value is stored rather than the value itself.
	By using a pointer (*int), the zero function is able
	to modify the original variable. */
	
	// An asterisk is also used to dereference pointer variables
	// Dereferencing a pointer gives us access to the value 
	// the pointer points to
	*xPtr = 0
}

func main() {
	x := 5
	// We use the & operator to find the address of a variable
	zero(&x)
	fmt.Println(x) // x is 0
}
