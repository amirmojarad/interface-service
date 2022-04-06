package handlers

import (
	"os"
)

func Open() (*os.File, error) {
	if file, err := os.Open("subs/Rick and Morty.English.srt"); err != nil {
		return nil, err
	} else {
		return file, nil
	}
}
