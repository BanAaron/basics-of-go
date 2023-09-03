package main

import (
	"fmt"
	"time"
)

func foo(text string) {
	for i := 0; i < 10; i++ {
		fmt.Println(text)
		time.Sleep(time.Second / 2)
	}
}

func main() {
	go foo("Aaron")
	go foo("Chris")
	foo("Drew")
}
