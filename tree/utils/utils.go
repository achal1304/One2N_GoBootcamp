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

	if !info.IsDir() {
		return false
	}
	return true
}

// Converts file mode to a permission string like "rw-r--r--"
func GetPermissionString(mode os.FileMode) string {
	return "[" + mode.String() + "]"
}
