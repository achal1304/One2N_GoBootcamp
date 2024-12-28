package wchandler

import (
	"errors"
	"os"
	"testing"

	"github.com/achal1304/One2N_GoBootcamp/wordcount/contract"
	"github.com/stretchr/testify/assert"
)

func TestProcessWCCommand(t *testing.T) {
	tests := []struct {
		name           string
		fileName       string
		prepare        func(fileName string) error
		expectedValues contract.WcValues
		setFlags       contract.WcFlags
		expectedErr    error
	}{
		{
			name:     "HappyPathLineCount",
			fileName: "happyline.txt",
			prepare: func(fileName string) error {
				lines := "line1\nline2"
				return os.WriteFile(fileName, []byte(lines), 0644)
			},
			expectedValues: contract.WcValues{LineCount: 2, WordCount: 2},
			setFlags:       contract.WcFlags{LineCount: true},
			expectedErr:    nil,
		},
		{
			name:     "HappyPathLineWhiteSpacesCount",
			fileName: "happyline.txt",
			prepare: func(fileName string) error {
				lines := "            \n             \n   "
				return os.WriteFile(fileName, []byte(lines), 0644)
			},
			expectedValues: contract.WcValues{LineCount: 3, WordCount: 0},
			setFlags:       contract.WcFlags{LineCount: true},
			expectedErr:    nil,
		},
		{
			name:     "HappyPathWordCount",
			fileName: "happyword.txt",
			prepare: func(fileName string) error {
				lines := "line1-line1new\nline2"
				return os.WriteFile(fileName, []byte(lines), 0644)
			},
			expectedValues: contract.WcValues{WordCount: 2, LineCount: 2},
			setFlags:       contract.WcFlags{WordCount: true},
			expectedErr:    nil,
		},
		{
			name:     "HappyPathWordWhiteSpacesCount",
			fileName: "happyword.txt",
			prepare: func(fileName string) error {
				lines := "                                              "
				return os.WriteFile(fileName, []byte(lines), 0644)
			},
			expectedValues: contract.WcValues{WordCount: 0, LineCount: 1},
			setFlags:       contract.WcFlags{WordCount: true},
			expectedErr:    nil,
		},
		{
			name:     "NoFileFound",
			fileName: "notfound.txt",
			prepare: func(fileName string) error {
				return nil
			},
			expectedValues: contract.WcValues{LineCount: 0},
			setFlags:       contract.WcFlags{LineCount: true},
			expectedErr:    os.ErrNotExist,
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
			expectedValues: contract.WcValues{LineCount: 2, WordCount: 4},
			setFlags:       contract.WcFlags{LineCount: true},
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

			actualCount, actualErr := ProcessWCCommand(tt.fileName, tt.setFlags)
			if tt.expectedErr != nil {
				assert.Error(t, actualErr)
				assert.True(t, errors.Is(actualErr, tt.expectedErr))
			}

			tt.expectedValues.FileName = tt.fileName
			assert.Equal(t, tt.expectedValues, actualCount)
		})
	}
}
