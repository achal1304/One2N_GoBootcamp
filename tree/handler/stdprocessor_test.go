package handler

import (
	"bytes"
	"testing"

	"github.com/achal1304/One2N_GoBootcamp/tree/contract"
	"github.com/stretchr/testify/assert"
)

func TestPrintResponseStdOut(t *testing.T) {
	tests := []struct {
		name           string
		response       contract.TreeResponse
		expectedOutput string
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
				DirectoryCount: 1,
				FileCount:      1,
			},
			expectedOutput: `testDir
|-- subDir
|-- file1.txt

1 directories, 1 files
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture output
			var output bytes.Buffer
			PrintResponseStdOut(&output, tt.response)

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
				Name:  "dir1",
				IsDir: true,
				NextDir: []*contract.TreeNode{
					{Name: "file1", IsDir: false},
					{Name: "file2", IsDir: false},
				},
			},
			{
				Name:  "dir2",
				IsDir: true,
				NextDir: []*contract.TreeNode{
					{Name: "file3", IsDir: false},
				},
			},
			{
				Name:  "file4",
				IsDir: false,
			},
			{
				Name:  "dir3",
				IsDir: true,
				NextDir: []*contract.TreeNode{
					{Name: "file5", IsDir: false},
				},
			},
		},
	}

	// Use a buffer to capture the output
	var buf bytes.Buffer

	// Call PrintTree
	PrintTree(&buf, tree, 0, "")

	// Define the expected output
	expected := `|-- dir1
|   |-- file1
|   |-- file2
|-- dir2
|   |-- file3
|-- file4
|-- dir3
    |-- file5
`

	// Compare captured output with the expected output
	output := buf.String()
	if output != expected {
		t.Errorf("Expected:\n%s\nBut got:\n%s", expected, output)
	}
}

func TestPrintStdOut(t *testing.T) {
	var buffer bytes.Buffer

	expectedOutput := "line1\nline2\nline3"
	PrintStdOut(&buffer, expectedOutput)
	actualOutput := buffer.String()
	assert.Equal(t, expectedOutput+"\n", actualOutput)
}
