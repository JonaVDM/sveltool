package cmd

import (
	"fmt"
	"strings"

	"github.com/jonavdm/sveltool/gen"
	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

// formCmd represents the form command
var formCmd = &cobra.Command{
	Use:   "form [name:type]",
	Short: "Generate a simple form in sveltekit",
	Run: func(cmd *cobra.Command, args []string) {
		fields := make([]gen.FormItem, 0)
		for _, arg := range args {
			spl := strings.Split(arg, ":")
			fields = append(fields, gen.FormItem{
				Name:  spl[0],
				Kind:  spl[1],
				Label: utils.FirstLetterUpper(spl[0]),
			})
		}

		if err := gen.Form(fields); err != nil {
			fmt.Println("Error: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(formCmd)
}
