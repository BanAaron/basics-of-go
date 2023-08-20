package main

import (
	"fmt"
)

func main() {
	// if ... else ...
	if true {
		fmt.Println(true)
	} else if false {
		fmt.Println(false)
	} else {
		panic("This state should never be possible")
	}

	// we can create variables inside an if/else
	if message := "this is a message"; true {
		fmt.Println(message)
	} else {
		// message exists in if and else scope
		fmt.Println(message)
	}

	// switch case
	day := "wednesday"
	switch day {
	case "wednesday":
		fmt.Println("Humpday")
	case "saturday":
		// we can skip over this case and execute the next one with fallthrough
		fallthrough
	case "sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	// c style for loop
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %d", i+1)
		fmt.Println()
	}

	// for in
	numbers := [5]int{1, 2, 3, 4, 5}
	for index := range numbers {
		fmt.Println(numbers[index])
	}

	// for each
	for index, value := range numbers {
		fmt.Println(index, value)
	}
	// also works with maps
	users := map[int]string{
		10: "aaron",
		15: "drew",
		20: "chris",
	}
	for key, value := range users {
		fmt.Println(key, value)
	}

	// we can make while loops with for loops
	running := true
	for running {
		fmt.Println("running")
		running = false
	}

	// while count < 10
	var count int
	for count < 10 {
		fmt.Println("counting...")
		count += 1
	}

	// infinite loop
	for {
		fmt.Println("infinite combo")
		break
	}
}
