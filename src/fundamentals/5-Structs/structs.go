package main

import "fmt"

type user struct {
	name    string
	age     uint
	address address
}

type address struct {
	street string
	number int
}

func main() {

	// Struct type variables can have values defined in order of definition
	// This requires all values to be defined
	var user1 = user{"Roger", 21, address{"street example", 30}}

	// Or the values can be passed directly
	// This allows to define only some values
	user2 := user{name: "Roger2", address: address{street: "Street2", number: 122}}

	// Or the variable can be created
	var user3 user

	// So the values can be passed directly this way
	user3.name = "Roger3"
	user3.age = 21
	user3.address = address{
		street: "Street3",
		number: 20,
	}

	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)

}
