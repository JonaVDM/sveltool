package gen

func TailwindConfig() error {
	return simpleTemplate("tailwind/tailwind.config.cjs", "tailwind.config.cjs")
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
