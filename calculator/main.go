package main

import (
	"fmt"
)

func main() {
	var operator string
	var a, b int

	fmt.Println("CALCULATOR GO 1.0")
	fmt.Println("=================")
	fmt.Println("(add, subtract, multiply, divide)")
	fmt.Scanf("%s", &operator)
	fmt.Println("Enter first number:")
	fmt.Scanf("%d", &a)
	fmt.Println("Enter second number:")
	fmt.Scanf("%d", &b)
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
