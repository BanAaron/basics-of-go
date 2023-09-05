package main

import (
	"fmt"
	"time"
)

func foo(text string) {
	for i := 0; i < 10; i++ {
		fmt.Println(text)
		// we use a sleep here, so we can observe execution order in real time
		time.Sleep(time.Second / 2)
	}
}

func main() {
	// to run a go routine all we need to do is add 'go' before a function
	go foo("Aaron")
	go foo("Chris")
	// anything without 'go' runs in the main go routine
	// it is important that non-main go routines are executed before the main function exits
	foo("Drew")
}
