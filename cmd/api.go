package cmd

import (
	"fmt"

	"github.com/jonavdm/sveltool/gen"
	"github.com/spf13/cobra"
)

// apiWrapperCmd represents the apiWrapper command
var apiWrapperCmd = &cobra.Command{
	Use:   "api",
	Short: "Generate a simple API wrapper, works okay with basic auth",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Copying api.ts: ")
		if err := gen.BasicAuthLayout(); err != nil {
			fmt.Println("something went wrong!", err)
		} else {
			fmt.Println("Ok")
		}

		fmt.Println("Add the following code to app.d.ts")
		fmt.Println(`
interface Locals {
  ...
  api: import('$lib/api').Api,
}`)
	},
}

func init() {
	rootCmd.AddCommand(apiWrapperCmd)
}
