package pocketbase

import (
	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

var composeCmd = &cobra.Command{
	Use:     "docker",
	Short:   "Creates a docker compose file with pocketbase & some tools",
	Aliases: []string{"compose"},
	Run: func(cmd *cobra.Command, args []string) {
		utils.SimpleTemplate("docker/pocketbase-docker-compose.yml", "docker-compose.yml")
	},
}

func init() {
	PocketbaseCmd.AddCommand(composeCmd)
}
