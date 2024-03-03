package main

import (
	"fmt"
	"github.com/banaaron/cantrips"
)

func main() {
	var operator string
	var a, b int

	cantrips.Pop()
	cantrips.KeysOrdered()

	fmt.Println("CALCULATOR GO 1.0")
	fmt.Println("=================")
	fmt.Println("(add, subtract, multiply, divide)")
	_, operatorError := fmt.Scanf("%s", &operator)
	if operatorError != nil {
		return
	}
	fmt.Println("Enter first number:")
	_, aError := fmt.Scanf("%d", &a)
	if aError != nil {
		return
	}
	fmt.Println("Enter second number:")
	_, bError := fmt.Scanf("%d", &b)
	if bError != nil {
		return
	}
	fmt.Println("Result:")

	switch operator {
	case "add":
		fmt.Println(a + b)
	case "subtract":
		fmt.Println(a - b)
	case "multiply":
		fmt.Println(a * b)
	case "divide":
		fmt.Println(a / b)
	default:
		panic("invalid operator")
	}
}
