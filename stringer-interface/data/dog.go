package data

import (
	"fmt"
)

// Dog first create a struct
type Dog struct {
	Name  string
	Age   int
	Breed string
}

// then add a method called String()
// this essentially works like __repr__() in python
// whatever is returned here will be the string representation
func (d Dog) String() string {
	return fmt.Sprintf("-- %s -- %d -- %s\n", d.Name, d.Age, d.Breed)
}
