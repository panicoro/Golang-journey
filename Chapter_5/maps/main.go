package main

import "fmt"

func main() {
	// that maps have to be initialized before they can be used
	x := make(map[string]int)
	x["key"] = 10
	// not working like this: it has to be initialized
	// var x map[string]int
	// fmt.Println(x)
	fmt.Println(x["key"])
	
	y := make(map[int]int)
	y[1] = 10
	fmt.Println(y[1])
	fmt.Println(len(x))
	delete(x, "key")
	fmt.Println(len(x))

}
