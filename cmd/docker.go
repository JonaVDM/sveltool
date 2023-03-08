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

		fmt.Print("Copying .dockerignore: ")
		if err := gen.DockerIgnore(); err != nil {
			fmt.Println("something went wrong!", err)
		} else {
			fmt.Println("Ok")
		}

		fmt.Print("Copying Dockerfile: ")
		if err := gen.BasicDockerFile(); err != nil {
			fmt.Println("something went wrong!", err)
		} else {
			fmt.Println("Ok")
		}
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
	dockerCmd.Flags().BoolP("skip-install", "s", false, "Skip the install of the node adapter")
}
