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
			ProcessDirectory(req.FileName, req, searchResponse)
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
	response := contract.GrepResponse{
		SearchedText: make(map[string][][]byte),
		Flags:        req.Flags,
	}
	lowerCaseSearchText := bytes.ToLower(req.SearchString)

	var isFound bool
	isNSearchOperation := req.Flags.BeforeSearch > 0 || req.Flags.AfterSearch > 0 ||
		req.Flags.BetweenSearch > 0
	// Update BeforeSearch and AfterSearch if BetweenFlag is provided
	if req.Flags.BetweenSearch > 0 {
		req.Flags.BeforeSearch = req.Flags.BetweenSearch
		req.Flags.AfterSearch = req.Flags.BetweenSearch
	}

	beforeBufLength := req.Flags.BeforeSearch
	afterBufLength := req.Flags.AfterSearch

	// +1 because it will include the matched element as well
	beforeBuffer := make([][]byte, beforeBufLength+1)
	afterBuffer := make([][]byte, 0, afterBufLength)

	var nextMatchCheck bool
	var lineCopy []byte
	afterCounter := 0 // Tracks lines after the match block

	for scanner.Scan() {
		line := scanner.Bytes()

		// Update beforeBuffer as a sliding window
		beforeBuffer = append(beforeBuffer[1:], append(line, '\n'))

		// Check if the current line contains the search term
		if !req.Flags.CaseInsensitive {
			if bytes.Contains(line, req.SearchString) {
				isFound = true
			}
		} else {
			lowerCaseLine := bytes.ToLower(line)
			if bytes.Contains(lowerCaseLine, lowerCaseSearchText) {
				isFound = true
			}
		}

		// If the current line matches
		if isNSearchOperation {
			if isFound {
				if !nextMatchCheck {
					// Add the beforeBuffer to the lineCopy (start of a match block)
					for _, ele := range beforeBuffer {
						lineCopy = append(lineCopy, ele...)
					}
					nextMatchCheck = true
				} else {
					// Append only the current line to the lineCopy
					lineCopy = append(lineCopy, beforeBuffer[beforeBufLength]...)
				}
				// Reset the afterBuffer and counter
				afterBuffer = afterBuffer[:0]
				afterCounter = afterBufLength
			} else if nextMatchCheck && afterCounter > 0 {
				// Collect lines in the afterBuffer after a match
				afterBuffer = append(afterBuffer, append(line, '\n'))
				lineCopy = append(lineCopy, line...)
				lineCopy = append(lineCopy, '\n')
				afterCounter--
			}
		} else {
			lineCopy = append([]byte{}, beforeBuffer[0]...)
		}

		// If no match and afterBuffer is empty, end the current match block
		if (!isFound && nextMatchCheck && afterCounter == 0) || (!isNSearchOperation && isFound) {
			utils.UpdateResponseMap(response.SearchedText, req.FileName, lineCopy)
			lineCopy = []byte{} // Reset lineCopy
			nextMatchCheck = false
			isFound = false
		}

		isFound = false // Reset match state for the next line
	}

	if err := scanner.Err(); err != nil {
		return response, fmt.Errorf("grep: %s: %v\n", req.FileName, err)
	}

	return response, nil
}
