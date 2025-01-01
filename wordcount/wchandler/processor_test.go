package wchandler

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/achal1304/One2N_GoBootcamp/wordcount/contract"
	"github.com/achal1304/One2N_GoBootcamp/wordcount/wcerrors"
	"github.com/stretchr/testify/assert"
)

func TestProcessWCCommand(t *testing.T) {
	tests := []struct {
		name           string
		fileName       string
		prepare        func(fileName string, lines string) error
		inputLines     string
		expectedValues contract.WcValues
		setFlags       contract.WcFlags
		expectedErr    error
	}{
		{
			name:       "HappyPathLineCount",
			fileName:   "happyline.txt",
			inputLines: "line1\nline2",
			prepare: func(fileName string, lines string) error {
				return os.WriteFile(fileName, []byte(lines), 0644)
			},
			expectedValues: contract.WcValues{LineCount: 1, WordCount: 2, CharacterCount: 11},
			setFlags:       contract.WcFlags{LineCount: true},
			expectedErr:    nil,
		},
		{
			name:       "HappyPathLineCountStdIn",
			fileName:   "",
			inputLines: "line1-line1new\nline2new\n",
			prepare: func(fileName string, lines string) error {
				return nil
			},
			expectedValues: contract.WcValues{LineCount: 2, WordCount: 2, CharacterCount: 24},
			setFlags:       contract.WcFlags{LineCount: true},
			expectedErr:    nil,
		},
		{
			name:       "HappyPathLineWhiteSpacesCount",
			fileName:   "happyline.txt",
			inputLines: "            \n             \n   ",
			prepare: func(fileName string, lines string) error {
				return os.WriteFile(fileName, []byte(lines), 0644)
			},
			expectedValues: contract.WcValues{LineCount: 2, WordCount: 0, CharacterCount: 30},
			setFlags:       contract.WcFlags{LineCount: true},
			expectedErr:    nil,
		},
		{
			name:       "HappyPathWordCount",
			fileName:   "happyword.txt",
			inputLines: "line1-line1new\nline2",
			prepare: func(fileName string, lines string) error {
				return os.WriteFile(fileName, []byte(lines), 0644)
			},
			expectedValues: contract.WcValues{LineCount: 1, WordCount: 2, CharacterCount: 20},
			setFlags:       contract.WcFlags{WordCount: true},
			expectedErr:    nil,
		},
		{
			name:       "HappyPathWordWhiteSpacesCount",
			fileName:   "happyword.txt",
			inputLines: "                                              ",
			prepare: func(fileName string, lines string) error {
				return os.WriteFile(fileName, []byte(lines), 0644)
			},
			expectedValues: contract.WcValues{LineCount: 0, WordCount: 0, CharacterCount: 46},
			setFlags:       contract.WcFlags{WordCount: true},
			expectedErr:    nil,
		},
		{
			name:     "NoFileFound",
			fileName: "notfound.txt",
			prepare: func(fileName string, lines string) error {
				return nil
			},
			expectedValues: contract.WcValues{Err: &wcerrors.WcError{Err: os.ErrNotExist}},
			setFlags:       contract.WcFlags{LineCount: true},
			expectedErr:    os.ErrNotExist,
		},
		{
			name:       "ReadPermissionDenied",
			fileName:   "read.txt",
			inputLines: "Line 1\nLine 2\n",
			prepare: func(fileName string, lines string) error {
				err := os.WriteFile(fileName, []byte(lines), 0000)
				if err != nil {
					return err
				}
				return nil
			},
			expectedValues: contract.WcValues{LineCount: 2, WordCount: 4, CharacterCount: 14},
			setFlags:       contract.WcFlags{LineCount: true},
		},
		{
			name:     "DirectoryFailure",
			fileName: "cmd",
			prepare: func(fileName string, lines string) error {
				err := os.Mkdir("cmd", 0644)
				if err != nil {
					return err
				}
				return nil
			},
			expectedValues: contract.WcValues{Err: &wcerrors.WcError{Err: fmt.Errorf("cmd: Is a directory")}},
			setFlags:       contract.WcFlags{},
			expectedErr:    fmt.Errorf("cmd: Is a directory"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buffer bytes.Buffer
			if tt.fileName != "" {
				err := tt.prepare(tt.fileName, tt.inputLines)
				if err != nil {
					t.Error("unable to prepare file", err)
					return
				}
				defer os.Remove(tt.fileName)
			} else {
				_, err := buffer.WriteString(tt.inputLines)
				if err != nil {
					t.Error("unable to write to buffer file", err)
					return
				}
			}

			var wg sync.WaitGroup
			wcValuesCh := make(chan contract.WcValues, 1)

			wg.Add(1)
			go ProcessWCCommand(&wg, tt.fileName, tt.setFlags, wcValuesCh, &buffer)
			wg.Wait()
			close(wcValuesCh)

			actualValues := <-wcValuesCh

			if tt.expectedErr != nil {
				assert.Error(t, actualValues.Err.Err)
				if tt.expectedErr.Error() == actualValues.Err.Err.Error() {
					assert.Equal(t, tt.expectedErr.Error(), actualValues.Err.Err.Error())
				} else {
					assert.True(t, errors.Is(actualValues.Err.Err, tt.expectedErr))
				}
			}

			tt.expectedValues.FileName = tt.fileName
			assert.Equal(t, tt.expectedValues.LineCount, actualValues.LineCount)
			assert.Equal(t, tt.expectedValues.WordCount, actualValues.WordCount)
			assert.Equal(t, tt.expectedValues.CharacterCount, actualValues.CharacterCount)
			assert.Equal(t, tt.expectedValues.FileName, actualValues.FileName)
		})
	}
}

