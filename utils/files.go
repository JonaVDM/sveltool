package utils

import (
	"errors"
	"fmt"
	"os"
)

func CreateFolder(folder string) error {
	_, err := os.Stat(folder)
	if errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			return err
		}
		return nil
	}
	return err
}

func RunTemplate(file string, templ func() error) {
	fmt.Print("Copying " + file + ": ")
	if err := templ(); err != nil {
		fmt.Println("something went wrong!", err)
	} else {
		fmt.Println("Ok")
	}
}
