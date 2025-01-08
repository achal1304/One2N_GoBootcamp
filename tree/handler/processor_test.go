package handler

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/achal1304/One2N_GoBootcamp/tree/contract"
	"github.com/stretchr/testify/assert"
)

func TestProcessTreeRequest(t *testing.T) {
	tests := []struct {
		name          string
		req           contract.TreeRequest
		expectedResp  *contract.TreeResponse
		errorExpected bool
		fileNames     []string
		filesData     []string
		errorMessage  string
		prepareDir    func(dirName string, fileNames []string, data []string) error
	}{
		{
			name: "Happy Path with FolderName Provided",
			req: contract.TreeRequest{
				FolderName: "testDir",
			},
			fileNames: []string{"subDir"},
			filesData: []string{""},
			expectedResp: &contract.TreeResponse{
				Root: &contract.TreeNode{
					Name:         "testDir",
					IsDir:        true,
					Path:         "testDir",
					RelativePath: "testDir",
					NextDir: []*contract.TreeNode{
						{
							Name:         "subDir",
							IsDir:        true,
							Path:         filepath.Join("testDir", "subDir"),
							RelativePath: "testDir/subDir",
							NextDir:      []*contract.TreeNode{},
						},
					},
				},
			},
			errorExpected: false,
			prepareDir:    DirCreator,
		},
		{
			name: "Error when FolderName is not a directory",
			req: contract.TreeRequest{
				FolderName: "notADir.txt",
			},
			errorExpected: true,
			errorMessage:  "notADir.txt [error opening dir]",
			prepareDir:    DirCreator,
		},
		{
			name: "Happy Path with Empty FolderName (uses current directory)",
			req: contract.TreeRequest{
				FolderName: "",
			},
			expectedResp: &contract.TreeResponse{
				Root: &contract.TreeNode{
					Name:         "handler",
					IsDir:        true,
					Path:         "handler",
					RelativePath: "handler",
				},
			},
			errorExpected: false,
			prepareDir:    DirCreator,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare the test environment
			if tt.req.FolderName != "" {
				err := tt.prepareDir(tt.req.FolderName, tt.fileNames, tt.filesData)
				if err != nil {
					t.Fatalf("error preparing test directory: %v", err)
				}
			}
			defer func() {
				_ = os.RemoveAll(tt.req.FolderName)
			}()

			// Execute the function
			resp, err := ProcessTreeRequest(tt.req)

			// Validate error
			if tt.errorExpected {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMessage)
				return
			} else {
				assert.NoError(t, err)
			}

			// Validate response
			if tt.req.FolderName != "" {
				assert.Equal(t, tt.expectedResp.Root, resp.Root)
			}
		})
	}
}
