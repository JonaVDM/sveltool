package gen

import (
	"os"
	"text/template"
)

type FormItem struct {
	Name  string
	Kind  string
	Label string
}

func Form(fields []FormItem) error {
	contents, err := templates.ReadFile("templates/form/form.html.stub")
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
	contents, err := templates.ReadFile("templates/form/action.ts.stub")
	if err != nil {
		return err
	}

	tmpl := template.Must(template.New("form-action").Parse(string(contents)))

	data := map[string]any{
		"Fields": fields,
	}

	return tmpl.Execute(os.Stdout, data)
}
