package cmd

import (
	"fmt"

	"github.com/jonavdm/sveltool/gen"
	"github.com/spf13/cobra"
)

// pocketbaseComposeCmd represents the pocketbaseCompose command
var pocketbaseComposeCmd = &cobra.Command{
	Use:     "docker",
	Short:   "Creates a docker compose file with pocketbase & some tools",
	Aliases: []string{"compose"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nCreating docker-compose.yml")
		if err := gen.PocketbaseDocker(); err != nil {
			fmt.Printf("Could not create: %s\n", err.Error())
		}
	},
}

func init() {
	pocketbaseCmd.AddCommand(pocketbaseComposeCmd)
}
