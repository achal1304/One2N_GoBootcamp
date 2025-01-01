package wcerrors

import (
	"errors"
	"fmt"
	"os"
)

type WcError struct {
	Err      error
	FileName string
}

func HandleErrors(wcError WcError) WcError {
	if errors.Is(wcError.Err, os.ErrNotExist) {
		return WcError{
			Err:      fmt.Errorf("wc: %s: read: %s", wcError.FileName, "No such file or directory\n"),
			FileName: wcError.FileName,
		}
	} else if errors.Is(wcError.Err, os.ErrPermission) {
		return WcError{
			Err:      fmt.Errorf("wc: %s: read: %s", wcError.FileName, "Permission denied\n"),
			FileName: wcError.FileName,
		}
	} else {
		return WcError{
			Err:      fmt.Errorf("wc: %s\n", wcError.Err.Error()),
			FileName: wcError.FileName,
		}
	}
}
