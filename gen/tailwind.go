package gen

func TailwindConfig(daisy bool) error {
	return complexTemplate("tailwind/tailwind.config.cjs.stub", "tailwind.config.cjs", map[string]bool{
		"Daisy": daisy,
	})
}
