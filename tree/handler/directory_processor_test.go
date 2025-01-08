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
		name         string
		fileNames    []string
		filesData    []string
		expectedResp *contract.TreeNode
		req          contract.TreeRequest
		errorResonse string
		prepareDir   func(dirName string, fileNames []string, data []string) error
	}{
		{
			name:      "Happy Path Folder Provider",
			fileNames: []string{"test1"},
			filesData: []string{""},
			req: contract.TreeRequest{
				FolderName: "dir1",
			},
			expectedResp: &contract.TreeNode{
				Name:         "dir1",
				IsDir:        true,
				RelativePath: "dir1",
				Path:         "dir1",
				NextDir: []*contract.TreeNode{
					{
						Name:         "test1",
						IsDir:        true,
						Path:         filepath.Join("dir1", "test1"),
						RelativePath: "dir1/test1",
						NextDir:      []*contract.TreeNode{},
					},
				},
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

			resp := &contract.TreeResponse{}

			ProcessDirectory(tt.req, resp)

			assert.Equal(t, tt.expectedResp, resp.Root)
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
			},
			expectedResp: &contract.TreeNode{
				Name:         "dir1",
				IsDir:        true,
				RelativePath: "dir1",
				Path:         "dir1",
				NextDir: []*contract.TreeNode{
					{
						Name:         "test1",
						IsDir:        true,
						Path:         filepath.Join("dir1", "test1"),
						RelativePath: "dir1/test1",
						NextDir:      []*contract.TreeNode{},
					},
					{
						Name:         "test2.txt",
						IsDir:        false,
						Path:         filepath.Join("dir1", "test2.txt"),
						RelativePath: "dir1/test2.txt",
					},
				},
			},
			prepareDir:     DirCreator,
			expfolderCount: 1,
			expfileCount:   1,
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
			},
			expectedResp: &contract.TreeNode{
				Name:         "dir1",
				IsDir:        true,
				RelativePath: "dir1",
				Path:         "dir1",
				NextDir:      []*contract.TreeNode{},
			},
			prepareDir:     DirCreator,
			expfolderCount: 0,
			expfileCount:   0,
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

			actualDCount, actualFCount := ReadDirectory(tt.req)

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
