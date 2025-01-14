package handler

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/achal1304/One2N_GoBootcamp/tree/contract"
	"github.com/stretchr/testify/assert"
)

func TestWriteJSON(t *testing.T) {
	tests := []struct {
		name           string
		req            contract.TreeRequest
		response       contract.TreeResponse
		expectedOutput []interface{}
	}{
		{
			name: "Single Directory with One File",
			req:  contract.TreeRequest{},
			response: contract.TreeResponse{
				Root: &contract.TreeNode{
					Name:  "root",
					IsDir: true,
					NextDir: []*contract.TreeNode{
						{Name: "file1.txt", IsDir: false},
					},
				},
				DirectoryCount: 1,
				FileCount:      1,
			},
			expectedOutput: []interface{}{
				map[string]interface{}{
					"type": "directory",
					"name": "root",
					"contents": []interface{}{
						map[string]interface{}{
							"type": "file",
							"name": "file1.txt",
						},
					},
				},
				map[string]interface{}{
					"type":        "report",
					"directories": float64(1),
					"files":       float64(1),
				},
			},
		},
		{
			name: "Empty Directory",
			req:  contract.TreeRequest{},
			response: contract.TreeResponse{
				Root: &contract.TreeNode{
					Name:    "emptyDir",
					IsDir:   true,
					NextDir: []*contract.TreeNode{},
				},
				DirectoryCount: 0,
				FileCount:      0,
			},
			expectedOutput: []interface{}{
				map[string]interface{}{
					"type": "directory",
					"name": "emptyDir",
				},
				map[string]interface{}{
					"type":        "report",
					"directories": float64(0),
					"files":       float64(0),
				},
			},
		},
		{
			name: "Nested Directories with Files",
			req:  contract.TreeRequest{},
			response: contract.TreeResponse{
				Root: &contract.TreeNode{
					Name:  "root",
					IsDir: true,
					NextDir: []*contract.TreeNode{
						{
							Name:  "dir1",
							IsDir: true,
							NextDir: []*contract.TreeNode{
								{Name: "file1.txt", IsDir: false},
								{Name: "file2.txt", IsDir: false},
							},
						},
						{Name: "file3.txt", IsDir: false},
					},
				},
				DirectoryCount: 2,
				FileCount:      3,
			},
			expectedOutput: []interface{}{
				map[string]interface{}{
					"type": "directory",
					"name": "root",
					"contents": []interface{}{
						map[string]interface{}{
							"type": "directory",
							"name": "dir1",
							"contents": []interface{}{
								map[string]interface{}{
									"type": "file",
									"name": "file1.txt",
								},
								map[string]interface{}{
									"type": "file",
									"name": "file2.txt",
								},
							},
						},
						map[string]interface{}{
							"type": "file",
							"name": "file3.txt",
						},
					},
				},
				map[string]interface{}{
					"type":        "report",
					"directories": float64(2),
					"files":       float64(3),
				},
			},
		},
		{
			name: "Directory Only Report",
			req: contract.TreeRequest{
				Flags: contract.TreeFlags{DirectoryPrint: true},
			},
			response: contract.TreeResponse{
				Root: &contract.TreeNode{
					Name:  "root",
					IsDir: true,
					NextDir: []*contract.TreeNode{
						{Name: "file1.txt", IsDir: false},
						{
							Name:  "subdir",
							IsDir: true,
						},
					},
				},
				DirectoryCount: 2,
				FileCount:      1,
			},
			expectedOutput: []interface{}{
				map[string]interface{}{
					"type": "directory",
					"name": "root",
					"contents": []interface{}{
						map[string]interface{}{
							"type": "file",
							"name": "file1.txt",
						},
						map[string]interface{}{
							"type": "directory",
							"name": "subdir",
						},
					},
				},
				map[string]interface{}{
					"type":        "report",
					"directories": float64(2),
				},
			},
		},
		{
			name: "Permission And Relative Path Flag",
			req:  contract.TreeRequest{Flags: contract.TreeFlags{RelativePath: true, Permission: true}},
			response: contract.TreeResponse{
				Root: &contract.TreeNode{
					Name:            "root",
					IsDir:           true,
					Permission:      "drwxrwxrwx",
					PermissionOctal: "0777",
					RelativePath:    "root",
					NextDir: []*contract.TreeNode{
						{Name: "file1.txt", IsDir: false, Permission: "-rw-rw-rw", PermissionOctal: "0666",
							RelativePath: "root/file1.txt",
						},
					},
				},
				DirectoryCount: 1,
				FileCount:      1,
			},
			expectedOutput: []interface{}{
				map[string]interface{}{
					"type": "directory",
					"mode": "0777",
					"prot": "drwxrwxrwx",
					"name": "root",
					"contents": []interface{}{
						map[string]interface{}{
							"type": "file",
							"mode": "0666",
							"prot": "-rw-rw-rw",
							"name": "root/file1.txt",
						},
					},
				},
				map[string]interface{}{
					"type":        "report",
					"directories": float64(1),
					"files":       float64(1),
				},
			},
		},
		{
			name: "Permission And Relative Path Flag With Graphics Option",
			req: contract.TreeRequest{Flags: contract.TreeFlags{RelativePath: true,
				Permission: true, Graphics: true}},
			response: contract.TreeResponse{
				Root: &contract.TreeNode{
					Name:            "root",
					IsDir:           true,
					Permission:      "drwxrwxrwx",
					PermissionOctal: "0777",
					RelativePath:    "root",
					NextDir: []*contract.TreeNode{
						{Name: "file1.txt", IsDir: false, Permission: "-rw-rw-rw", PermissionOctal: "0666",
							RelativePath: "root/file1.txt",
						},
					},
				},
				DirectoryCount: 1,
				FileCount:      1,
			},
			expectedOutput: []interface{}{
				map[string]interface{}{
					"type": "directory",
					"mode": "0777",
					"prot": "drwxrwxrwx",
					"name": "root",
					"contents": []interface{}{
						map[string]interface{}{
							"type": "file",
							"mode": "0666",
							"prot": "-rw-rw-rw",
							"name": "root/file1.txt",
						},
					},
				},
				map[string]interface{}{
					"type":        "report",
					"directories": float64(1),
					"files":       float64(1),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := WriteJSON(&buf, tt.req, tt.response)
			assert.NoError(t, err)

			var actualOutput []interface{}
			err = json.Unmarshal(buf.Bytes(), &actualOutput)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedOutput, actualOutput)
		})
	}
}
