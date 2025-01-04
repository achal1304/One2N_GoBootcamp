package handler

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/achal1304/One2N_GoBootcamp/grep/contract"
)

func ReadDirectory(dir string, req contract.GrepRequest, resp contract.GrepResponse) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", dir, err)
		return
	}

	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			ReadDirectory(path, req, resp)
		} else {
			file, _, err := ReadFile(path)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", path, err)
				continue
			}
			defer file.Close()
			searchResp, err := SearchForText(req, file)
			if err != nil {
				fmt.Printf("Error searching for text %s: %v\n", path, err)
				continue
			}

			for _, v := range searchResp.SearchedText {
				resp.SearchedText[path] = v
			}
		}
	}
}
