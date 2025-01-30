package check

import (
	"fmt"
	"os"

	"github.com/RevoltSecurities/Crlfix/crlfix/modules/logger"
)

func Permission(filename string) bool {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		if pathErr, ok := err.(*os.PathError); ok && pathErr.Err == os.ErrPermission {
			logger.Stdlogger(fmt.Sprintf("You don't have write permission in this %s file", filename), "warn")
			return false
		} else {
			logger.Stdlogger(fmt.Sprintf("Exception occured in the permission check module due to: %s", err.Error()), "warn")
			return false
		}
	}
	defer file.Close()
	return true
}
