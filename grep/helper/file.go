package helper

import (
	"fmt"
	"os"
)

func GenerateFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666)
	if err != nil {
		if os.IsExist(err) {
			return nil, fmt.Errorf("grep: %s: File already exists", fileName)
		} else {
			return nil, fmt.Errorf("grep: %s: Error while creating file", fileName)
		}
	}
	return file, err
}
