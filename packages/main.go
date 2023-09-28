package main

import (
	// we import from another module like this
	"github.com/banaaron/packges/aaron"
	// import from core modules like this
	"fmt"
)

func main() {
	// we can access anything from the same package
	fmt.Println(charmander)
	fmt.Println(pikachu)

	// the same for data types
	var squirtle = poke{
		"squirtle",
		[]string{"water"},
	}
	fmt.Println(squirtle)

	// because Exported is title case it is exported from aaron.go
	aaron.Exported()
	// but notExported() does not work because it is camel case
	// aaron.notExported()
}
