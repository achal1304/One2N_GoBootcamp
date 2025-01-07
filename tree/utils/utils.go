package utils

import (
	"fmt"
	"os"
)

func UpdateResponseMap(respMap map[string][][]byte, key string, value []byte) {
	if _, ok := respMap[key]; ok {
		respMap[key] = append(respMap[key], value)
	} else {
		respMap[key] = [][]byte{value}
	}
}

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

	if !info.IsDir() {
		return false
	}
	return true
}
