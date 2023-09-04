package gen

import "github.com/jonavdm/sveltool/utils"

func CheckboxComponent() error {
	if err := utils.CreateFolder("src/lib/inputs"); err != nil {
		return err
	}

	return simpleTemplate("components/Checkbox.svelte.stub", "src/lib/inputs/Checkbox.svelte")
}

func TextboxComponent() error {
	if err := utils.CreateFolder("src/lib/inputs"); err != nil {
		return err
	}

	return simpleTemplate("components/Textbox.svelte.stub", "src/lib/inputs/Textbox.svelte")
}

func TextInputComponents() error {
	if err := utils.CreateFolder("src/lib/inputs"); err != nil {
		return err
	}

	return simpleTemplate("components/TextInput.svelte.stub", "src/lib/inputs/TextInput.svelte")
}