func TestComputeTotalCount(t *testing.T) {
	tests := []struct {
		name             string
		multipleFiles    bool
		flagsOptions     contract.WcFlags
		wcValues         []contract.WcValues
		expectedTotal    contract.WcValues
		expectedExitCode int
		expectedPrint    []string
		exitStatusCode   int
	}{
		{
			name:          "HappyPathTwoFileLineCount",
			multipleFiles: true,
			flagsOptions:  contract.WcFlags{LineCount: true},
			wcValues: []contract.WcValues{
				{LineCount: 2, FileName: "1.txt"},
				{LineCount: 5, FileName: "2.txt"},
			},
			expectedTotal: contract.WcValues{
				LineCount: 7,
			},
			expectedPrint: []string{
				"       2 1.txt\n       5 2.txt\n",
				"       5 2.txt\n       2 1.txt\n",
			},
			exitStatusCode: 0,
		},
		{
			name:          "HappyPathSingleFileLineCount",
			multipleFiles: false,
			flagsOptions:  contract.WcFlags{LineCount: true},
			wcValues: []contract.WcValues{
				{LineCount: 2, FileName: "1.txt"},
			},
			expectedTotal: contract.WcValues{LineCount: 2},
			expectedPrint: []string{
				"       2 1.txt\n",
			},
			exitStatusCode: 0,
		},
		{
			name:          "ErrorSingleFileLineCount",
			multipleFiles: false,
			flagsOptions:  contract.WcFlags{LineCount: true},
			wcValues: []contract.WcValues{
				{Err: &wcerrors.WcError{Err: os.ErrNotExist, FileName: "1.txt"}, FileName: "1.txt"},
			},
			expectedTotal: contract.WcValues{},
			expectedPrint: []string{
				"wc: 1.txt: read: No such file or directory\n",
			},
			exitStatusCode: 1,
		},
		{
			name:          "ExitStatus1MultipleFileLineCount",
			multipleFiles: true,
			flagsOptions:  contract.WcFlags{LineCount: true},
			wcValues: []contract.WcValues{
				{Err: &wcerrors.WcError{Err: os.ErrNotExist, FileName: "1.txt"}},
				{LineCount: 5, FileName: "2.txt"},
			},
			expectedTotal: contract.WcValues{
				LineCount: 5,
			},
			expectedPrint: []string{
				"wc: 1.txt: read: No such file or directory\n       5 2.txt\n",
				"       5 2.txt\nwc: 1.txt: read: No such file or directory\n",
			},
			exitStatusCode: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wcValuesCh := make(chan contract.WcValues)
			done := make(chan struct{})
			actualTotal := contract.WcValues{}
			actualExitCode := make(chan int)
			defer close(actualExitCode)
			var buffer bytes.Buffer

			go func(wcValues []contract.WcValues) {
				for _, value := range wcValues {
					wcValuesCh <- value
				}
				close(wcValuesCh)
			}(tt.wcValues)

			go func() {
				actualExitStatusCode := ComputeTotalCount(tt.multipleFiles, wcValuesCh, tt.flagsOptions, &actualTotal, done, &buffer)
				actualExitCode <- actualExitStatusCode
			}()

			<-done
			bufferOutput := buffer.String()
			if tt.exitStatusCode == 0 {
				assert.Contains(t, tt.expectedPrint, bufferOutput)
			}

			assert.Equal(t, tt.exitStatusCode, <-actualExitCode)

			assert.Equal(t, tt.expectedTotal.FileName, actualTotal.FileName)
			assert.Equal(t, tt.expectedTotal.WordCount, actualTotal.WordCount)
			assert.Equal(t, tt.expectedTotal.CharacterCount, actualTotal.CharacterCount)
			assert.Equal(t, tt.expectedTotal.LineCount, actualTotal.LineCount)
		})
	}
}

