package main

import (
	"fmt"
	"time"
)

func main() {

	// CONCURRENCY !== PARALELISM
	// When creating goroutines, they compete for resource
	// When the keyword go comes before the calling of a function
	// the program will understand that it does not need to wait
	// for the function to stop executing before following the rest of the program
	go write("Hello world!")

	// It's needed to be careful on putting the word go before function calls
	// because if the next line is an EOF or return statement, the called function
	// will probably not run, because the return will be called
	write("Go programming")
	return
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
