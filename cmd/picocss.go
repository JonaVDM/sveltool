package cmd

import (
	"fmt"

	"github.com/jonavdm/sveltool/gen"
	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

// picocssCmd represents the picocss command
var picocssCmd = &cobra.Command{
	Use:   "picocss",
	Short: "Adds Pico.css to the project for quick and easy styling",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing Pico.css")
		manager, _ := cmd.Flags().GetString("manager")
		if err := utils.InstallPackages(manager, "@picocss/pico"); err != nil {
			fmt.Printf("Could not intsall Pico.css: %s\n", err.Error())
		}

		fmt.Println("\nCreating +layout.svelte")
		if err := gen.PicoCssLayout(); err != nil {
			fmt.Printf("Could not create: %s\n", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(picocssCmd)
	picocssCmd.Flags().StringP("manager", "m", "pnpm", "Set node package manager")
}
