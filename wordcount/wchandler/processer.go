package wchandler

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/achal1304/One2N_GoBootcamp/wordcount/contract"
)

func ProcessWCCommand(fileName string, wcFlags contract.WcFlags) (contract.WcValues, error) {
	wcCounterValues := contract.WcValues{FileName: fileName}
	file, err := os.Open(fileName)
	if err != nil {
		return wcCounterValues, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := scanner.Text()
		wcCounterValues.LineCount++
		words := strings.Fields(lines)
		wcCounterValues.WordCount += len(words)
	}
	fileStats, err := file.Stat()
	if err != nil {
		return wcCounterValues, fmt.Errorf("error reading file stats: %v", err)
	}
	wcCounterValues.CharacterCount = int(fileStats.Size())

	if err := scanner.Err(); err != nil {
		return wcCounterValues, fmt.Errorf("error reading file: %v", err)
	}

	return wcCounterValues, nil
}
