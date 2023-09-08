package cmd

import (
	"github.com/jonavdm/sveltool/gen"
	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

// tailwindCmd represents the tailwind command
var tailwindCmd = &cobra.Command{
	Use:   "tailwind",
	Short: "Add Tailwind to the application",
	Run: func(cmd *cobra.Command, args []string) {
		daisy, _ := cmd.Flags().GetBool("daisy")
		utils.InstallPackages(cmd, "tailwindcss", "postcss", "autoprefixer", "@tailwindcss/typography")

		if daisy {
			utils.InstallPackages(cmd, "daisyui")
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
