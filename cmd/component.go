package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jonavdm/sveltool/cmd/components"
)

// componentCmd represents the component command
var componentCmd = &cobra.Command{
	Use:   "component",
	Short: "A collection of some components",
}

func init() {
	rootCmd.AddCommand(componentCmd)

	componentCmd.AddCommand(components.InputCmd)
}
