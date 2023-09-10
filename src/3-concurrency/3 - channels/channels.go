package main

import (
	"fmt"
	"time"
)

func main() {
	// Channels are the most common way to sync go routines
	// They can be used to receive and send signals/data
	var channel = make(chan string)

	go write("Hello world!", channel)

	fmt.Println("Waiting for channel to be received")

	// This syntax is equivalent to the one commented below it
	// When iterating through a channel, an internal validation
	// occurs to check if it's open or closed
	for message := range channel {
		fmt.Println(message)
	}

	/* for {
		// The program waits for the channel to receive data
		text, channelIsOpen := <-channel

		// This verification is needed to avoid DEADLOCKS
		// They happen when a channel is open but will never receive
		// any signal or data, is that happens, the program crashes
		// Deadlocks are NOT identified in compilation, only in runtime
		if !channelIsOpen {
			break
		}

		fmt.Println(text)
	} */

}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel) // This is needed to avoid deadlocks
}
