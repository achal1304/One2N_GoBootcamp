package wchandler

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
	var scanner *bufio.Scanner
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

		fileStats, err := file.Stat()
		if err != nil {
			wcCounterValues.Err = &wcerrors.WcError{Err: fmt.Errorf("error reading file stats: %v", err), FileName: fileName}
			wcValuesCh <- wcCounterValues
			return
		}
		wcCounterValues.CharacterCount = int(fileStats.Size())
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(reader)
	}

	for scanner.Scan() {
		lines := scanner.Text()
		wcCounterValues.LineCount++
		words := strings.Fields(lines)
		wcCounterValues.WordCount += len(words)
		if fileName == "" {
			wcCounterValues.CharacterCount += len(lines) + 1
		}
	}

	if err := scanner.Err(); err != nil {
		wcCounterValues.Err = &wcerrors.WcError{Err: fmt.Errorf("error reading file: %v %T", err, err), FileName: fileName}
		wcValuesCh <- wcCounterValues
	}

	wcValuesCh <- wcCounterValues
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
