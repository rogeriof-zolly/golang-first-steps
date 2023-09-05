package main

import "fmt"

// DEFER postpones the code execution to the last possible moment in the function
// Usually before returning or ending the function
// It's useful for closing database connections, where the connection will be closed
// either way, if the query was successful or not, making it unnecessary to code it
// for every possible case

func isStudentApproved(g1, g2 uint) bool {

	defer fmt.Println("Mean calculated. Displaying result:")

	fmt.Println("Calculating mean...")

	mean := (g1 + g2) / 2

	if mean >= 7 {
		return true
	}

	return false
}

func main() {
	defer fmt.Println("End of the program")

	fmt.Println(isStudentApproved(7, 8))
}
