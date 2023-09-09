package cmd

import (
	"fmt"

	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

// apiWrapperCmd represents the apiWrapper command
var apiWrapperCmd = &cobra.Command{
	Use:   "api",
	Short: "Generate a simple API wrapper, works okay with basic auth",
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.CreateFolder("src/lib"); err != nil {
			fmt.Println(err)
			return
		}

		utils.SimpleTemplate("misc/api.ts", "src/lib/api.ts")

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
