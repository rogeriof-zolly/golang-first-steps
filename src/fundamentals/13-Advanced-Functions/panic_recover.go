package main

import "fmt"

func recovery() {
	// this verifies if the program panicked
	// and treats accordingly
	if r := recover(); r != nil {
		fmt.Println("Program recovered!")
	}
}

func isStudentApproved(g1, g2 float64) bool {

	// the recovery function must be defered, because it will run
	// at any condition
	defer recovery()

	fmt.Println("Calculating mean...")

	mean := (g1 + g2) / 2

	if mean > 7 {
		return true
	} else if mean < 7 {
		return false
	}

	// This is similar, but not equal, to an error
	// This covers a possible unpredicted case that
	// needs to be treated, but not to stop the execution
	// when 'panic' runs, it looks for the 'defered' functions
	// and runs them, that's why the recovery func is called with defer
	// in the beggining
	panic("The mean is equal to 7!")
}

func main() {
	fmt.Println(isStudentApproved(7, 7))
}
