package fileutils

import (
	"os"
)

func ReadTextFile(fileName string) (string, error) {
	cwd, _ := os.Getwd()
	content, err := os.ReadFile(cwd + fileName)
	return string(content), err
}

func WriteTextFile(fileName string, content string) error {
	cwd, _ := os.Getwd()
	return os.WriteFile(cwd+fileName, []byte(content), 0644)
}
