package gen

import (
	"embed"
	"errors"
	"os"
	"text/template"
)

//go:embed templates/*
var templates embed.FS

func gen(filename string, contents []byte) error {
	if _, err := os.Stat(filename); err == nil {
		return errors.New("file already exists")
	}

	return os.WriteFile(filename, contents, 0664)
}

func simpleTemplate(source, dest string) error {
	contents, err := templates.ReadFile("templates/" + source)

	if err != nil {
		return err
	}

	return gen(dest, contents)
}

func complexTemplate(source, dest string, args any) error {
	if _, err := os.Stat(dest); err == nil {
		return errors.New("file already exists")
	}

	contents, err := templates.ReadFile("templates/" + source)
	if err != nil {
		return err
	}

	tmpl := template.Must(template.New(source).Parse(string(contents)))
	file, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer file.Close()
	return tmpl.Execute(file, args)
}
