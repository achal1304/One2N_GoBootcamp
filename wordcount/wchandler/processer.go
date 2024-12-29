package wchandler

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"unicode"

	"github.com/achal1304/One2N_GoBootcamp/wordcount/contract"
	"github.com/achal1304/One2N_GoBootcamp/wordcount/wcerrors"
)

func ProcessWCCommand(wg *sync.WaitGroup,
	fileName string,
	wcFlags contract.WcFlags,
	wcValuesCh chan contract.WcValues,
	reader io.Reader) {
	defer wg.Done()
	wcCounterValues := contract.WcValues{FileName: fileName}
	if fileName != "" {
		fileInfo, err := os.Stat(fileName)
		if err != nil {
			wcCounterValues.Err = &wcerrors.WcError{Err: err, FileName: fileName}
			wcValuesCh <- wcCounterValues
			return
		}

		if fileInfo.IsDir() {
			wcCounterValues.Err = &wcerrors.WcError{Err: fmt.Errorf("%s: Is a directory", fileName), FileName: fileName}
			wcValuesCh <- wcCounterValues
			return
		}

		file, err := os.Open(fileName)
		if err != nil {
			wcCounterValues.Err = &wcerrors.WcError{Err: err, FileName: fileName}
			wcValuesCh <- wcCounterValues
			return
		}
		defer file.Close()
		reader = file
	}

	readUsingBuffers(fileName, wcValuesCh, reader)
}

func ComputeTotalCount(multipleFiles bool,
	wcValuesCh chan contract.WcValues,
	flagsOptions contract.WcFlags,
	total *contract.WcValues,
	done chan struct{},
	writer io.Writer) int {
	for wcValues := range wcValuesCh {

		total.LineCount += wcValues.LineCount
		total.CharacterCount += wcValues.CharacterCount
		total.WordCount += wcValues.WordCount
		if wcValues.Err != nil {
			if multipleFiles {
				PrintStdOut(writer, wcerrors.HandleErrors(*wcValues.Err).Err.Error())
			} else {
				PrintStdOut(os.Stderr, wcerrors.HandleErrors(*wcValues.Err).Err.Error())
				close(done)
				return 1
			}
		} else {
			PrintStdOut(writer, GenerateOutput(wcValues, flagsOptions))
		}
	}
	close(done)
	return 0
}

func readUsingBuffers(fileName string,
	wcValuesCh chan contract.WcValues,
	reader io.Reader) {
	buffer := make([]byte, 4096)
	// lineBuffer := bytes.Buffer{}

	wcCounterValues := contract.WcValues{FileName: fileName}
	leftOver := ""
	for {
		n, err := reader.Read(buffer)
		if err != nil && !errors.Is(err, io.EOF) {
			wcCounterValues.Err = &wcerrors.WcError{Err: err, FileName: fileName}
			wcValuesCh <- wcCounterValues
			return
		}

		text := leftOver + string(buffer[:n])
		leftOver = ""
		lines := bytes.Split([]byte(text), []byte("\n"))
		for i, line := range lines {
			if i == len(lines)-1 {
				leftOver = string(line)
			} else {
				// adding \n character as we are splitting based on that which excludes the character
				line = append(line, '\n')
				wcCounterValues.CharacterCount += len(line)
				wcCounterValues.LineCount++
				wcCounterValues.WordCount += countWords(string(line))
			}

		}

		if errors.Is(err, io.EOF) {
			break
		}
	}
	if len(leftOver) > 0 {
		wcCounterValues.CharacterCount += len(leftOver)
		wcCounterValues.WordCount += countWords(leftOver)
		leftOver = ""
	}

	wcValuesCh <- wcCounterValues
}

func countWords(line string) int {
	inWord := false
	wordCount := 0
	for _, r := range line {
		if unicode.IsSpace(r) {
			if inWord {
				wordCount++
			}
			inWord = false
		} else {
			inWord = true
		}
	}
	if inWord {
		wordCount++
	}
	return wordCount
}
