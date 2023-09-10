package main

import (
	"fmt"
	"time"
)

func main() {

	// SELECTS are like switch statements but for channels
	// They help on working with multiple channels that are
	// waiting for messages
	chan1, chan2 := make(chan string), make(chan string)

	// If the messages are received in different frequencies
	// one can delay the receiving of messages of the other
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			chan1 <- "Message sent to channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			chan2 <- "Message sent to channel 2"
		}
	}()

	// So, select keep track of messages being received in channels
	// and runs the corresponding code when a channel receives a
	// message without waiting for other channel
	for {
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
		}
	}
}
