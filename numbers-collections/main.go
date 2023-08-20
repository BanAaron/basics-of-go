package main

import (
	"fmt"
)

func main() {
	// types aren't converted automatically
	age := 27
	pi := 3.14
	// this is not valid
	// result := age * pi

	// both variables must be of the same type
	var result = float64(age) * pi
	fmt.Println(result)

	// arrays are of fixed length
	numberArray := [3]int{1, 2, 3}

	// slices can have any length
	numberSlice := []int{1, 2}

	// maps
	numberMap := map[int]string{
		1: "hello",
		2: "world",
	}

	fmt.Println(numberArray, numberSlice, numberMap)

	// we find the length of collections using len()
	length := len(numberSlice)
	// there is also capacity
	capacity := cap(numberSlice)

	fmt.Println(length, capacity)
}
