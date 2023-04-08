package cmd

import (
	"fmt"

	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

// superformsCmd represents the superforms command
var superformsCmd = &cobra.Command{
	Use:   "superforms",
	Short: "Installs the packages needed for superforms",
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.InstallPackages("pnpm", "sveltekit-superforms", "zod"); err != nil {
			fmt.Println("Something went wrong whilst installing packges:", err)
			return
		}

		fmt.Println("Installed: Superforms, Zod")
	},
}

func init() {
	rootCmd.AddCommand(superformsCmd)
}
