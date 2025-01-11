package handler

import (
	"os"
	"path/filepath"

	"github.com/achal1304/One2N_GoBootcamp/tree/contract"
)

func ProcessDirectory(req contract.TreeRequest, dir string, resp *contract.TreeResponse) {
	currDir := filepath.Base(dir)

	var root contract.TreeNode
	if req.FolderName != "" {
		root = contract.TreeNode{
			Name:         currDir,
			Path:         req.FolderName,
			RelativePath: currDir,
			IsDir:        true,
		}
	} else {
		root = contract.TreeNode{
			Name:         ".",
			Path:         ".",
			RelativePath: ".",
			IsDir:        true,
		}
	}
	entries, _ := os.ReadDir(root.Path)
	if len(entries) > 0 {
		resp.DirectoryCount, resp.FileCount = ReadDirectory(&root, 0, req.Flags.Levels)
		// adding the current directory count if it contains even a single directory
		if req.Flags.DirectoryPrint {
			for _, entry := range entries {
				if entry.IsDir() {
					resp.DirectoryCount++
					break
				}
			}
		} else {
			resp.DirectoryCount++
		}
	}
	resp.Root = &root
}

func ReadDirectory(
	root *contract.TreeNode,
	currLevel int,
	maxLevel int) (int, int) {
	var dCount, fCount int
	entries, err := os.ReadDir(root.Path)
	if err != nil || currLevel >= maxLevel {
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
				IsDir:        true,
				RelativePath: relativePath,
			}
			nextDir = append(nextDir, nextNode)
			nextDCount, nextFCount := ReadDirectory(nextNode, currLevel+1, maxLevel)
			dCount += nextDCount
			fCount += nextFCount
		} else {
			fCount++
			nextNode := &contract.TreeNode{
				Name:         entry.Name(),
				Path:         path,
				IsDir:        false,
				RelativePath: relativePath,
			}
			nextDir = append(nextDir, nextNode)
		}
	}
	root.NextDir = nextDir
	return dCount, fCount
}
