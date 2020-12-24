package main

import "fmt"

func main() {
	silce1 := []int{1, 2, 3}
	slice2 := append(silce1, 4, 5)
	fmt.Println(silce1, slice2)
	slice3 := make([]int64, 0, 5)
	slice4 := append(slice3, 4, 5)
	fmt.Println(slice3, slice4)
}
