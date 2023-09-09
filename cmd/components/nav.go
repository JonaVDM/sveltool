package components

import (
	"log"

	"github.com/jonavdm/sveltool/gen"
	"github.com/spf13/cobra"
)

// navCmd represents the nav command
var navCmd = &cobra.Command{
	Use:   "nav",
	Short: "Generate a simple nav bar.",
	Run: func(cmd *cobra.Command, args []string) {
		authed, err := cmd.Flags().GetBool("authed")
		if err != nil {
			log.Fatal(err)
		}

		gen.NavBar(authed)
	},
}

func init() {
	ComponentCmd.AddCommand(navCmd)
	navCmd.Flags().BoolP("authed", "a", false, "Separate authed/non-authed bars")
}
