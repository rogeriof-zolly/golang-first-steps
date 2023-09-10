package main

import "fmt"

func main() {
	// A buffered channel has a limit to number of messages
	// it keeps at the same time
	// To use a channel outside goroutines, you need to use buffered channels
	// so deadlocks don't occur
	var channel = make(chan string, 2)

	// It can receive a max number of messages until they're read
	channel <- "Hello world"
	channel <- "Go programming!"

	// This 'consumes' the first element from the channel
	msg1 := <-channel

	// So, now it only has 1 message, then, one more can be added
	channel <- "One more"

	msg2 := <-channel
	msg3 := <-channel

	// If a message be added to the channel buffer when it is full
	// the program will crash because of a deadlock

	// It's implementation follows the First In, First Out method

	fmt.Println(msg1)
	fmt.Println(msg2)
	fmt.Println(msg3)
}
