package main

import "fmt"

func main() {

	anonFunc := func(text string) string {
		return fmt.Sprintf("Received string: %s", text)
	}("Parameter")

	fmt.Println(anonFunc)
}
