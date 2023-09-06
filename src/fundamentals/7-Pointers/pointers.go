package main

import "fmt"

func main() {

	var var1 int = 100
	var var2 int = var1

	fmt.Println("======ASSINGING VARIABLE TO VARIABLE======")
	fmt.Println(var1, var2)

	fmt.Println("------INCREMENTING ASSIGNED VARIABLE------")
	var1++

	fmt.Println(var1, var2)

	// Pointers are memory references

	var var3 int
	var var4 *int

	var3 = 100
	var4 = &var3 // use & to point to memory registry

	fmt.Println("======ASSINGING VARIABLE TO VARIABLE MEMORY POINT======")
	fmt.Println(var3, *var4)

	fmt.Println("------INCREMENTING ASSIGNED VARIABLE TO POINTER------")
	var3++

	fmt.Println(var3, *var4) // use * to derefer (get value at that memory point)

}
