package main

import "fmt"

// we use panic when the program gets into a state which should not be possible
// this is different to error handling

func checkAge(age uint) {
	if age > 200 {
		panic("Too old to be real")
	}
}

func main() {
	// we can use defer to execute this line at the end of the function all
	defer fmt.Println("World")
	// defers are executed in reverse order
	defer fmt.Println("Hello")

	checkAge(27)
	fmt.Println("All is Good")
}
