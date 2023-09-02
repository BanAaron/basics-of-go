package main

import (
	"fmt"
)

var url = "https://www.aaronbarratt.dev"

func main() {
	// declare variables like this
	var x int
	// type declaration is optional on constants
	const pi = 3.14
	// variables default to their zero state
	// 0, false, and "" for int, bool, and string

	// can use a shorthand which will infer type like this
	var text = "Hello, World!"

	// variables in a block only exist in the blocks scope
	{
		block := "block"
		fmt.Println(block)
	}
	// therefore this doesn't work:
	// fmt.Println(block)

	fmt.Println(x, pi)
	fmt.Println(text)
	fmt.Println(url)
}
