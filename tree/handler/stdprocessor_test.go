package handler

import (
	"bytes"
	"testing"

	"github.com/achal1304/One2N_GoBootcamp/tree/contract"
	"github.com/stretchr/testify/assert"
)

func TestWritePlainText(t *testing.T) {
	tests := []struct {
		name           string
		response       contract.TreeResponse
		expectedOutput string
		req            contract.TreeRequest
	}{
		{
			name: "Happy Path - Single Directory with One Subdirectory",
			response: contract.TreeResponse{
				Root: &contract.TreeNode{
					Name:         "testDir",
					IsDir:        true,
					Path:         "testDir",
					RelativePath: "testDir",
					NextDir: []*contract.TreeNode{
						{
							Name:         "subDir",
							IsDir:        true,
							Path:         "testDir/subDir",
							RelativePath: "testDir/subDir",
						},
					},
				},
				DirectoryCount: 2,
				FileCount:      0,
			},
			expectedOutput: `testDir
|-- subDir

2 directories, 0 files
`,
		},
		{
			name: "Empty Directory",
			response: contract.TreeResponse{
				Root: &contract.TreeNode{
					Name:         "emptyDir",
					IsDir:        true,
					Path:         "emptyDir",
					RelativePath: "emptyDir",
					NextDir:      []*contract.TreeNode{},
				},
				DirectoryCount: 0,
				FileCount:      0,
			},
			expectedOutput: `emptyDir

0 directories, 0 files
`,
		},
		{
			name: "Directory Print Should Print Directories Only",
			response: contract.TreeResponse{
				Root: &contract.TreeNode{
					Name:         "testDir",
					IsDir:        true,
					Path:         "testDir",
					RelativePath: "testDir",
					NextDir: []*contract.TreeNode{
						{
							Name:         "subDir",
							IsDir:        true,
							Path:         "testDir/subDir",
							RelativePath: "testDir/subDir",
						},
					},
				},
				DirectoryCount: 2,
				FileCount:      0,
			},
			req: contract.TreeRequest{Flags: contract.TreeFlags{DirectoryPrint: true}},
			expectedOutput: `testDir
|-- subDir

2 directories
`,
		},
		{
			name: "Permission Print Should Print Permission With Names",
			response: contract.TreeResponse{
				Root: &contract.TreeNode{
					Name:         "testDir",
					IsDir:        true,
					Path:         "testDir",
					RelativePath: "testDir",
					Permission:   "drwxrwxrwx",
					NextDir: []*contract.TreeNode{
						{
							Name:         "subDir",
							IsDir:        true,
							Path:         "testDir/subDir",
							RelativePath: "testDir/subDir",
							Permission:   "drwxrwxrwx",
						},
						{
							Name:         "file1.txt",
							IsDir:        false,
							Path:         "testDir/file1.txt",
							RelativePath: "testDir/file1.txt",
							Permission:   "-rw-rw-rw-",
						},
					},
				},
				DirectoryCount: 2,
				FileCount:      1,
			},
			req: contract.TreeRequest{Flags: contract.TreeFlags{Permission: true}},
			expectedOutput: `[drwxrwxrwx] testDir
|-- [drwxrwxrwx] subDir
|-- [-rw-rw-rw-] file1.txt

2 directories, 1 files
`,
		},
		{
			name: "Permission Print With Relative Path",
			response: contract.TreeResponse{
				Root: &contract.TreeNode{
					Name:         "testDir",
					IsDir:        true,
					Path:         "testDir",
					RelativePath: "testDir",
					Permission:   "drwxrwxrwx",
					NextDir: []*contract.TreeNode{
						{
							Name:         "subDir",
							IsDir:        true,
							Path:         "testDir/subDir",
							RelativePath: "testDir/subDir",
							Permission:   "drwxrwxrwx",
						},
						{
							Name:         "file1.txt",
							IsDir:        false,
							Path:         "testDir/file1.txt",
							RelativePath: "testDir/file1.txt",
							Permission:   "-rw-rw-rw-",
						},
					},
				},
				DirectoryCount: 2,
				FileCount:      1,
			},
			req: contract.TreeRequest{Flags: contract.TreeFlags{Permission: true, RelativePath: true}},
			expectedOutput: `[drwxrwxrwx] testDir
|-- [drwxrwxrwx] testDir/subDir
|-- [-rw-rw-rw-] testDir/file1.txt

2 directories, 1 files
`,
		},
		{
			name: "Permission Print With Direcotory Only Path",
			response: contract.TreeResponse{
				Root: &contract.TreeNode{
					Name:         "testDir",
					IsDir:        true,
					Path:         "testDir",
					RelativePath: "testDir",
					Permission:   "drwxrwxrwx",
					NextDir: []*contract.TreeNode{
						{
							Name:         "subDir",
							IsDir:        true,
							Path:         "testDir/subDir",
							RelativePath: "testDir/subDir",
							Permission:   "drwxrwxrwx",
						},
					},
				},
				DirectoryCount: 2,
				FileCount:      0,
			},
			req: contract.TreeRequest{Flags: contract.TreeFlags{Permission: true, DirectoryPrint: true}},
			expectedOutput: `[drwxrwxrwx] testDir
|-- [drwxrwxrwx] subDir

2 directories
`,
		},
		{
			name: "Directory with Files and Subdirectories",
			response: contract.TreeResponse{
				Root: &contract.TreeNode{
					Name:         "testDir",
					IsDir:        true,
					Path:         "testDir",
					RelativePath: "testDir",
					NextDir: []*contract.TreeNode{
						{
							Name:         "subDir",
							IsDir:        true,
							Path:         "testDir/subDir",
							RelativePath: "testDir/subDir",
						},
						{
							Name:         "file1.txt",
							IsDir:        false,
							Path:         "testDir/file1.txt",
							RelativePath: "testDir/file1.txt",
						},
					},
				},
				DirectoryCount: 2,
				FileCount:      1,
			},
			expectedOutput: `testDir
|-- subDir
|-- file1.txt

2 directories, 1 files
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture output
			var output bytes.Buffer
			WritePlainText(&output, tt.req, tt.response)

			// Validate output
			assert.Equal(t, tt.expectedOutput, output.String())
		})
	}
}

func TestPrintTree(t *testing.T) {
	// Mock TreeNode structure
	tree := &contract.TreeNode{
		Name:  "root",
		IsDir: true,
		NextDir: []*contract.TreeNode{
			{
				Name:         "dir1",
				IsDir:        true,
				RelativePath: "root/dir1",
				NextDir: []*contract.TreeNode{
					{Name: "file1", IsDir: false, RelativePath: "root/dir1/file1"},
					{Name: "file2", IsDir: false, RelativePath: "root/dir1/file2"},
				},
			},
			{
				Name:         "dir2",
				IsDir:        true,
				RelativePath: "root/dir2",
				NextDir: []*contract.TreeNode{
					{Name: "file3", IsDir: false, RelativePath: "root/dir2/file3"},
				},
			},
			{
				Name:         "file4",
				IsDir:        false,
				RelativePath: "root/file4",
			},
			{
				Name:         "dir3",
				IsDir:        true,
				RelativePath: "root/dir3",
				NextDir: []*contract.TreeNode{
					{Name: "file5", IsDir: false, RelativePath: "root/dir3/file5"},
				},
			},
		},
	}

	tests := []struct {
		name     string
		req      contract.TreeRequest
		expected string
	}{
		{
			name: "Default Path",
			req:  contract.TreeRequest{},
			expected: `|-- dir1
|   |-- file1
|   |-- file2
|-- dir2
|   |-- file3
|-- file4
|-- dir3
    |-- file5
`,
		},
		{
			name: "With Relative Path",
			req: contract.TreeRequest{
				Flags: contract.TreeFlags{RelativePath: true},
			},
			expected: `|-- root/dir1
|   |-- root/dir1/file1
|   |-- root/dir1/file2
|-- root/dir2
|   |-- root/dir2/file3
|-- root/file4
|-- root/dir3
    |-- root/dir3/file5
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use a buffer to capture the output
			var buf bytes.Buffer

			// Call PrintTree
			PrintTree(&buf, tt.req, tree, 0, "")

			// Compare captured output with the expected output
			output := buf.String()
			if output != tt.expected {
				t.Errorf("Test case: %s\nExpected:\n%s\nBut got:\n%s", tt.name, tt.expected, output)
			}
		})
	}
}

func TestPrintStdOut(t *testing.T) {
	var buffer bytes.Buffer

	expectedOutput := "line1\nline2\nline3"
	PrintStdOut(&buffer, expectedOutput)
	actualOutput := buffer.String()
	assert.Equal(t, expectedOutput+"\n", actualOutput)
}
