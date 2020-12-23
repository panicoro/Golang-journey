package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	/* This is the same */
	/*
		for i:= 1; i <= 10; i++ {
			do something...
			}
	*/
}
