package cmd

import (
	"fmt"

	"github.com/jonavdm/sveltool/gen"
	"github.com/spf13/cobra"
)

// actionCmd represents the action command
var actionCmd = &cobra.Command{
	Use:   "action [names]",
	Short: "Generate a simple form action to go along with a form",
	Run: func(cmd *cobra.Command, args []string) {
		if err := gen.FormActions(args); err != nil {
			fmt.Println("Error: ", err)
		}
	},
}

func init() {
	formCmd.AddCommand(actionCmd)
}
