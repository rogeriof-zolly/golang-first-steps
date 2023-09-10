package main

import "fmt"

func main() {

	// The generator pattern consists on escapsulating the
	// goroutines declarations inside functions that are
	// not the main function
	channel := write("Hello world!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}

}

// These functions return a communication tunnel in form
// of a Channel of some type
func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Received string: %s", text)
		}
	}()

	return channel
}
