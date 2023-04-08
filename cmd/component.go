package cmd

import (
	"github.com/spf13/cobra"
)

// componentCmd represents the component command
var componentCmd = &cobra.Command{
	Use:   "component",
	Short: "A collection of some components",
}

func init() {
	rootCmd.AddCommand(componentCmd)
}
