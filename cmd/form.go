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
		binds, err := cmd.Flags().GetBool("binds")
		if err != nil {
			fmt.Println(err)
			return
		}

		fields := make([]gen.FormItem, 0)
		for _, arg := range args {
			spl := strings.Split(arg, ":")
			fields = append(fields, gen.FormItem{
				Name:  spl[0],
				Kind:  spl[1],
				Label: utils.FirstLetterUpper(spl[0]),
			})
		}

		if err := gen.Form(fields, binds); err != nil {
			fmt.Println("Error: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(formCmd)

	formCmd.Flags().BoolP("binds", "b", false, "If the form needs to use bindings")
}
