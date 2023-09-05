package main

import (
	"fmt"
	"time"
)

// channels are used to communicate data between goroutines
// you can also use them to wait for a routine to end

func count(number int, channel chan string) {
	for i := 0; i < number; i++ {
		fmt.Println(i + 1)
		time.Sleep(time.Second / 2)
	}
	channel <- "Done!"
}

func main() {
	// create a channels with make
	channel := make(chan string)
	go count(5, channel)
	// this waits for the goroutine that takes this channel to complete
	<-channel

	// we can make buffer channels by providing a number
	buffer := make(chan string, 3)
	// assign up to 3 values to the buffer channel
	buffer <- "Hello"
	buffer <- "World"
	buffer <- "Noob"
	// we then consume each value in the buffer in order of insertion
	fmt.Println(<-buffer)
	fmt.Println(<-buffer)
	fmt.Println(<-buffer)
}
