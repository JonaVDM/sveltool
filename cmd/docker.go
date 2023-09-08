package cmd

import (
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
			utils.InstallPackages(cmd, "@sveltejs/adapter-node")
		}

		utils.RunTemplate(".dockerignore", gen.DockerIgnore)
		utils.RunTemplate("Dockerfile", gen.BasicDockerFile)
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
	dockerCmd.Flags().BoolP("skip-install", "s", false, "Skip the install of the node adapter")
}
