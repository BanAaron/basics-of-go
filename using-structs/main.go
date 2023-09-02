package main

import (
	"fmt"

	"aaronbarratt.dev/go/using-structs/data"
)

func main() {
	johnMalkovich := data.Instructor{
		Id:        1,
		FirstName: "John",
		LastName:  "Malkovich",
		Score:     100,
	}
	fmt.Println(johnMalkovich.Print())
}
