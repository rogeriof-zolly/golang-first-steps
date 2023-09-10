package main

import "fmt"

type user struct {
	name string
	age  int
}

/*
Methods are linked to a struct
They can do anything related to the struct
Create, read, delete or update any data structured that way
*/
func (u user) save() {
	fmt.Printf("Saving user with name %s, %d years old, in database\n", u.name, u.age)
}

func (u user) isAdult() bool {
	return u.age >= 18
}

// To update data from a struct, a pointer referencing the struct can be passed
// It is not needed to dereference the variable inside the method
func (u *user) ageUp() {
	u.age++
}

func main() {
	user1 := user{name: "Rogerio", age: 21}
	fmt.Println(user1)
	user1.save()
	fmt.Println(user1.isAdult())
	user1.ageUp()
	user1.save()
}
