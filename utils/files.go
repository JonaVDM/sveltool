package utils

import (
	"errors"
	"fmt"
	"html/template"
	"os"

	"github.com/jonavdm/sveltool/templates"
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

func SimpleTemplate(source, dest string) error {
	fmt.Print("Creating " + dest + ": ")

	if _, err := os.Stat(dest); err == nil {
		fmt.Println("Error! file already exists")
		return errors.New("file already exists")
	}

	contents, err := templates.Load(source)
	if err != nil {
		fmt.Println("Error! " + err.Error())
		return err
	}

	if err = os.WriteFile(dest, contents, 0664); err != nil {
		fmt.Println("Error! " + err.Error())
		return err
	}

	fmt.Println("Ok")
	return nil
}

func ComplexTemplate(source, dest string, args any) error {
	fmt.Print("Creating " + dest + ": ")

	if _, err := os.Stat(dest); err == nil {
		fmt.Println("Error! file already exists")
		return errors.New("file already exists")
	}

	contents, err := templates.Load(source)
	if err != nil {
		fmt.Println("Error! " + err.Error())
		return err
	}

	tmpl := template.Must(template.New(source).Parse(string(contents)))
	file, err := os.Create(dest)
	if err != nil {
		fmt.Println("Error! " + err.Error())
		return err
	}
	defer file.Close()

	if err := tmpl.Execute(file, args); err != nil {
		fmt.Println("Error! " + err.Error())
		return err
	}

	return nil
}
