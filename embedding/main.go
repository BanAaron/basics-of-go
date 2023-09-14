package main

import (
	"fmt"
)

// animal is the base struct
type animal struct {
	species string
	name    string
}

// create a method foo on animal
func (a animal) foo() string {
	return "bar"
}

// dog the embeds animal
type dog struct {
	animal
	name  string
	breed string
	age   int
}

// newDog is like an init, we can pass arguments into animal with it
func newDog(species string, name string, breed string, age int) dog {
	d := dog{
		animal: animal{species, name},
		name:   name,
		breed:  breed,
		age:    age,
	}
	return d
}

// stringer is like a __repr__ method from Python
func (d dog) String() string {
	return fmt.Sprintf(
		"%s %s %s %s %d",
		d.species,
		d.animal.species, // if we have name collisions we can specify like this
		d.name,
		d.breed,
		d.age,
	)
}

func main() {
	// when we use the type we get the base type with it
	charlie := newDog(
		"Dog",
		"Charlie",
		"French Bulldog",
		3,
	)
	// here we can see what have inherited the stringer method
	fmt.Println(charlie)
	// we also get any methods on the base type
	fmt.Println(charlie.foo())
}
