package main

import (
	"fmt"

	"github.com/banaaron/files/fileutils"
)

func main() {
	passwords, readFileError := fileutils.ReadTextFile("/data/secret_passwords.txt")
	if readFileError != nil {
		panic(readFileError)
	}

	writeFileError := fileutils.WriteTextFile("/data/stolen_passwords.txt", passwords)
	if writeFileError != nil {
		panic(writeFileError)
	}
	fmt.Println("Passwords stolen!")
}
