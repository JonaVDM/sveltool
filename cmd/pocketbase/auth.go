package pocketbase

import (
	"fmt"

	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Add authentication by pocketbasem, not including oauth",
	Run: func(cmd *cobra.Command, args []string) {
		utils.InstallPackages(cmd, "zod", "sveltekit-superforms")

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
		if err := utils.CreateFolder("src/routes/(auth)/signup"); err != nil {
			fmt.Println(err)
			return
		}

		utils.SimpleTemplate("pocketbase_auth/hooks.client.ts", "src/hooks.client.ts")
		utils.SimpleTemplate("pocketbase_auth/hooks.server.ts", "src/hooks.server.ts")
		utils.SimpleTemplate("pocketbase_auth/+layout.server.ts", "src/routes/(auth)/+layout.server.ts")
		utils.SimpleTemplate("pocketbase_auth/login/+page.server.ts", "src/routes/(auth)/login/+page.server.ts")
		utils.SimpleTemplate("pocketbase_auth/login/+page.svelte", "src/routes/(auth)/login/+page.svelte")
		utils.SimpleTemplate("pocketbase_auth/signup/+page.server.ts", "src/routes/(auth)/signup/+page.server.ts")
		utils.SimpleTemplate("pocketbase_auth/signup/+page.svelte", "src/routes/(auth)/signup/+page.svelte")
		utils.SimpleTemplate("pocketbase_auth/logout/+server.ts", "src/routes/(auth)/logout/+server.ts")

		fmt.Println(`
Lastly add the following to src/app.d.ts

interface Locals {
	pb: import('pocketbase').default,
	user: import('pocketbase').default['authStore']['model']
}`)
	},
}

func init() {
	PocketbaseCmd.AddCommand(authCmd)
}
