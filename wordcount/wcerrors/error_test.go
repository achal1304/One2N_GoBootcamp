package wcerrors

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleErrors(t *testing.T) {
	tests := []struct {
		name           string
		inputError     WcError
		expectedOutput WcError
	}{
		{
			name: "FileNotFoundError",
			inputError: WcError{
				Err:      os.ErrNotExist,
				FileName: "file1.txt",
			},
			expectedOutput: WcError{
				Err:      fmt.Errorf("wc: %s: read: %s", "file1.txt", "No such file or directory\n"),
				FileName: "file1.txt",
			},
		},
		{
			name: "PermissionDeniedError",
			inputError: WcError{
				Err:      os.ErrPermission,
				FileName: "file2.txt",
			},
			expectedOutput: WcError{
				Err:      fmt.Errorf("wc: %s: read: %s", "file2.txt", "Permission denied\n"),
				FileName: "file2.txt",
			},
		},
		{
			name: "GenericError",
			inputError: WcError{
				Err:      errors.New("some generic error"),
				FileName: "file3.txt",
			},
			expectedOutput: WcError{
				Err:      fmt.Errorf("wc: %s\n", "some generic error"),
				FileName: "file3.txt",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualOutput := HandleErrors(tt.inputError)

			// Validate the FileName
			assert.Equal(t, tt.expectedOutput.FileName, actualOutput.FileName)

			// Validate the Error
			assert.EqualError(t, actualOutput.Err, tt.expectedOutput.Err.Error())
		})
	}
}
