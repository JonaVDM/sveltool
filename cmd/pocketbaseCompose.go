package cmd

import (
	"github.com/jonavdm/sveltool/gen"
	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

// pocketbaseComposeCmd represents the pocketbaseCompose command
var pocketbaseComposeCmd = &cobra.Command{
	Use:     "docker",
	Short:   "Creates a docker compose file with pocketbase & some tools",
	Aliases: []string{"compose"},
	Run: func(cmd *cobra.Command, args []string) {
		utils.RunTemplate("docker-compose.yml", gen.PocketbaseDocker)
	},
}

func init() {
	pocketbaseCmd.AddCommand(pocketbaseComposeCmd)
}
