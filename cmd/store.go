package cmd

import (
	"fmt"

	"github.com/jonavdm/sveltool/gen"
	"github.com/spf13/cobra"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store [storename]",
	Short: "Generate a quick store",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := gen.SvelteStore(args[0]); err != nil {
			fmt.Println("error: ", err)
		} else {
			fmt.Println("Should be there now")
		}
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
}
