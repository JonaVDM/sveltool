package gen

import (
	"errors"
	"os"
)

func gen(filename string, contents []byte) error {
	if _, err := os.Stat(filename); err == nil {
		return errors.New("file already exists")
	}

	return os.WriteFile(filename, contents, 0664)
}

func TailwindConfig() error {
	return gen("tailwind.config.cjs", tailwindConfigTemplate())
}

func PostCssConfig() error {
	return gen("postcss.config.cjs", tailwindPostcssConfigTemplate())
}

func TailwindStyles() error {
	return gen("src/app.css", tailwindStylesTemplate())
}

func TailwindLayout() error {
	return gen("src/routes/+layout.svelte", tailwindLayoutTemplate())
}