package gen

import (
	"os"
	"text/template"

	"github.com/jonavdm/sveltool/templates"
)

type FormItem struct {
	Name  string
	Kind  string
	Label string
}

func Form(fields []FormItem) error {
	contents, err := templates.Load("form/form.html")
	if err != nil {
		return err
	}

	tmpl := template.Must(template.New("form-form").Parse(string(contents)))

	data := map[string]any{
		"Fields": fields,
	}

	return tmpl.Execute(os.Stdout, data)
}

func FormActions(fields []FormItem) error {
	contents, err := templates.Load("form/action.ts")
	if err != nil {
		return err
	}

	tmpl := template.Must(template.New("form-action").Parse(string(contents)))

	data := map[string]any{
		"Fields": fields,
	}

	return tmpl.Execute(os.Stdout, data)
}
