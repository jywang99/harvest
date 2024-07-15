package files

import (
	"os"
	"path/filepath"
	"strings"

	"jy.org/harvest/src/config"
	"jy.org/harvest/src/logging"
)

var logger = logging.Logger
var cfg = config.Config

func RemoveExt(path string) string {
    return path[:len(path)-len(filepath.Ext(path))]
}

func IsParent(parent, path string) bool {
    return parent == filepath.Dir(path)
}

func VerifyAndGetBasename(path string) (*string, bool) {
    _, err := os.Stat(path)
    if err != nil {
        logger.ERROR.Printf("Error when getting file info for %s: %v\n", path, err)
        return nil, false
    }

    var ppng *string
    png, err := filepath.Rel(cfg.Ingest.ThumbDir, path)
    if err != nil {
        logger.ERROR.Printf("Error when getting relative path for %s: %v\n", path, err)
        return nil, false
    }
    ppng = &png
    return ppng, true
}

func getExt(path string) string {
    ext := filepath.Ext(path)
    if ext == "" {
        return ""
    }
    return strings.ToLower(ext[1:])
}

func ignoreEntry(path string) bool {
    ignore := cfg.Ingest.IgnoreMap

    if !cfg.Ingest.DotFiles && strings.HasPrefix(filepath.Base(path), ".") {
        // logger.INFO.Printf("Skipping dot directory/file: %v\n", path)
        return true
    }

    if ignore[filepath.Base(path)] {
        // logger.INFO.Printf("Skipping ignored directory: %v\n", path)
        return true
    }

    return false
}

