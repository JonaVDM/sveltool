package cmd

import (
	"fmt"

	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

// basicAuthCmd represents the basicAuth command
var basicAuthCmd = &cobra.Command{
	Use:   "basic-auth",
	Short: "Adds (really) basic authentication in svelte (batteries not included)",
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.CreateFolder("src/routes/(auth)"); err != nil {
			fmt.Println(err)
			return
		}
		if err := utils.CreateFolder("src/routes/(auth)/login"); err != nil {
			fmt.Println(err)
			return
		}
		if err := utils.CreateFolder("src/routes/(auth)/logout"); err != nil {
			fmt.Println(err)
			return
		}

		utils.SimpleTemplate("basic_auth/+layout.server.ts", "src/routes/(auth)/+layout.server.ts")
		utils.SimpleTemplate("basic_auth/login/+page.ts", "src/routes/(auth)/login/+page.ts")
		utils.SimpleTemplate("basic_auth/login/+page.svelte", "src/routes/(auth)/login/+page.svelte")
		utils.SimpleTemplate("basic_auth/logout/+server.ts", "src/routes/(auth)/logout/+server.ts")
		utils.SimpleTemplate("basic_auth/hooks.client.ts", "src/hooks.client.ts")
		utils.SimpleTemplate("basic_auth/hooks.server.ts", "src/hooks.server.ts")
	},
}

func init() {
	rootCmd.AddCommand(basicAuthCmd)
}
