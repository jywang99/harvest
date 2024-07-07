package files

import (
	"os"
	"path/filepath"

	"jy.org/harvest/src/logging"
)

var logger = logging.Logger

func VerifyFileExists(path string) bool {
    _, err := os.Stat(path)
    return err == nil
}

func RemoveExt(path string) string {
    return path[:len(path)-len(filepath.Ext(path))]
}

func IsParent(parent, path string) bool {
    return parent == filepath.Dir(path)
}

