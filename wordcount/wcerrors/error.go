package wcerrors

import (
	"errors"
	"fmt"
	"os"
)

func HandleErrors(errorMessage error, fileName string) error {
	if errors.Is(errorMessage, os.ErrNotExist) {
		return fmt.Errorf("wc: %s: read: %s", fileName, "No such file or directory")
	} else if errors.Is(errorMessage, os.ErrPermission) {
		return fmt.Errorf("wc: %s: read: %s", fileName, "Permission denied")
	} else {
		return fmt.Errorf("wc: %s", errorMessage.Error())
	}
}
