package main

import "fmt"

func sumSub(n1, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	sum, sub := sumSub(30, 10)

	fmt.Println(sum, sub)
}
