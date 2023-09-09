package cmd

import (
	"strings"

	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store [storename]",
	Short: "Generate a quick store",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.ComplexTemplate("misc/store.ts", "src/lib/"+strings.ToLower(args[0])+".ts", map[string]string{
			"Name":       cases.Title(language.Georgian).String(args[0]),
			"ExportName": strings.ToLower(args[0]),
		})
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
}
