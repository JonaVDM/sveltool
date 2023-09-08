package packages

import (
	"github.com/jonavdm/sveltool/gen"
	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

var picoCmd = &cobra.Command{
	Use:     "picocss",
	Short:   "Adds Pico.css, really easy basic styling",
	Aliases: []string{"pico"},
	Run: func(cmd *cobra.Command, args []string) {
		utils.InstallPackages(cmd, "@picocss/pico")

		utils.RunTemplate("+layout.svelte", gen.PicoCssLayout)
	},
}

func init() {
	PackageCmd.AddCommand(picoCmd)
}
