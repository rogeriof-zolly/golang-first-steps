package main

import "fmt"

type person struct {
	name   string
	age    uint
	height uint
}

type student struct {
	person
	course     string
	university string
}

func main() {

	p1 := person{"Rogerio", 21, 174}

	s1 := student{p1, "Computer Engineering", "UNISC"}

	s2 := student{person{"Rogério", 21, 174}, "CE", "UNISC"}

	fmt.Println(s1)

	// There are various ways to call properties from composited structs
	fmt.Println(s2.name)
	fmt.Println(s2.person.age)
	fmt.Println(s2.university)
}
