package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/achal1304/One2N_GoBootcamp/grep/contract"
)

const WORKERS = 4

func ProcessDirectory(dir string, req contract.GrepRequest, resp contract.GrepResponse) {
	filePathsCh := make(chan string)
	var wg sync.WaitGroup
	var readwg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(1)
	go ReadDirectory(&wg, dir, resp, filePathsCh)
	for i := 0; i < WORKERS; i++ {
		readwg.Add(1)
		j := i
		go ReadFilesInParallel(j, &readwg, &mu, filePathsCh, req, resp)
	}
	wg.Wait()
	close(filePathsCh)
	readwg.Wait()
}

func ReadDirectory(wg *sync.WaitGroup,
	dir string,
	resp contract.GrepResponse,
	filePathsCh chan string) {
	defer wg.Done()
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", dir, err)
		return
	}

	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			wg.Add(1)
			go ReadDirectory(wg, path, resp, filePathsCh)
		} else if filepath.Ext(entry.Name()) == ".txt" {
			filePathsCh <- path
		}
	}
}

func ReadFilesInParallel(i int, readwg *sync.WaitGroup, mu *sync.Mutex,
	filePathsCh chan string,
	req contract.GrepRequest,
	resp contract.GrepResponse) {
	defer readwg.Done()
	localResultsMap := make(map[string][][]byte)
	for path := range filePathsCh {
		file, _, err := ReadFile(path)
		if err != nil {
			PrintStdOut(os.Stderr, err.Error())
			continue
		}
		searchResp, err := SearchForText(req, file)
		if err != nil {
			PrintStdOut(os.Stderr, err.Error())
			file.Close()
			continue
		}
		file.Close()
		mu.Lock()
		localResultsMap[path] = searchResp.SearchedText[req.FileName]
		mu.Unlock()
	}

	mu.Lock()
	for path, grepMatches := range localResultsMap {
		if grepMatches != nil && len(grepMatches) > 0 {
			resp.SearchedText[path] = grepMatches
		}
	}
	mu.Unlock()
}
