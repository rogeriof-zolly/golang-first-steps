package main

import "fmt"

func main() {
	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
	}

	for j := 10; j < 20; j++ {
		fmt.Println(j)
	}

	names := [3]string{"John", "Roger", "Edward"}

	fmt.Println("------ARRAYS------")
	for idx, name := range names {
		fmt.Println(idx, ":", name)
	}

	// When looping through strings, each char is a rune
	// So, it is needed to cast each one as a string
	fmt.Println("-------STRING RUNES-------")
	for idx, char := range "WORD" {
		fmt.Println(idx, ":", char)
	}
	fmt.Println("-------STRING CHARS-------")
	for idx, char := range "WORD" {
		fmt.Println(idx, ":", string(char))
	}

	fmt.Println("-------MAPS-------")

	user := map[string]string{
		"name":    "Joe",
		"surname": "Doe",
	}

	for key, value := range user {
		fmt.Println(key, ":", value)
	}
}
