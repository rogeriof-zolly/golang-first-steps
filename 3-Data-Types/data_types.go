package main

import "fmt"

func main() {

	fmt.Println("==========INTEGERS==========")

	// For non-inferred integers, you must specify the number of bits
	// There's => int8, int16, int32, int64
	var num1 int64 = 10000000000
	fmt.Println("64 bit number:", num1)

	// For inferred integers, the program will use the CPU architechture (32 or 64 bits)
	num2 := 11113000
	fmt.Println("inferred typed integer:", num2)

	// Just defining the variable as int will also use the CPU architecture
	var num3 int = 13332626
	fmt.Println("int type without bits specified:", num3)

	fmt.Println("==========FLOATS==========")

	// Just like integers, you must specify the number of bits
	// There's => float32 and float64
	// You can't specify the type with only float like with ints
	var floatNum1 float32 = 125.65
	fmt.Println("floating point number (32 bits):", floatNum1)

	// In inferred floats, the CPU architecture will also be used
	floatNum2 := 145.56
	fmt.Println("inferred typed float:", floatNum2)

	// SETTING A VALUE GREATER THAN THE SUPPORTED SPECIFIED BITS WILL CAUSE AN OVERFLOW

	fmt.Println("==========ALIASES==========")

	// There are also Alias to be used for documentation purposes like:+

	// Rune is the same as an 32 bit integer, but is commonly used to say that the number
	// should be considered as an special character, like when working with ASCII
	var aliasNum rune = 32
	fmt.Println("Rune alias:", aliasNum)

	// Byte is the same as an 8 bit integer, because that's what a byte is
	var byteNum byte = 64
	fmt.Println("Byte alias:", byteNum)

}
