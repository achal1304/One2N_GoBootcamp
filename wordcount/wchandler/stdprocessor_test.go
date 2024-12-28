package wchandler

import (
	"bytes"
	"testing"

	"github.com/achal1304/One2N_GoBootcamp/wordcount/contract"
	"github.com/stretchr/testify/assert"
)

func TestGenerateOutput(t *testing.T) {
	tests := []struct {
		name     string
		wcValues contract.WcValues
		wcFlags  contract.WcFlags
		expected string
	}{
		{
			name: "Only Line Count",
			wcValues: contract.WcValues{
				LineCount: 10,
				FileName:  "file1.txt",
			},
			wcFlags: contract.WcFlags{
				LineCount: true,
			},
			expected: "      10 file1.txt",
		},
		{
			name: "Only Word Count",
			wcValues: contract.WcValues{
				WordCount: 20,
				FileName:  "file2.txt",
			},
			wcFlags: contract.WcFlags{
				WordCount: true,
			},
			expected: "      20 file2.txt",
		},
		{
			name: "Only Character Count",
			wcValues: contract.WcValues{
				CharacterCount: 20,
				FileName:       "file2.txt",
			},
			wcFlags: contract.WcFlags{
				CharacterCount: true,
			},
			expected: "      20 file2.txt",
		},
		{
			name: "Line and Word Count",
			wcValues: contract.WcValues{
				LineCount: 15,
				WordCount: 25,
				FileName:  "file3.txt",
			},
			wcFlags: contract.WcFlags{
				LineCount: true,
				WordCount: true,
			},
			expected: "      15      25 file3.txt",
		},
		{
			name: "Line Word And Character Count",
			wcValues: contract.WcValues{
				LineCount:      15,
				WordCount:      25,
				CharacterCount: 50,
				FileName:       "file3.txt",
			},
			wcFlags:  contract.WcFlags{},
			expected: "      15      25      50 file3.txt",
		},
		{
			name: "No Flags And No Counts",
			wcValues: contract.WcValues{
				FileName: "file4.txt",
			},
			wcFlags:  contract.WcFlags{},
			expected: "       0       0       0 file4.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := GenerateOutput(tt.wcValues, tt.wcFlags)
			assert.Equal(t, tt.expected, output, "Output mismatch for test: %s", tt.name)
		})
	}
}

func TestPrintStdOut(t *testing.T) {
	var buffer bytes.Buffer

	expectedOutput := "       5       0      56 file.txt"
	PrintStdOut(&buffer, expectedOutput)
	actualOutput := buffer.String()
	assert.Equal(t, expectedOutput, actualOutput)
}
