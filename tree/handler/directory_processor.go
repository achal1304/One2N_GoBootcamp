package handler

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/achal1304/One2N_GoBootcamp/tree/contract"
)

func ProcessDirectory(req contract.TreeRequest, resp *contract.TreeResponse) {
	currDir := filepath.Base(req.FolderName)
	fmt.Println("foldername is ", currDir)

	root := contract.TreeNode{
		Name:         currDir,
		Path:         req.FolderName,
		RelativePath: currDir,
	}
	entries, err := os.ReadDir(root.Path)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", root.Name, err)
	}
	if len(entries) > 0 {
		resp.DirectoryCount, resp.FileCount = ReadDirectory(&root)
		// adding the current directory count
		resp.DirectoryCount++
	}
	resp.Root = &root
}

func ReadDirectory(
	root *contract.TreeNode) (int, int) {
	var dCount, fCount int
	entries, err := os.ReadDir(root.Path)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", root.Name, err)
		return dCount, fCount
	}

	nextDir := []*contract.TreeNode{}
	for _, entry := range entries {
		relativePath := root.RelativePath + "/" + entry.Name()
		path := filepath.Join(root.Path, entry.Name())
		if entry.IsDir() {
			dCount++
			nextNode := &contract.TreeNode{
				Name:         entry.Name(),
				Path:         path,
				RelativePath: relativePath,
			}
			nextDir = append(nextDir, nextNode)
			nextDCount, nextFCount := ReadDirectory(nextNode)
			dCount += nextDCount
			fCount += nextFCount
		} else {
			fCount++
			nextNode := &contract.TreeNode{
				Name:         entry.Name(),
				Path:         path,
				RelativePath: relativePath,
			}
			nextDir = append(nextDir, nextNode)
		}
	}
	root.NextDir = nextDir
	return dCount, fCount
}

// func ReadFilesInParallel(i int, readwg *sync.WaitGroup, mu *sync.Mutex,
// 	filePathsCh chan string,
// 	req contract.GrepRequest,
// 	resp contract.GrepResponse) {
// 	defer readwg.Done()
// 	localResultsMap := make(map[string][][]byte)
// 	for path := range filePathsCh {
// 		file, _, err := ReadFile(path)
// 		if err != nil {
// 			PrintStdOut(os.Stderr, err.Error())
// 			continue
// 		}
// 		searchResp, err := SearchForText(req, file)
// 		if err != nil {
// 			PrintStdOut(os.Stderr, err.Error())
// 			file.Close()
// 			continue
// 		}
// 		file.Close()
// 		mu.Lock()
// 		localResultsMap[path] = searchResp.SearchedText[req.FileName]
// 		mu.Unlock()
// 	}

// 	mu.Lock()
// 	for path, grepMatches := range localResultsMap {
// 		if grepMatches != nil && len(grepMatches) > 0 {
// 			resp.SearchedText[path] = grepMatches
// 		}
// 	}
// 	mu.Unlock()
// }
