package utils

import (
	"errors"
	"fmt"
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

func RunTemplate(file string, templ func() error) {
	fmt.Print("Copying " + file + ": ")
	if err := templ(); err != nil {
		fmt.Println("something went wrong!", err)
	} else {
		fmt.Println("Ok")
	}
}

func gen(filename string, contents []byte) error {
	if _, err := os.Stat(filename); err == nil {
		return errors.New("file already exists")
	}

	return os.WriteFile(filename, contents, 0664)
}

func SimpleTemplate(source, dest string) error {
	fmt.Print("Creating " + dest + ": ")
	contents, err := templates.Load(source)

	if err != nil {
		fmt.Println("Error! " + err.Error())
		return err
	}

	if err = gen(dest, contents); err != nil {
		fmt.Println("Error! " + err.Error())
		return err
	}

	fmt.Println("Ok")
	return nil
}

// func ComplexTemplate(source, dest string, args any) error {
// 	if _, err := os.Stat(dest); err == nil {
// 		return errors.New("file already exists")
// 	}

// 	contents, err := templates.ReadFile("templates/" + source)
// 	if err != nil {
// 		return err
// 	}

// 	tmpl := template.Must(template.New(source).Parse(string(contents)))
// 	file, err := os.Create(dest)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()
// 	return tmpl.Execute(file, args)
// }
