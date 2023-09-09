package components

import (
	"github.com/spf13/cobra"
)

var ComponentCmd = &cobra.Command{
	Use:   "component",
	Short: "A collection of some components",
}

func init() {
	ComponentCmd.AddCommand(inputCmd)
}
