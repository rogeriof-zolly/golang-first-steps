package main

import "fmt"

func invertNumber(num int) int {
	return num * -1
}

func invertNumberWithPointers(num *int) {
	*num = *num * -1
}

func main() {
	num1 := 20
	negative_num1 := invertNumber(num1)
	fmt.Println(num1, negative_num1)

	num2 := 40
	fmt.Println(num2)
	invertNumberWithPointers(&num2)
	fmt.Println(num2)
}
