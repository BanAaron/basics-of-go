package main

import "fmt"

func addYear(number *int) {
	*number++
}

func printPointer(pointer *int) {
	// we use %verb to format. %v is value, %d is digit
	fmt.Printf("Memory Address: %v Pointer Value: %d", pointer, *pointer)
}

func addYearNoPointer(number int) int {
	number++
	return number
}

func main() {
	age := 27
	addYear(&age)
	printPointer(&age)

	age = addYearNoPointer(age)
}
