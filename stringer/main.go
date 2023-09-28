package main

import (
	"fmt"

	"github.com/banaaron/stringer-interface/data"
)

func main() {
	charlie := data.Dog{
		Name:  "Charlie",
		Age:   3,
		Breed: "French Bulldog",
	}
	fmt.Println(charlie)
}
