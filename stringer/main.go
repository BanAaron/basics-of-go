package main

import (
	"fmt"

	"aaronbarratt.dev/go/stringer-interface/data"
)

func main() {
	charlie := data.Dog{
		Name:  "Charlie",
		Age:   3,
		Breed: "French Bulldog",
	}
	fmt.Println(charlie)
}
