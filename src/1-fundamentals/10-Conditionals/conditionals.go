package main

import "fmt"

func main() {

	number := 10

	if number > 15 {
		fmt.Println("Greater than 15")
	} else if number >= 10 {
		fmt.Println("Between 10 and 15")
	} else {
		fmt.Println("Lower than 10")
	}

	// if inits: defining a if-scoped variable that will be unaccessable
	// out of the if-else

	if otherNum := 16; otherNum > 15 {
		fmt.Println("Number defined is greater than 15")
	} else {
		fmt.Println("Number defined is lower or equal to 15")
	}
}
