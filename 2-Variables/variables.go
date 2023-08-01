package main

import "fmt"

func main() {

	fmt.Println("==========VARIABLES==========")

	// Declaring variables

	var variable1 string = "sample text 1"
	variable2 := "sample text 2"

	var (
		variable3 string = "sample text 3"
		variable4 string = "sample text 4"
	)

	variable5, variable6 := "sample text 5", "sample text 6"

	fmt.Println("variable1:", variable1)
	fmt.Println("variable2:", variable2)
	fmt.Println("variable3:", variable3)
	fmt.Println("variable4:", variable4)
	fmt.Println("variable5:", variable5)
	fmt.Println("variable6:", variable6)

	// Declaring constants

	fmt.Println("==========CONSTANTS==========")

	const constant1 string = "constant sample text"

	fmt.Println("constant1:", constant1)

	fmt.Println("==========SWITCHING VARIABLE VALUES==========")

	// Switching variable values
	variable5, variable6 = variable6, variable5
	fmt.Println("variable5:", variable5)
	fmt.Println("variable6:", variable6)
}
