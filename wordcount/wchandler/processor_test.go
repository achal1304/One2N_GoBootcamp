package wchandler

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessWCCommand(t *testing.T) {
	tests := []struct {
		name        string
		fileName    string
		prepare     func(fileName string) error
		expected    int
		expectedErr error
	}{
		{
			name:     "HappyPathLineCount",
			fileName: "happy.txt",
			prepare: func(fileName string) error {
				lines := "line1\nline2"
				return os.WriteFile(fileName, []byte(lines), 0644)
			},
			expected:    2,
			expectedErr: nil,
		},
		{
			name:     "NoFileFound",
			fileName: "notfound.txt",
			prepare: func(fileName string) error {
				return nil
			},
			expected:    0,
			expectedErr: os.ErrNotExist,
		},
		{
			name:     "ReadPermissionDenied",
			fileName: "read.txt",
			prepare: func(fileName string) error {
				lines := "Line 1\nLine 2\n"
				err := os.WriteFile(fileName, []byte(lines), 0000)
				if err != nil {
					return err
				}
				return nil
			},
			expected:    2,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.prepare(tt.fileName)
			if err != nil {
				t.Error("unable to prepare file", err)
				return
			}
			defer os.Remove(tt.fileName)

			actualCount, actualErr := ProcessWCCommand(tt.fileName)
			if tt.expectedErr != nil {
				assert.Error(t, actualErr)
				assert.True(t, errors.Is(actualErr, tt.expectedErr))
			}

			assert.Equal(t, tt.expected, actualCount)
		})
	}
}
