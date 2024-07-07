package files

import (
	"bufio"
	"os"
)

type FileReader struct{
    file *os.File
    scanner *bufio.Scanner
}

func NewFileReader(path string) (*FileReader, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
        logger.ERROR.Printf("Error when opening file: %v\n", err)
		return nil, err
	}

    return &FileReader{
        file: file,
        scanner: bufio.NewScanner(file),
    }, nil
}

func (fr *FileReader) ReadNextLine() (string, bool) {
    if fr.scanner.Scan() {
        return fr.scanner.Text(), true
    }
    return "", false
}

func (fr *FileReader) Close() {
    fr.scanner = nil
    fr.file.Close()
}

