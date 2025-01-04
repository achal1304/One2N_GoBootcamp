package handler

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/achal1304/One2N_GoBootcamp/grep/contract"
	"github.com/stretchr/testify/assert"
)

func TestProcessGrepRequest(t *testing.T) {
	tests := []struct {
		name             string
		req              contract.GrepRequest
		mockSetup        func(fileName string)
		expectedResponse contract.GrepResponse
		expectedError    string
		reader           io.Reader
	}{
		{
			name: "FileDoesNotExist",
			req: contract.GrepRequest{
				FileName:     "nonexistent.txt",
				SearchString: []byte("test"),
			},
			mockSetup: func(fileName string) {
				_ = os.Remove(fileName)
			},
			expectedResponse: contract.GrepResponse{},
			expectedError:    "grep: nonexistent.txt: No such file or directory",
		},
		{
			name: "FileIsDirectory",
			req: contract.GrepRequest{
				FileName:     "testdir",
				SearchString: []byte("test"),
			},
			mockSetup: func(fileName string) {
				_ = os.Mkdir(fileName, 0755)
			},
			expectedResponse: contract.GrepResponse{},
			expectedError:    "grep: testdir: Is a directory",
		},
		{
			name: "SuccessfulFileSearch",
			req: contract.GrepRequest{
				FileName:     "testfile.txt",
				SearchString: []byte("line 1"),
			},
			mockSetup: func(fileName string) {
				_ = os.WriteFile(fileName, []byte("line 1\nline 2\nline 3"), 0644)
			},
			expectedResponse: contract.GrepResponse{
				SearchedText: map[string][][]byte{
					"testfile.txt": {[]byte("line 1")},
				},
			},
			expectedError: "",
		},
		{
			name: "SuccessfulInputSearchStdInInput",
			req: contract.GrepRequest{
				SearchString: []byte("line 1"),
			},
			mockSetup: func(fileName string) {
				return
			},
			expectedResponse: contract.GrepResponse{
				SearchedText: map[string][][]byte{
					"": {[]byte("line 1")},
				},
			},
			expectedError: "",
			reader:        bytes.NewBufferString("line 1\nline 2\nline 3"),
		},
		{
			name: "SuccessfulInputSearchStdInInput Multiple Lines",
			req: contract.GrepRequest{
				SearchString: []byte("line 1"),
			},
			mockSetup: func(fileName string) {
				return
			},
			expectedResponse: contract.GrepResponse{
				SearchedText: map[string][][]byte{
					"": {[]byte("line 1"), []byte("line 4 and line 1")},
				},
			},
			expectedError: "",
			reader:        bytes.NewBufferString("line 1\nline 2\nline 3\nline 4 and line 1"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(tt.req.FileName)
			defer func() {
				_ = os.Remove(tt.req.FileName)
				_ = os.Remove("testdir")
			}()

			actualResponse, err := ProcessGrepRequest(tt.req, tt.reader)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				assert.Equal(t, actualResponse.SearchedText, contract.GrepResponse{}.SearchedText)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResponse.SearchedText, actualResponse.SearchedText)
			}
		})
	}
}

func TestReadFile(t *testing.T) {
	tests := []struct {
		name          string
		fileName      string
		mockSetup     func()
		expectedError string
	}{
		{
			name:     "FileDoesNotExist",
			fileName: "nonexistent.txt",
			mockSetup: func() {
			},
			expectedError: "grep: nonexistent.txt: No such file or directory",
		},
		{
			name:     "FileIsDirectory",
			fileName: "testdir",
			mockSetup: func() {
				_ = os.Mkdir("testdir", 0755)
			},
			expectedError: "grep: testdir: Is a directory",
		},
		{
			name:     "FileOpensSuccessfully",
			fileName: "testfile.txt",
			mockSetup: func() {
				_ = os.WriteFile("testfile.txt", []byte("test data"), 0644)
			},
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			defer func() {
				// Clean up after each test
				_ = os.Remove(tt.fileName)
				_ = os.Remove("testdir")
			}()

			file, _, err := ReadFile(tt.fileName)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, file)
				_ = file.Close()
			}
		})
	}
}

func TestSearchForText(t *testing.T) {
	defaultPrepareFunc := func(inputText string) (bytes.Buffer, error) {
		var buffer bytes.Buffer
		_, err := buffer.WriteString(inputText)
		if err != nil {
			return buffer, fmt.Errorf("unable to write to buffer: %w", err)
		}
		return buffer, nil
	}

	tests := []struct {
		name             string
		req              contract.GrepRequest
		expectedResponse [][]byte
		expectedError    error
		inputText        string
		prepareFile      func(inputText string) (bytes.Buffer, error)
	}{
		{
			name:             "HappyPathSingleLine Case Sensitive",
			req:              contract.GrepRequest{FileName: "test1.txt", SearchString: []byte("test1")},
			inputText:        "this is line 1 test1\nline 2 with test2",
			expectedResponse: [][]byte{[]byte("this is line 1 test1")},
			prepareFile:      defaultPrepareFunc,
		},
		{
			name:             "HappyPath MultipleLines Case Sensitive",
			req:              contract.GrepRequest{FileName: "test2.txt", SearchString: []byte("test1")},
			inputText:        "this is line 1 test1\nline 2 with test1\nline 3 with TEST1\nline 4 with test4",
			expectedResponse: [][]byte{[]byte("this is line 1 test1"), []byte("line 2 with test1")},
			prepareFile:      defaultPrepareFunc,
		},
		{
			name: "HappyPath Case Insensitive Search",
			req: contract.GrepRequest{
				FileName:     "test2.txt",
				SearchString: []byte("test1"),
				Flags: contract.GrepFlags{
					CaseInsensitive: true,
				},
			},
			inputText: "this is line 1 test1\nline 2 with test1\nline 3 with TEST1\nline 4 with test4",
			expectedResponse: [][]byte{[]byte("this is line 1 test1"), []byte("line 2 with test1"),
				[]byte("line 3 with TEST1")},
			prepareFile: defaultPrepareFunc,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bufferInput, err := tt.prepareFile(tt.inputText)

			if err != nil {
				t.Error("error from prepareFile ", err)
			}

			actualResponse, err := SearchForText(tt.req, &bufferInput)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError.Error())
			}
			if len(tt.expectedResponse) > 1 {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResponse, actualResponse.SearchedText[tt.req.FileName])
			}
		})
	}
}
