package main

import "fmt"

// we can declare functions like this
func add(a int, b int) int {
	return a + b
}

// functions can return multiple types:
func addAndSubtract(a int, b int) (int, int) {
	return a + b, a - b
}

func main() {
	// call a function like so
	result := add(1, 2)
	fmt.Println(result)

	// we need to capture both returned values like this
	a, b := addAndSubtract(1, 2)
	fmt.Println(a, b)

	// we can throw away a value with _
	c, _ := addAndSubtract(3, 4)
	fmt.Println(c)
}
