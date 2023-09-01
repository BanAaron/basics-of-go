package main

import "fmt"

// we can alias a type like this
type integer = int
type str = string
type sliceOfStrings = []string

// then use them like normal
var x integer
var y str
var z sliceOfStrings

// we can create new types
type miles float64
type kilometres float64

// constant for our conversion calculations
const milesToKilometerRatio float64 = 1.60934

// we can add methods to our type
// distance is the argument, miles is type we want to inject this method to
// these are called receiver functions
func (distance miles) toKm() kilometres {
	return kilometres(float64(distance) * milesToKilometerRatio)
}

// the same for km to miles
func (distance kilometres) toMiles() miles {
	return miles(float64(distance) / milesToKilometerRatio)
}

func main() {
	// print aliases
	fmt.Println(x, y, z)

	// created types
	distanceMiles := miles(1)
	distanceKm := kilometres(1)

	fmt.Println(distanceMiles.toKm())
	fmt.Println(distanceKm.toMiles())
}
