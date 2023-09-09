package components

import (
	"fmt"
	"log"

	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

// navCmd represents the nav command
var navCmd = &cobra.Command{
	Use:   "nav",
	Short: "Generate a simple nav bar.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.CreateFolder("src/lib"); err != nil {
			fmt.Println(err)
			return
		}

		authed, err := cmd.Flags().GetBool("authed")
		if err != nil {
			log.Fatal(err)
		}

		utils.ComplexTemplate("misc/nav.svelte", "src/lib/Nav.svelte", map[string]bool{
			"Authed": authed,
		})
	},
}

func init() {
	ComponentCmd.AddCommand(navCmd)
	navCmd.Flags().BoolP("authed", "a", false, "Separate authed/non-authed bars")
}
