package main

import (
	"fmt"

	"aaronbarratt.dev/go/factories/data"
)

func main() {
	charlie := dog.NewDog("Charlie", "Woof")
	fmt.Println(charlie.Bark())
}
