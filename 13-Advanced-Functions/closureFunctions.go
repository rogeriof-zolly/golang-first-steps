package main

import "fmt"

func closure() func() {
	str_text := "Inside closure function"

	closure_func := func() {
		fmt.Println(str_text)
	}

	return closure_func
}

func main() {
	// Even if this variable has the same name as the str_text
	// from closure function. This is scoped to the main function
	// and the other is scoped to the closure function
	str_text := "Inside main function"
	fmt.Println(str_text)

	new_function := closure()
	new_function()
}
