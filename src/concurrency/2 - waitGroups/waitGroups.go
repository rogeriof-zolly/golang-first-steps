package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		write("Hello world!")
		waitGroup.Done() // This will subtract 1 from the defined number in Add above
	}()

	go func() {
		write("Go programming!")
		waitGroup.Done()
	}()

	waitGroup.Wait() // This waits the counter defined to reach 0 (only when both functions ran)
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
