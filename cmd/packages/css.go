package packages

import (
	"fmt"

	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

var picoCmd = &cobra.Command{
	Use:     "picocss",
	Short:   "Adds Pico.css, really easy basic styling",
	Aliases: []string{"pico"},
	Run: func(cmd *cobra.Command, args []string) {
		utils.InstallPackages(cmd, "@picocss/pico")

		if err := utils.CreateFolder("src/lib"); err != nil {
			fmt.Println(err)
			return
		}

		utils.SimpleTemplate("misc/picocss-layout.svelte", "src/routes/+layout.svelte")
	},
}

func init() {
	PackageCmd.AddCommand(picoCmd)
}
