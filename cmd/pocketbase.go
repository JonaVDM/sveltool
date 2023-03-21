package cmd

import (
	"fmt"

	"github.com/jonavdm/sveltool/gen"
	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

// pocketbaseCmd represents the pocketbase command
var pocketbaseCmd = &cobra.Command{
	Use:   "pocketbase",
	Short: "Install different parts of pocketbase",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing pocketbase")
		manager, _ := cmd.Flags().GetString("manager")
		if err := utils.InstallPackages(manager, "pocketbase"); err != nil {
			fmt.Printf("Could not intsall pocketbase: %s\n", err.Error())
		}

		utils.RunTemplate("pocketbase.ts", gen.PocketBase)
	},
}

func init() {
	rootCmd.AddCommand(pocketbaseCmd)

	pocketbaseCmd.Flags().StringP("manager", "m", "pnpm", "Set node package manager")
}
