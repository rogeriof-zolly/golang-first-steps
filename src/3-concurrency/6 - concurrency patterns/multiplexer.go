package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// A multiplexer is a output controller, where there are two or more inputs
	// and one output is delivered
	multiplexer := multiplexer(write("Hello world!"), write("Go programming!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-multiplexer)
	}

}

// The multiplexer receives two channels and returns one
func multiplexer(chan1, chan2 <-chan string) <-chan string {

	multiplexerChannel := make(chan string)

	go func() {
		// It constantly waits for messages to be received at
		// one of the channels
		for {
			// Using a select, the output will be the data received
			// in one of the channels
			select {
			case msg := <-chan1:
				multiplexerChannel <- msg
			case msg := <-chan2:
				multiplexerChannel <- msg
			}
		}
	}()

	return multiplexerChannel
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Received string: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
