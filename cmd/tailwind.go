package cmd

import (
	"fmt"

	"github.com/jonavdm/sveltool/gen"
	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

// tailwindCmd represents the tailwind command
var tailwindCmd = &cobra.Command{
	Use:   "tailwind",
	Short: "Add Tailwind to the application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing packages")
		manager, _ := cmd.Flags().GetString("manager")
		daisy, _ := cmd.Flags().GetBool("daisy")
		if err := utils.InstallPackages(manager, "tailwindcss", "postcss", "autoprefixer", "@tailwindcss/typography"); err != nil {
			fmt.Printf("Could not intsall tailwind: %s\n", err.Error())
		}

		if daisy {
			if err := utils.InstallPackages(manager, "daisyui"); err != nil {
				fmt.Printf("Could not intsall daisy: %s\n", err.Error())
			}
		}

		utils.RunTemplate("tailwind.config.cjs", func() error { return gen.TailwindConfig(daisy) })
		utils.RunTemplate("postcss.config.cjs", gen.PostCssConfig)
		utils.RunTemplate("style.css", gen.TailwindStyles)
		utils.RunTemplate("+layout.svelte", gen.TailwindLayout)
	},
}

func init() {
	rootCmd.AddCommand(tailwindCmd)

	tailwindCmd.Flags().StringP("manager", "m", "pnpm", "Set node package manager")
	tailwindCmd.Flags().Bool("daisy", false, "Add in daisy ui")
}
