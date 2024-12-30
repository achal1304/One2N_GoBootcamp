package wchandler

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

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

	ReadUsingBuffers(fileName, wcValuesCh, reader)
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

func ReadUsingBuffers(fileName string,
	wcValuesCh chan contract.WcValues,
	reader io.Reader) {
	buffer := make([]byte, 1024*1024)
	// lineBuffer := bytes.Buffer{}

	wcCounterValues := contract.WcValues{FileName: fileName}
	resultCh := make(chan contract.WcValues)
	done := make(chan struct{})
	var wg sync.WaitGroup
	leftOver := []byte{}

	go func(resultCh chan contract.WcValues, done chan struct{},
		wcCounterValues *contract.WcValues) {
		for res := range resultCh {
			wcCounterValues.LineCount += res.LineCount
			wcCounterValues.CharacterCount += res.CharacterCount
			wcCounterValues.WordCount += res.WordCount
		}
		close(done)
	}(resultCh, done, &wcCounterValues)

	workerPool := make(chan struct{}, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil && !errors.Is(err, io.EOF) {
			wcCounterValues.Err = &wcerrors.WcError{Err: err, FileName: fileName}
			wcValuesCh <- wcCounterValues
			return
		}

		text := append(leftOver, buffer[:n]...)
		leftOver = nil
		lines := bytes.Split([]byte(text), []byte("\n"))
		totalLines := len(lines)
		leftOver = lines[totalLines-1]

		lines = lines[:totalLines-1]

		wg.Add(1)
		workerPool <- struct{}{} // Acquire a worker
		go func(wg *sync.WaitGroup, lines [][]byte, resultCh chan contract.WcValues) {
			defer wg.Done()
			var wcCounterValues contract.WcValues
			for _, line := range lines {
				// adding \n character as we are splitting based on that which excludes the character
				line = append(line, '\n')
				wcCounterValues.CharacterCount += len(line)
				wcCounterValues.LineCount++
				wcCounterValues.WordCount += len(bytes.Fields(line))
			}

			resultCh <- wcCounterValues
			<-workerPool // Release the worker
		}(&wg, append([][]byte(nil), lines...), resultCh)

		if errors.Is(err, io.EOF) {
			break
		}
	}

	wg.Wait()
	close(resultCh)
	<-done

	if len(leftOver) > 0 {
		wcCounterValues.CharacterCount += len(leftOver)
		wcCounterValues.WordCount += len(bytes.Fields(leftOver))
		leftOver = nil
	}

	wcValuesCh <- wcCounterValues
}

// func ReadLines(wg *sync.WaitGroup,
// 	lines [][]byte,
// 	resultCh chan contract.WcValues) {
// 	wcCounterValues := contract.WcValues{}
// 	defer wg.Done()
// 	for _, line := range lines {
// 		// adding \n character as we are splitting based on that which excludes the character
// 		line = append(line, '\n')
// 		wcCounterValues.CharacterCount += len(line)
// 		wcCounterValues.LineCount++
// 		wcCounterValues.WordCount += countWords(line)
// 	}

// 	resultCh <- wcCounterValues
// }

func countWords(line []byte) int {
	inWord := false
	wordCount := 0
	for _, b := range line {
		if isSpace(b) {
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

func isSpace(b byte) bool {
	// Covers common whitespace characters: space, tab, and newlines
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}
