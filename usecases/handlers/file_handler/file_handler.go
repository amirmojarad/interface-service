package file_handler

import (
	"os"
)

func Open(filePath string) (*os.File, error) {
	if file, err := os.Open(filePath); err != nil {
		return nil, err
	} else {
		return file, nil
	}
}
