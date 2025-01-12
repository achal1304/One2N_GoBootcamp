package handler

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/achal1304/One2N_GoBootcamp/tree/contract"
	"github.com/stretchr/testify/assert"
)

func TestProcessDirectory(t *testing.T) {
	tests := []struct {
		name             string
		fileNames        []string
		filesData        []string
		expectedResp     *contract.TreeNode
		req              contract.TreeRequest
		expectedTreeResp *contract.TreeResponse
		errorResonse     string
		prepareDir       func(dirName string, fileNames []string, data []string) error
	}{
		{
			name:      "Happy Path Folder Provider",
			fileNames: []string{"test1"},
			filesData: []string{""},
			req: contract.TreeRequest{
				FolderName: "dir1",
				Flags:      contract.TreeFlags{Levels: contract.MaxLevel},
			},
			expectedResp: &contract.TreeNode{
				Name:         "dir1",
				IsDir:        true,
				RelativePath: "dir1",
				Path:         "dir1",
				Permission:   "[drwxrwxrwx]",
				NextDir: []*contract.TreeNode{
					{
						Name:         "test1",
						IsDir:        true,
						Path:         filepath.Join("dir1", "test1"),
						RelativePath: "dir1/test1",
						Permission:   "[drwxrwxrwx]",
						NextDir:      []*contract.TreeNode{},
					},
				},
			},
			expectedTreeResp: &contract.TreeResponse{
				DirectoryCount: 2,
				FileCount:      0,
			},
			prepareDir: DirCreator,
		},
		{
			name:      "Directory Print No Dir In Parent Dir - Should Give Zero ",
			fileNames: []string{"test1.txt"},
			filesData: []string{""},
			req: contract.TreeRequest{
				FolderName: "dir1",
				Flags:      contract.TreeFlags{DirectoryPrint: true, Levels: contract.MaxLevel},
			},
			expectedResp: &contract.TreeNode{
				Name:         "dir1",
				IsDir:        true,
				RelativePath: "dir1",
				Path:         "dir1",
				Permission:   "[drwxrwxrwx]",
				NextDir: []*contract.TreeNode{
					{
						Name:         "test1.txt",
						IsDir:        false,
						Path:         filepath.Join("dir1", "test1.txt"),
						RelativePath: "dir1/test1.txt",
						Permission:   "[-rw-rw-rw-]",
					},
				},
			},
			expectedTreeResp: &contract.TreeResponse{
				DirectoryCount: 0,
				FileCount:      0,
			},
			prepareDir: DirCreator,
		},
		{
			name:      "Directory Print 1 Dir In Parent Dir - Should Print 2 ",
			fileNames: []string{"test1"},
			filesData: []string{""},
			req: contract.TreeRequest{
				FolderName: "dir1",
				Flags:      contract.TreeFlags{DirectoryPrint: true, Levels: contract.MaxLevel},
			},
			expectedResp: &contract.TreeNode{
				Name:         "dir1",
				IsDir:        true,
				RelativePath: "dir1",
				Path:         "dir1",
				Permission:   "[drwxrwxrwx]",
				NextDir: []*contract.TreeNode{
					{
						Name:         "test1",
						IsDir:        true,
						Path:         filepath.Join("dir1", "test1"),
						RelativePath: "dir1/test1",
						NextDir:      []*contract.TreeNode{},
						Permission:   "[drwxrwxrwx]",
					},
				},
			},
			expectedTreeResp: &contract.TreeResponse{
				DirectoryCount: 2,
				FileCount:      0,
			},
			prepareDir: DirCreator,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.prepareDir(tt.req.FolderName, tt.fileNames, tt.filesData)
			if err != nil {
				t.Error("error while creating dir ", err)
			}
			defer func() {
				_ = os.RemoveAll(tt.req.FolderName)
			}()

			ProcessDirectory(tt.req, tt.req.FolderName, tt.expectedTreeResp)

			assert.Equal(t, tt.expectedResp, tt.expectedTreeResp.Root)
		})
	}
}

