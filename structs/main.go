package main

import "fmt"

// we create structs like this
type dog struct {
	name  string
	age   int
	breed string
	bark  string
}

// we can then add methods to a struct with the receiver syntax
func (d dog) speak() string {
	return fmt.Sprintf("%s goes %s", d.name, d.bark)
}

func main() {
	// create a dog
	charlie := dog{
		name:  "Charlie",
		age:   3,
		breed: "French Bulldog",
		bark:  "Woof",
	}
	// we can do this positionally as well
	jake := dog{"Jake", 5, "British Bulldog", "Arf"}

	// call methods like so
	fmt.Println(charlie.speak())
	fmt.Println(jake.speak())

}
