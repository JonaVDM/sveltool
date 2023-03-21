package cmd

import (
	"fmt"
	"strings"

	"github.com/jonavdm/sveltool/gen"
	"github.com/spf13/cobra"
)

// actionCmd represents the action command
var actionCmd = &cobra.Command{
	Use:   "action [names:type]",
	Short: "Generate a simple form action to go along with a form",
	Run: func(cmd *cobra.Command, args []string) {
		fields := make([]gen.FormItem, 0)
		for _, arg := range args {
			spl := strings.Split(arg, ":")
			fields = append(fields, gen.FormItem{
				Name: spl[0],
				Kind: spl[1],
			})
		}

		if err := gen.FormActions(fields); err != nil {
			fmt.Println("Error: ", err)
		}
	},
}

func init() {
	formCmd.AddCommand(actionCmd)
}
