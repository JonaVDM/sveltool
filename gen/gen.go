package gen

import (
	"errors"
	"os"
	"text/template"

	"github.com/jonavdm/sveltool/templates"
)

func complexTemplate(source, dest string, args any) error {
	if _, err := os.Stat(dest); err == nil {
		return errors.New("file already exists")
	}

	contents, err := templates.Load(source)
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
