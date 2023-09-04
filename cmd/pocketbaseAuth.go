package cmd

import (
	"fmt"

	"github.com/jonavdm/sveltool/gen"
	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

// pocketbaseAuthCmd represents the pocketbaseAuth command
var pocketbaseAuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Add authentication by pocketbasem, not including oauth (yet)",
	Run: func(cmd *cobra.Command, args []string) {
		utils.InstallPackages("pnpm", "zod", "sveltekit-superforms")

		// utils.RunTemplate("Input.svelte", gen.InputComponent)
		utils.RunTemplate("hooks.client.ts", gen.ClientHook)
		utils.RunTemplate("hooks.server.ts", gen.ServerHook)
		utils.RunTemplate("layout.server.ts", gen.AuthLayout)
		utils.RunTemplate("(login) page.server.ts", gen.LoginServer)
		utils.RunTemplate("(login) page.svelte", gen.LoginPage)
		utils.RunTemplate("(signup) page.server.ts", gen.SignupServer)
		utils.RunTemplate("(signup) page.svelte", gen.SignupPage)
		utils.RunTemplate("(logout) server.ts", gen.LogoutServer)

		fmt.Println(`
Lastly add the following to src/app.d.ts

interface Locals {
	pb: import('pocketbase').default,
	user: import('pocketbase').default['authStore']['model']
}`)
	},
}

func init() {
	pocketbaseCmd.AddCommand(pocketbaseAuthCmd)
}
