package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/achal1304/One2N_GoBootcamp/tree/contract"
	"github.com/achal1304/One2N_GoBootcamp/tree/utils"
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
	info, err := os.Stat(root.Path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	entries, _ := os.ReadDir(root.Path)
	root.Permission = utils.GetPermissionString(info.Mode())
	if len(entries) > 0 {
		resp.DirectoryCount, resp.FileCount = ReadDirectory(&root, 0, req)
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
	req contract.TreeRequest) (int, int) {
	var dCount, fCount int
	entries, err := os.ReadDir(root.Path)
	if err != nil || currLevel >= req.Flags.Levels {
		return dCount, fCount
	}

	if req.Flags.RecentlyModified {
		sort.Slice(entries, func(i, j int) bool {
			infoI, err := entries[i].Info()
			if err != nil {
				PrintStdOut(os.Stderr, fmt.Sprintf("tree: error while reading %s", infoI.Name()))
			}
			infoJ, err := entries[j].Info()
			if err != nil {
				PrintStdOut(os.Stderr, fmt.Sprintf("tree: error while reading %s", infoJ.Name()))
			}
			return infoI.ModTime().Before(infoJ.ModTime())
		})
	}

	nextDir := []*contract.TreeNode{}
	for _, entry := range entries {
		relativePath := root.RelativePath + "/" + entry.Name()
		path := filepath.Join(root.Path, entry.Name())
		// Get full file info to access permissions
		fileInfo, err := entry.Info()
		if err != nil {
			continue // Skip if file info cannot be read
		}

		permission := utils.GetPermissionString(fileInfo.Mode())

		if entry.IsDir() {
			dCount++
			nextNode := &contract.TreeNode{
				Name:         entry.Name(),
				Path:         path,
				IsDir:        true,
				Permission:   permission,
				RelativePath: relativePath,
			}
			nextDir = append(nextDir, nextNode)
			nextDCount, nextFCount := ReadDirectory(nextNode, currLevel+1, req)
			dCount += nextDCount
			fCount += nextFCount
		} else {
			fCount++
			nextNode := &contract.TreeNode{
				Name:         entry.Name(),
				Path:         path,
				IsDir:        false,
				Permission:   permission,
				RelativePath: relativePath,
			}
			nextDir = append(nextDir, nextNode)
		}
	}
	root.NextDir = nextDir
	return dCount, fCount
}
