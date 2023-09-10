package main

import "fmt"

func main() {

	fmt.Println("======ARRAYS======")

	// Defined without inference
	var arr1 [5]string
	arr1[0] = "Data"

	// Defined with inference
	arr2 := [5]string{"data1", "data2", "data3", "data4", "data5"}

	fmt.Println(arr1)
	fmt.Println(arr2)

	// This is not dinamic size, but it sets the size to the array lenght
	arr3 := [...]string{"data1", "data2", "data3"}
	fmt.Println(arr3)

	fmt.Println("======SLICES======")

	// Slices work like arrays, but are not the same and can have dinamic lenght
	slc1 := []string{}

	fmt.Println("slc1 =", slc1)

	slc1 = append(slc1, "Added data")
	fmt.Println(slc1)

	slc1 = append(slc1, "Added more data")
	fmt.Println(slc1)

	// Slices can come from parts of arrays as pointers
	// So, if you change the sliced array, the slice will change too
	slc2 := arr2[1:3]

	fmt.Println("slc2 =", slc2)

	arr2[1] = "Altered data"

	fmt.Println("slc2 =", slc2)

	// Internal arrays

	fmt.Println("======INTERAL ARRAYS======")

	slc3 := make([]float32, 10, 11)

	fmt.Println("slc3:", slc3)
	fmt.Println("slc3 length", len(slc3))
	fmt.Println("slc3 capacity", cap(slc3))
	fmt.Println("---------------------------")

	slc3 = append(slc3, 5)
	slc3 = append(slc3, 10)

	// When an internal array capacity is reached, go creates a second one internally
	// with the same capacity so you don't run out of space
	fmt.Println("slc3:", slc3)
	fmt.Println("slc3 length", len(slc3))
	fmt.Println("slc3 capacity", cap(slc3))
	fmt.Println("---------------------------")

	slc4 := make([]float32, 5)
	fmt.Println("slc4:", slc4)
	fmt.Println("slc4 length", len(slc4))
	fmt.Println("slc4 capacity", cap(slc4))
	fmt.Println("---------------------------")

	slc4 = append(slc4, 10)
	slc4 = append(slc4, 20)
	fmt.Println("slc4:", slc4)
	fmt.Println("slc4 length", len(slc4))
	fmt.Println("slc4 capacity", cap(slc4))
	fmt.Println("---------------------------")
}
