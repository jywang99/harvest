package files

import (
	"os"
	"path/filepath"
)

// get relative paths of all files in a directory, recursively
func GetFilesInDir(dir string) ([]string, error) {
    logger.INFO.Printf("Exploring directory: %v\n", dir)

    exts := cfg.Ingest.ExtMap

    files := make([]string, 0)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
        if info.IsDir() || ignoreEntry(path) {
            return nil
        }

        // no extension
        ext := getExt(path)
        if ext == "" {
            return nil
        }

        // add to list depending on file type
        if exts[ext] {
            rpath, _ := filepath.Rel(dir, path)
            files = append(files, rpath)
        }

        return nil
	})
    if err != nil {
        logger.ERROR.Printf("Error when reading directory")
        return nil, err
    }

    logger.INFO.Printf("Found %v files\n", len(files))
	return files, nil
}

