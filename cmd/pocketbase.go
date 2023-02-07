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

		fmt.Println("\nCreating pocketbase config")
		if err := gen.PocketBase(); err != nil {
			fmt.Printf("Could not create: %s\n", err.Error())
		}

		fmt.Println("Insert the following interface into src/app.d.ts:")
		fmt.Println(`
interface Locals {
	pb: import('pocketbase').default,
}
	`)
	},
}

func init() {
	rootCmd.AddCommand(pocketbaseCmd)

	pocketbaseCmd.Flags().StringP("manager", "m", "pnpm", "Set node package manager")
}
