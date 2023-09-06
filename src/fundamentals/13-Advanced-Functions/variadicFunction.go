package main

import "fmt"

func sum(numbers ...int) {
	fmt.Println(numbers)

	total := 0

	for _, num := range numbers {
		total += num
	}

	fmt.Println(total)
}

func write(text string, numbers ...int) {
	for _, num := range numbers {
		fmt.Println(text, num)
	}
}

func main() {
	sum(1, 2, 3, 4, 5, 6)
	write("Hello", 1, 5, 454, 76, 786, 86, 213)
}
