package cmd

import (
	"github.com/jonavdm/sveltool/gen"
	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

// basicAuthCmd represents the basicAuth command
var basicAuthCmd = &cobra.Command{
	Use:   "basic-auth",
	Short: "Adds (really) basic authentication in svelte (batteries not included)",
	Run: func(cmd *cobra.Command, args []string) {
		utils.RunTemplate("+layout.server.ts", gen.BasicAuthLayout)
		utils.RunTemplate("(login) +page.server.ts", gen.BasicAuthLoginServer)
		utils.RunTemplate("(login) +page.svelte", gen.BasicAuthLogin)
		utils.RunTemplate("(logout) +server.ts", gen.BasicAuthLogout)
		utils.RunTemplate("hooks.client.ts", gen.BasicAuthClientHook)
		utils.RunTemplate("hooks.server.ts", gen.BasicAuthServerHook)
	},
}

func init() {
	rootCmd.AddCommand(basicAuthCmd)
}
