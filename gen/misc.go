package gen

import (
	"strings"

	"github.com/jonavdm/sveltool/utils"
)

func PicoCssLayout() error {
	return simpleTemplate("misc/picocss-layout.svelte", "src/routes/+layout.svelte")
}

func SvelteStore(name string) error {
	if err := utils.CreateFolder("src/lib"); err != nil {
		return err
	}

	args := struct {
		Name       string
		ExportName string
	}{
		strings.Title(name),
		strings.ToLower(name),
	}

	return complexTemplate("misc/store.ts", "src/lib/"+strings.ToLower(name)+".ts", args)
}

func NavBar(auth bool) error {
	if err := utils.CreateFolder("src/lib"); err != nil {
		return err
	}

	args := struct {
		Auth bool
	}{
		auth,
	}

	return complexTemplate("misc/nav.svelte", "src/lib/Nav.svelte", args)
}

func BasicApi() error {
	if err := utils.CreateFolder("src/lib"); err != nil {
		return err
	}

	return simpleTemplate("misc/api.ts", "src/lib/api.ts")
}

func InputComponent() error {
	if err := utils.CreateFolder("src/lib"); err != nil {
		return err
	}

	return simpleTemplate("misc/Input.svelte.stub", "src/lib/Input.svelte")
}
