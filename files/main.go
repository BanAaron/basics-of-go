package main

import (
	"aaronbarratt.dev/go/files/fileutils"
	"fmt"
)

func main() {
	content, err := fileutils.ReadTextFile("aaron")
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}
