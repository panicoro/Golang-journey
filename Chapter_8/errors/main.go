package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("this is a error")
	if err != nil {
		fmt.Print(err)
	}
}
