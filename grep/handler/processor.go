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

func ProcessGrepRequest(req contract.GrepRequest) (contract.GrepResponse, error) {
	var reader io.Reader
	var searchResponse contract.GrepResponse
	if req.FileName != "" {
		file, err := ReadFile(req.FileName)
		if err != nil {
			return contract.GrepResponse{}, err
		}
		reader = file
		defer file.Close()
		searchResponse, err = SearchForText(req, reader)
		if err != nil {
			return contract.GrepResponse{}, err
		}
	} else {
		return searchResponse, fmt.Errorf("filename is not specified")
	}

	return searchResponse, nil
}

func ReadFile(fileName string) (*os.File, error) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return nil, fmt.Errorf("grep: %s: No such file or directory", fileName)
	}

	if fileInfo.IsDir() {
		return nil, fmt.Errorf("grep: %s: Is a directory", fileName)
	}

	file, err := os.Open(fileName)
	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			return nil, fmt.Errorf("grep: %s: Permission denied", fileName)
		} else {
			return nil, err
		}
	}
	return file, nil
}

func SearchForText(req contract.GrepRequest, reader io.Reader) (contract.GrepResponse, error) {
	scanner := bufio.NewScanner(reader)
	response := contract.GrepResponse{SearchedText: make(map[string][][]byte)}

	for scanner.Scan() {
		line := scanner.Bytes()
		if bytes.Contains(line, req.SearchString) {
			utils.UpdateResponseMap(response.SearchedText, req.FileName, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return response, fmt.Errorf("Error reading file: %v\n", err)
	}

	return response, nil
}
