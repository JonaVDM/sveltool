package cmd

import (
	"github.com/jonavdm/sveltool/gen"
	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

// inputComponentCmd represents the inputComponent command
var inputComponentCmd = &cobra.Command{
	Use:   "input",
	Short: "A \"simple\" input component",
	Run: func(cmd *cobra.Command, args []string) {
		utils.RunTemplate("Input.svelte", gen.InputComponent)
	},
}

func init() {
	componentCmd.AddCommand(inputComponentCmd)
}
