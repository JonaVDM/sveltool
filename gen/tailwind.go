package gen

func TailwindConfig(daisy bool) error {
	return complexTemplate("tailwind/tailwind.config.cjs.stub", "tailwind.config.cjs", map[string]bool{
		"Daisy": daisy,
	})
}

func PostCssConfig() error {
	return simpleTemplate("tailwind/postcss.config.cjs", "postcss.config.cjs")
}

func TailwindStyles() error {
	return simpleTemplate("tailwind/app.css", "src/app.css")
}

func TailwindLayout() error {
	return simpleTemplate("tailwind/+layout.svelte", "src/routes/+layout.svelte")
}
