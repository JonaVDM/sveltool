package cmd

import (
	"github.com/jonavdm/sveltool/cmd/pocketbase"
	"github.com/jonavdm/sveltool/gen"
	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

// pocketbaseCmd represents the pocketbase command
var pocketbaseCmd = &cobra.Command{
	Use:   "pocketbase",
	Short: "Install different parts of pocketbase",
	Run: func(cmd *cobra.Command, args []string) {
		utils.InstallPackages(cmd, "pocketbase")
		utils.RunTemplate("pocketbase.ts", gen.PocketBase)
	},
}

func init() {
	rootCmd.AddCommand(pocketbaseCmd)

	pocketbaseCmd.Flags().StringP("manager", "m", "pnpm", "Set node package manager")
}
