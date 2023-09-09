package pocketbase

import (
	"fmt"

	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

// PocketbaseCmd represents the pocketbase command
var PocketbaseCmd = &cobra.Command{
	Use:   "pocketbase",
	Short: "Install different parts of pocketbase",
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.CreateFolder("src/lib"); err != nil {
			fmt.Println(err)
			return
		}

		utils.InstallPackages(cmd, "pocketbase")
		utils.SimpleTemplate("pocketbase/pocketbase.ts", "src/lib/pocketbase.ts")
	},
}

func init() {
}
