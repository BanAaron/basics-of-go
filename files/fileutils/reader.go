package fileutils

import (
	"os"
)

func ReadTextFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	return string(content), err
}
