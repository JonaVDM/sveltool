package gen

import (
	"embed"
	"errors"
	"os"
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

func PicoCssLayout() error {
	return simpleTemplate("misc/picocss-layout.svelte", "src/routes/+layout.svelte")
}
