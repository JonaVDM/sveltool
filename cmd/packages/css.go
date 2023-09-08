package packages

import (
	"fmt"

	"github.com/jonavdm/sveltool/gen"
	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

var picoCmd = &cobra.Command{
	Use:     "picocss",
	Short:   "Adds Pico.css, really easy basic styling",
	Aliases: []string{"pico"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing Pico.css")
		manager, _ := cmd.Flags().GetString("manager")
		if err := utils.InstallPackages(manager, "@picocss/pico"); err != nil {
			fmt.Printf("Could not intsall Pico.css: %s\n", err.Error())
		}

		utils.RunTemplate("+layout.svelte", gen.PicoCssLayout)
	},
}

func init() {
	PackageCmd.AddCommand(picoCmd)
}
