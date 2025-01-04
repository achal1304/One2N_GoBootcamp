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
	go ReadDirectory(&wg, dir, req, resp, filePathsCh)
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
	req contract.GrepRequest,
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
			go ReadDirectory(wg, path, req, resp, filePathsCh)
		} else {
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
		fmt.Println("reading in gorotune and path ", i, path)
		file, _, err := ReadFile(path)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", path, err)
			continue
		}
		searchResp, err := SearchForText(req, file)
		if err != nil {
			fmt.Printf("Error searching for text %s: %v\n", path, err)
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
		resp.SearchedText[path] = grepMatches
	}
	mu.Unlock()
}
