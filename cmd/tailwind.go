package cmd

import (
	"fmt"

	"github.com/jonavdm/sveltool/gen"
	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

// tailwindCmd represents the tailwind command
var tailwindCmd = &cobra.Command{
	Use:   "tailwind",
	Short: "Add Tailwind to the application",
	Long: `Add Tailwind to the application.
This install tailwind and sets up the configuration files for it.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing packages")
		manager, _ := cmd.Flags().GetString("manager")
		if err := utils.InstallPackages(manager, "tailwindcss", "postcss", "autoprefixer"); err != nil {
			fmt.Printf("Could not intsall tailwind: %s\n", err.Error())
		}

		fmt.Println("\nCreating config for tailwind")
		if err := gen.TailwindConfig(); err != nil {
			fmt.Printf("Could not create: %s\n", err.Error())
		}

		fmt.Println("\nCreating config for postcss")
		if err := gen.PostCssConfig(); err != nil {
			fmt.Printf("Could not create: %s\n", err.Error())
		}

		fmt.Println("\nCreating style.css")
		if err := gen.TailwindStyles(); err != nil {
			fmt.Printf("Could not create: %s\n", err.Error())
		}

		fmt.Println("\nCreating +layout.svelte")
		if err := gen.TailwindLayout(); err != nil {
			fmt.Printf("Could not create: %s\n", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(tailwindCmd)

	tailwindCmd.Flags().StringP("manager", "m", "pnpm", "Set node package manager")
}
