package main

import "fmt"

// Global scoped variable
var n int

// This function is commonly used for setting up the program file
// It will always run before the main function
func init() {
	n = 10
}

func main() {
	fmt.Println(n)
}
