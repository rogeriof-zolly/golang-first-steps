package main

import (
	"fmt"
)

func main() {

	tasks := make(chan int, 45)
	results := make(chan int, 45)

	// Four concurrent processes
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 45; i++ {
		result := <-results
		fmt.Println(result)
	}

}

// <-chan : receiver channel;
// chan<- : sender channel
func worker(tasks <-chan int, results chan<- int) {
	for num := range tasks {
		results <- fibonacci(num)
	}
}

func fibonacci(position int) int {

	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}
