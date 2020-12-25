package main

import "fmt"

func main() {
	/*
		We can handle a runtime panic with  the built-in recover function.
		recover stops the panic and returns the value that was passed
		to the call to panic. We might be tempted to recover from a panic
		like this:

			panic("PANIC")
		  str := recover()
		  fmt.Println(str)

		But the call to recover will never happen in this case,
		because the call to panic immediately stops execution of the function.
		Instead, we have to pair it with defer */
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