func TestReadDirectory(t *testing.T) {
	tests := []struct {
		name           string
		fileNames      []string
		filesData      []string
		expectedResp   *contract.TreeNode
		req            *contract.TreeNode
		errorResonse   string
		expfolderCount int
		expfileCount   int
		reqWithFlags   contract.TreeRequest
		prepareDir     func(dirName string, fileNames []string, data []string) error
	}{
		{
			name:      "Happy Path Folder Provider",
			fileNames: []string{"test1", "test2.txt"},
			filesData: []string{"", ""},
			req: &contract.TreeNode{
				Name:         "dir1",
				Path:         "dir1",
				RelativePath: "dir1",
				IsDir:        true,
				Permission:   "[drwxrwxrwx]",
			},
			expectedResp: &contract.TreeNode{
				Name:         "dir1",
				IsDir:        true,
				RelativePath: "dir1",
				Path:         "dir1",
				Permission:   "[drwxrwxrwx]",
				NextDir: []*contract.TreeNode{
					{
						Name:         "test1",
						IsDir:        true,
						Path:         filepath.Join("dir1", "test1"),
						RelativePath: "dir1/test1",
						Permission:   "[drwxrwxrwx]",
						NextDir:      []*contract.TreeNode{},
					},
					{
						Name:         "test2.txt",
						IsDir:        false,
						Path:         filepath.Join("dir1", "test2.txt"),
						RelativePath: "dir1/test2.txt",
						Permission:   "[-rw-rw-rw-]",
					},
				},
			},
			prepareDir:     DirCreator,
			expfolderCount: 1,
			expfileCount:   1,
			reqWithFlags:   contract.TreeRequest{Flags: contract.TreeFlags{Levels: 10}},
		},
		{
			name:      "Folder Not Present",
			fileNames: []string{""},
			filesData: []string{""},
			req: &contract.TreeNode{
				Name:         "dir1",
				Path:         "dir1",
				RelativePath: "dir1",
				IsDir:        true,
				Permission:   "[drwxrwxrwx]",
			},
			expectedResp: &contract.TreeNode{
				Name:         "dir1",
				IsDir:        true,
				RelativePath: "dir1",
				Path:         "dir1",
				NextDir:      []*contract.TreeNode{},
				Permission:   "[drwxrwxrwx]",
			},
			prepareDir:     DirCreator,
			expfolderCount: 0,
			expfileCount:   0,
			reqWithFlags:   contract.TreeRequest{Flags: contract.TreeFlags{Levels: 10}},
		},
		{
			name: "Happy Path Folder With Nested Levels",
			// creating test3.txt inside test1 directory which will be at nestedlevel 2 and should not be in resp
			fileNames: []string{"test1", "test2.txt", "test1/test3.txt"},
			filesData: []string{"", "", ""},
			req: &contract.TreeNode{
				Name:         "dir1",
				Path:         "dir1",
				RelativePath: "dir1",
				IsDir:        true,
				Permission:   "[drwxrwxrwx]",
			},
			expectedResp: &contract.TreeNode{
				Name:         "dir1",
				IsDir:        true,
				RelativePath: "dir1",
				Path:         "dir1",
				Permission:   "[drwxrwxrwx]",
				NextDir: []*contract.TreeNode{
					{
						Name:         "test1",
						IsDir:        true,
						Path:         filepath.Join("dir1", "test1"),
						Permission:   "[drwxrwxrwx]",
						RelativePath: "dir1/test1",
					},
					{
						Name:         "test2.txt",
						IsDir:        false,
						Path:         filepath.Join("dir1", "test2.txt"),
						RelativePath: "dir1/test2.txt",
						Permission:   "[-rw-rw-rw-]",
					},
				},
			},
			prepareDir:     DirCreator,
			expfolderCount: 1,
			expfileCount:   1,
			reqWithFlags:   contract.TreeRequest{Flags: contract.TreeFlags{Levels: 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.prepareDir(tt.expectedResp.Name, tt.fileNames, tt.filesData)
			if err != nil {
				t.Error("error while creating dir ", err)
			}
			defer func() {
				_ = os.RemoveAll(tt.expectedResp.Name)
			}()

			actualDCount, actualFCount := ReadDirectory(tt.req, 0, tt.reqWithFlags)

			assert.Equal(t, tt.expectedResp, tt.req)
			assert.Equal(t, tt.expfolderCount, actualDCount)
			assert.Equal(t, tt.expfileCount, actualFCount)
		})
	}
}

func DirCreator(dirName string, fileNames []string, data []string) error {
	if strings.Contains(dirName, "txt") {
		_ = os.WriteFile(dirName, []byte(""), 0644)
	} else {
		err := os.Mkdir(dirName, fs.FileMode(os.O_WRONLY|os.O_RDONLY))
		if err != nil {
			return fmt.Errorf("error while creating directory %v", err)
		}
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
