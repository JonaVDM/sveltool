package cmd

import (
	"fmt"

	"github.com/jonavdm/sveltool/gen"
	"github.com/spf13/cobra"
)

// basicAuthCmd represents the basicAuth command
var basicAuthCmd = &cobra.Command{
	Use:   "basic-auth",
	Short: "Adds (really) basic authentication in svelte (batteries not included)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("basicAuth called")

		fmt.Print("Copying +layout.server.ts: ")
		if err := gen.BasicAuthLayout(); err != nil {
			fmt.Println("something went wrong!", err)
		} else {
			fmt.Println("Ok")
		}

		fmt.Print("Copying (login) +page.server.ts: ")
		if err := gen.BasicAuthLoginServer(); err != nil {
			fmt.Println("something went wrong!", err)
		} else {
			fmt.Println("Ok")
		}

		fmt.Print("Copying (login) +page.svelte: ")
		if err := gen.BasicAuthLogin(); err != nil {
			fmt.Println("something went wrong!", err)
		} else {
			fmt.Println("Ok")
		}

		fmt.Print("Copying (logout) +server.ts: ")
		if err := gen.BasicAuthLogout(); err != nil {
			fmt.Println("something went wrong!", err)
		} else {
			fmt.Println("Ok")
		}

		fmt.Print("Copying hooks.client.ts: ")
		if err := gen.BasicAuthClientHook(); err != nil {
			fmt.Println("something went wrong!", err)
		} else {
			fmt.Println("Ok")
		}

		fmt.Print("Copying hooks.server.ts: ")
		if err := gen.BasicAuthServerHook(); err != nil {
			fmt.Println("something went wrong!", err)
		} else {
			fmt.Println("Ok")
		}
	},
}

func init() {
	rootCmd.AddCommand(basicAuthCmd)
}
