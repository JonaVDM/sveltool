package utils

import (
	"errors"
	"os"
)

func CreateFolder(folder string) error {
	_, err := os.Stat(folder)
	if errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(folder, os.ModePerm)
		if err != nil {
			return err
		}
		return nil
	}
	return err
}
