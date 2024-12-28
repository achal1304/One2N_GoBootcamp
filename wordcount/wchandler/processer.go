package wchandler

import (
	"bufio"
	"fmt"
	"os"
)

func ProcessWCCommand(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCounter := 0

	for scanner.Scan() {
		lineCounter++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	return lineCounter, nil
}