func TestReadLines(t *testing.T) {
	tests := []struct {
		name           string
		lines          [][]byte
		expectedCounts contract.WcValues
	}{
		{
			name: "Multiple lines with words",
			lines: [][]byte{
				[]byte("This is a test."),
				[]byte("Another line."),
				[]byte("Final line."),
			},
			expectedCounts: contract.WcValues{
				CharacterCount: 42,
				LineCount:      3,
				WordCount:      8,
			},
		},
		{
			name: "Empty lines",
			lines: [][]byte{
				[]byte(""),
				[]byte(""),
				[]byte(""),
			},
			expectedCounts: contract.WcValues{
				CharacterCount: 3,
				LineCount:      3,
				WordCount:      0,
			},
		},
		{
			name: "Lines with whitespace",
			lines: [][]byte{
				[]byte(" "),
				[]byte("    "),
				[]byte("  word  "),
			},
			expectedCounts: contract.WcValues{
				CharacterCount: 16,
				LineCount:      3,
				WordCount:      1,
			},
		},
		{
			name: "Single line with no newline",
			lines: [][]byte{
				[]byte("Single line with words."),
			},
			expectedCounts: contract.WcValues{
				CharacterCount: 24,
				LineCount:      1,
				WordCount:      4,
			},
		},
		{
			name:  "No lines (empty input)",
			lines: [][]byte{},
			expectedCounts: contract.WcValues{
				CharacterCount: 0,
				LineCount:      0,
				WordCount:      0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wg sync.WaitGroup
			resultCh := make(chan contract.WcValues)

			wg.Add(1)
			go func() {
				ReadLines(&wg, tt.lines, resultCh)
			}()
			result := <-resultCh
			wg.Wait()

			close(resultCh)

			assert.Equal(t, tt.expectedCounts.CharacterCount, result.CharacterCount, "CharacterCount mismatch")
			assert.Equal(t, tt.expectedCounts.LineCount, result.LineCount, "LineCount mismatch")
			assert.Equal(t, tt.expectedCounts.WordCount, result.WordCount, "WordCount mismatch")

		})
	}
}
