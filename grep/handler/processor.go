package handler

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/achal1304/One2N_GoBootcamp/grep/contract"
	"github.com/achal1304/One2N_GoBootcamp/grep/utils"
)

func ProcessGrepRequest(req contract.GrepRequest, reader io.Reader) (contract.GrepResponse, error) {
	searchResponse := contract.GrepResponse{
		SearchedText: make(map[string][][]byte),
		Flags:        req.Flags,
	}
	var err error
	var isDir bool
	if req.FileName != "" {
		var file *os.File
		file, isDir, err = ReadFile(req.FileName)
		if err != nil && !isDir {
			return contract.GrepResponse{}, err
		}
		reader = file
		defer file.Close()
	} else {
		if reader == nil {
			return contract.GrepResponse{}, fmt.Errorf("reader is nil, expected os.StdIn")
		}
	}

	if !isDir {
		searchResponse, err = SearchForText(req, reader)
		if err != nil {
			return contract.GrepResponse{}, err
		}
	} else {
		if !req.Flags.FolderCheck {
			return contract.GrepResponse{}, fmt.Errorf("grep: %s: Is a directory", req.FileName)
		} else {
			ReadDirectory(req.FileName, req, searchResponse)
		}
	}

	return searchResponse, nil
}

func ReadFile(fileName string) (*os.File, bool, error) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return nil, false, fmt.Errorf("grep: %s: No such file or directory", fileName)
	}
	if fileInfo.IsDir() {
		return nil, true, fmt.Errorf("grep: %s: Is a directory", fileName)
	}

	file, err := os.Open(fileName)
	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			return nil, false, fmt.Errorf("grep: %s: Permission denied", fileName)
		} else {
			return nil, false, err
		}
	}
	return file, false, nil
}

func SearchForText(req contract.GrepRequest, reader io.Reader) (contract.GrepResponse, error) {
	scanner := bufio.NewScanner(reader)
	response := contract.GrepResponse{SearchedText: make(map[string][][]byte)}
	lowerCaseSearchText := bytes.ToLower(req.SearchString)

	for scanner.Scan() {
		line := scanner.Bytes()
		if !req.Flags.CaseInsensitive {
			if bytes.Contains(line, req.SearchString) {
				// Copying the line as line variable points to the memory location of the buffer
				// When we append line to your map in UpdateResponseMap, the map ends up storing
				// multiple references to the same slice, which is updated in subsequent iterations.
				// which results in incorrect update in map
				lineCopy := append([]byte{}, line...)
				utils.UpdateResponseMap(response.SearchedText, req.FileName, lineCopy)
			}
		} else {
			lowerCaseLine := bytes.ToLower(line)
			if bytes.Contains(lowerCaseLine, lowerCaseSearchText) {
				// Copying the line as line variable points to the memory location of the buffer
				// When we append line to your map in UpdateResponseMap, the map ends up storing
				// multiple references to the same slice, which is updated in subsequent iterations.
				// which results in incorrect update in map
				lineCopy := append([]byte{}, line...)
				utils.UpdateResponseMap(response.SearchedText, req.FileName, lineCopy)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return response, fmt.Errorf("Error reading file: %v\n", err)
	}

	return response, nil
}
