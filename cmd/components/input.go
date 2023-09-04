package components

import (
	"github.com/jonavdm/sveltool/gen"
	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

var InputCmd = &cobra.Command{
	Use:     "input",
	Short:   "A simple input component",
	Aliases: []string{"inputs"},
	Run: func(cmd *cobra.Command, args []string) {
		utils.RunTemplate("Checkbox.svelte", gen.CheckboxComponent)
		utils.RunTemplate("Textbox.svelte", gen.TextboxComponent)
		utils.RunTemplate("TextInput.svelte", gen.TextInputComponents)
	},
}

func init() {
}
