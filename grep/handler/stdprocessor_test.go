package handler

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/achal1304/One2N_GoBootcamp/grep/contract"
	"github.com/achal1304/One2N_GoBootcamp/grep/helper"
	"github.com/stretchr/testify/assert"
)

func TestPrintResponseStdOut(t *testing.T) {
	tests := []struct {
		name         string
		grepResponse contract.GrepResponse
		filename     string
		writer       func(filname string, text string) (io.Writer, error)
		expected     string
		validator    func(write io.Writer) string
		cleanup      func(write io.Writer)
	}{
		{
			name: "Write to buffer or stdOut",
			grepResponse: contract.GrepResponse{
				SearchedText: map[string][][]byte{
					"testfile.txt": {[]byte("line1\nline2\nline3")},
				},
			},
			writer: func(filname, text string) (io.Writer, error) {
				var buffer bytes.Buffer
				return &buffer, nil
			},
			expected: "line1\nline2\nline3",
			validator: func(write io.Writer) string {
				buffer := write.(*bytes.Buffer)
				return buffer.String()
			},
			cleanup: func(write io.Writer) {},
		},
		{
			name: "Write To Output File",
			grepResponse: contract.GrepResponse{
				SearchedText: map[string][][]byte{
					"testfile.txt": {[]byte("line1\nline2\nline3")},
				},
			},
			filename: "outputtest.txt",
			writer: func(filename, text string) (io.Writer, error) {
				return helper.GenerateFile(filename)
			},
			expected: "line1\nline2\nline3",
			validator: func(write io.Writer) string {
				var file *os.File
				var ok bool
				if file, ok = write.(*os.File); ok {
					file.Sync() // Ensure all writes are saved
					file.Close()
				}
				fmt.Println("filename is ", file.Name())
				content, err := os.ReadFile(file.Name())
				if err != nil {
					t.Error("error reading file ", err)
					return err.Error()
				}
				return string(content)
			},
			cleanup: func(write io.Writer) {
				file := write.(*os.File)
				os.Remove(file.Name())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer, err := tt.writer(tt.filename, tt.expected)
			if err != nil {
				t.Error(err)
				return
			}
			defer tt.cleanup(writer)

			PrintResponseStdOut(writer, tt.grepResponse)

			actualOutput := tt.validator(writer)

			// as in the actual function, each line is printed on new line as we
			// are using fmt.FprintLn, hence adding \n to expected string
			assert.Equal(t, tt.expected+"\n", actualOutput)
		})
	}
}

func TestPrintStdOut(t *testing.T) {
	var buffer bytes.Buffer

	expectedOutput := "line1\nline2\nline3"
	PrintStdOut(&buffer, expectedOutput)
	actualOutput := buffer.String()
	assert.Equal(t, expectedOutput+"\n", actualOutput)
}
