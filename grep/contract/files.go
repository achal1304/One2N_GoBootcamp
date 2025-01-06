package contract

import "os"

type FileReader struct {
	File    *os.File
	ReadErr error
}
