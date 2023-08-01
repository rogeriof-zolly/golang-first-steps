package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("printing in main package")
	auxiliar.Write()
	err := checkmail.ValidateFormat("123")
	fmt.Println(err)
}
