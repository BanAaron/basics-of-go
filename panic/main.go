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
	// defer is good for closing database connections because it will always happen at the end
	// also works with panic(). If a panic is called defers are still respected

	checkAge(27)
	fmt.Println("All is Good")
}
