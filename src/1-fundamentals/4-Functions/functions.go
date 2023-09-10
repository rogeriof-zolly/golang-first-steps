package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

// Functions can return multiple values
// Important for error handling
func sumSub(a, b int) (int, int) {
	return a + b, a - b
}

func main() {

	// Functions can be assigned to variables
	var varFunc = func(a, b int) int {
		return a + b
	}

	varFuncRes := varFunc(10, 10)

	sumResult := sum(5, 5)

	// Use an _ to ignore one of the returns
	sum, sub := sumSub(15, 20)

	fmt.Println("sumResult", sumResult)
	fmt.Println("varFuncRes", varFuncRes)
	fmt.Println("sum", sum)
	fmt.Println("sub", sub)

}
