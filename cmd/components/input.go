package components

import (
	"fmt"

	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

var inputCmd = &cobra.Command{
	Use:     "input",
	Short:   "A simple input component",
	Aliases: []string{"inputs"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.CreateFolder("src/lib/inputs"); err != nil {
			fmt.Println("Error!")
			return
		}

		utils.SimpleTemplate("components/Checkbox.svelte", "src/lib/inputs/Checkbox.svelte")
		utils.SimpleTemplate("components/Textbox.svelte", "src/lib/inputs/Textbox.svelte")
		utils.SimpleTemplate("components/TextInput.svelte", "src/lib/inputs/TextInput.svelte")
	},
}

func init() {
}
