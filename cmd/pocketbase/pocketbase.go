package pocketbase

import (
	"github.com/jonavdm/sveltool/gen"
	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

// PocketbaseCmd represents the pocketbase command
var PocketbaseCmd = &cobra.Command{
	Use:   "pocketbase",
	Short: "Install different parts of pocketbase",
	Run: func(cmd *cobra.Command, args []string) {
		utils.InstallPackages(cmd, "pocketbase")
		utils.RunTemplate("pocketbase.ts", gen.PocketBase)
	},
}

func init() {
}
