package cmd

import (
	"os"

	"github.com/jonavdm/sveltool/cmd/components"
	"github.com/jonavdm/sveltool/cmd/packages"
	"github.com/jonavdm/sveltool/cmd/pocketbase"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sveltool",
	Short: "A simple tool that setups things in a svelte project",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sveltool.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringP("manager", "m", "pnpm", "Set node package manager")

	rootCmd.AddCommand(packages.PackageCmd)
	rootCmd.AddCommand(pocketbase.PocketbaseCmd)
	rootCmd.AddCommand(components.ComponentCmd)
}
