package handler

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/achal1304/One2N_GoBootcamp/grep/contract"
	"github.com/stretchr/testify/assert"
)

func dirCreator(dirName string, fileNames []string, data []string) error {
	err := os.Mkdir(dirName, fs.FileMode(os.O_WRONLY|os.O_RDONLY))
	if err != nil {
		return fmt.Errorf("error while creating directory %v", err)
	}

	for i, fileName := range fileNames {
		if strings.Contains(fileName, "txt") {
			_ = os.WriteFile(filepath.Join(dirName, fileName), []byte(data[i]), 0644)
		} else {
			os.Mkdir(filepath.Join(dirName, fileName), fs.FileMode(os.O_WRONLY|os.O_RDONLY))
		}
	}
	return nil
}

func TestProcessDirectory(t *testing.T) {
	// Input parameters
	dir := "mock_directory"
	req := contract.GrepRequest{FileName: dir, SearchString: []byte("test")}
	fileNames := []string{"test1.txt", "test2.txt", "test3.txt", "test4"}
	filesData := []string{"test data 1", "test data 2", "bad data 3", "this is folder"}
	actualResponse := contract.GrepResponse{SearchedText: make(map[string][][]byte)}
	expectedResp := contract.GrepResponse{
		SearchedText: map[string][][]byte{
			filepath.Join(dir, fileNames[0]): {append([]byte(filesData[0]), '\n')},
			filepath.Join(dir, fileNames[1]): {append([]byte(filesData[1]), '\n')},
		},
	}

	err := dirCreator(dir, fileNames, filesData)
	if err != nil {
		t.Error("error creating files and directories ", err)
	}
	defer func() {
		_ = os.RemoveAll(dir)
	}()

	ProcessDirectory(dir, req, actualResponse)

	assert.Equal(t, expectedResp.SearchedText, actualResponse.SearchedText)
}

func TestReadDirecotry(t *testing.T) {
	tests := []struct {
		name         string
		dirName      string
		fileNames    []string
		filesData    []string
		expectedResp []string
		errorResonse string
		prepareDir   func(dirName string, fileNames []string, data []string) error
	}{
		{
			name:      "Happy Path Return FileNames Without Directory",
			dirName:   "dir1",
			fileNames: []string{"test1.txt", "test2.txt", "test3.txt", "test4"},
			filesData: []string{"", "", "", ""},
			expectedResp: []string{filepath.Join("dir1", "test1.txt"),
				filepath.Join("dir1", "test2.txt"),
				filepath.Join("dir1", "test3.txt")},
			prepareDir: dirCreator,
		},
		{
			name:      "Failed Path Pass File Instead Of Dir",
			dirName:   "dir1.txt",
			fileNames: []string{"test1.txt", "test2.txt", "test3.txt", "test4"},
			filesData: []string{"", "", "", ""},
			prepareDir: func(dirName string, fileNames []string, data []string) error {
				_ = os.WriteFile(dirName, []byte{}, 0644)
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// cleanup the directory and all files into it
			defer func() {
				_ = os.RemoveAll(tt.dirName)
			}()
			filesCh := make(chan string)
			var wg sync.WaitGroup
			var readwg sync.WaitGroup
			var actualFiles []string

			err := tt.prepareDir(tt.dirName, tt.fileNames, tt.filesData)
			if err != nil {
				t.Error("error preapring directories", err)
			}

			readwg.Add(1)
			go func() {
				for files := range filesCh {
					actualFiles = append(actualFiles, files)
				}
				readwg.Done()
			}()

			wg.Add(1)
			ReadDirectory(&wg, tt.dirName, contract.GrepResponse{}, filesCh)
			wg.Wait()
			close(filesCh)

			readwg.Wait()

			assert.Equal(t, tt.expectedResp, actualFiles)
		})
	}
}

func TestReadFilesInParallel(t *testing.T) {
	// Setup
	dir := "mock_directory"
	req := contract.GrepRequest{FileName: dir, SearchString: []byte("test")}
	fileNames := []string{
		"test1.txt", // correct data
		"test2.txt", // correct data
		"test3.txt", // correct data
		"mockdir2",
		filepath.Join("mockdir2", "test4.txt"), // correct data
		filepath.Join("mockdir2", "test5.txt"),
	}
	filesData := []string{
		"test data 1", // correct data
		"test data 2", // correct data
		"test data 3", // correct data
		"this is folder",
		"test data 4", // correct data
		"bad data 5"}
	actualResponse := contract.GrepResponse{SearchedText: make(map[string][][]byte)}
	expectedResp := contract.GrepResponse{
		SearchedText: map[string][][]byte{
			filepath.Join(dir, fileNames[0]): {append([]byte(filesData[0]), '\n')},
			filepath.Join(dir, fileNames[1]): {append([]byte(filesData[1]), '\n')},
			filepath.Join(dir, fileNames[2]): {append([]byte(filesData[2]), '\n')},
			filepath.Join(dir, fileNames[4]): {append([]byte(filesData[4]), '\n')},
		},
	}
	err := dirCreator(dir, fileNames, filesData)
	if err != nil {
		t.Error("error creating files and directories ", err)
	}
	defer func() {
		_ = os.RemoveAll(dir)
	}()

	filePathsCh := make(chan string)

	var readwg sync.WaitGroup
	var mu sync.Mutex

	readwg.Add(1)
	go ReadFilesInParallel(1, &readwg, &mu, filePathsCh, req, actualResponse)
	go func() {
		for _, file := range fileNames {
			filePathsCh <- filepath.Join(dir, file)
		}
		close(filePathsCh)
	}()
	readwg.Wait()

	assert.Equal(t, expectedResp.SearchedText, actualResponse.SearchedText)
}
