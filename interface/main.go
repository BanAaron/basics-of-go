package main

import (
	"fmt"

	"aaronbarratt.dev/go/interface/animal"
)

func main() {
	charlie := animal.Dog{Name: "Charlie", Breed: "French Bulldog"}
	billy := animal.Cat{Name: "Billy", Breed: "Tabby"}
	// create an array of types that implement the CanSpeak interface
	animals := [2]animal.CanSpeak{charlie, billy}
	for _, a := range animals {
		fmt.Println(a.Speak())
		fmt.Println(a.ListensToCommands())
	}
}
