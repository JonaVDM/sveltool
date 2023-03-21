package cmd

import (
	"fmt"

	"github.com/jonavdm/sveltool/gen"
	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

// dockerCmd represents the docker command
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Generate a semi-useful dockerfile",
	Run: func(cmd *cobra.Command, args []string) {
		if skip, _ := cmd.Flags().GetBool("skip-install"); !skip {
			if err := utils.InstallPackages("pnpm", "@sveltejs/adapter-node"); err != nil {
				fmt.Println("Error while installing packages", err)
				return
			}
		}

		utils.RunTemplate(".dockerignore", gen.DockerIgnore)
		utils.RunTemplate("Dockerfile", gen.BasicDockerFile)
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
	dockerCmd.Flags().BoolP("skip-install", "s", false, "Skip the install of the node adapter")
}
