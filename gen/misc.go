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
