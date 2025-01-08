package utils

import (
	"fmt"
	"os"
)

func GetCurrentDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("error fetching the current directory:", err)
		return "", err
	}
	return currentDir, nil
}

func CheckDirectory(dir string) bool {
	info, err := os.Stat(dir)
	if err != nil {
		return false
	}

	fmt.Println("isdir ", info.IsDir(), info.Name())
	if !info.IsDir() {
		return false
	}
	return true
}
