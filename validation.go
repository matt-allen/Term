package term

import (
	"os"
)

func ValidateFilePath(s string) bool {
	_, err := os.Stat(s)
	return os.IsNotExist(err)
}
