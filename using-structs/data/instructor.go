package data

import (
	"fmt"
)

type Instructor struct {
	// properties need to be capitalized to be public 🤢
	Id        int
	FirstName string
	LastName  string
	Score     uint
}

func (i Instructor) Print() string {
	return fmt.Sprintf("%s, %s, (%d)", i.FirstName, i.LastName, i.Score)
}
