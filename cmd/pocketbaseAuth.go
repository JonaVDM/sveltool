package cmd

import (
	"fmt"

	"github.com/jonavdm/sveltool/gen"
	"github.com/spf13/cobra"
)

// pocketbaseAuthCmd represents the pocketbaseAuth command
var pocketbaseAuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Add authentication by pocketbasem, not including oauth (yet)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nCreating hooks.client.ts")
		if err := gen.ClientHook(); err != nil {
			fmt.Printf("Could not create: %s\n", err.Error())
		}

		fmt.Println("\nCreating hooks.server.ts")
		if err := gen.ServerHook(); err != nil {
			fmt.Printf("Could not create: %s\n", err.Error())
		}

		fmt.Println("\nCreating layout.server.ts")
		if err := gen.AuthLayout(); err != nil {
			fmt.Printf("Could not create: %s\n", err.Error())
		}

		fmt.Println("\nCreating login page.server.ts")
		if err := gen.LoginServer(); err != nil {
			fmt.Printf("Could not create: %s\n", err.Error())
		}

		fmt.Println("\nCreating login page.svelte")
		if err := gen.LoginPage(); err != nil {
			fmt.Printf("Could not create: %s\n", err.Error())
		}

		fmt.Println("\nCreating signup page.server.ts")
		if err := gen.SignupServer(); err != nil {
			fmt.Printf("Could not create: %s\n", err.Error())
		}

		fmt.Println("\nCreating signup page.svelte")
		if err := gen.SignupPage(); err != nil {
			fmt.Printf("Could not create: %s\n", err.Error())
		}

		fmt.Println("\nCreating logout server.ts")
		if err := gen.LogoutServer(); err != nil {
			fmt.Printf("Could not create: %s\n", err.Error())
		}

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
