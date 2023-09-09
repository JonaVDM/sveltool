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
		utils.SimpleTemplate("tailwind/postcss.config.cjs", "postcss.config.cjs")
		utils.SimpleTemplate("tailwind/app.css", "src/app.css")
		utils.SimpleTemplate("tailwind/+layout.svelte", "src/routes/+layout.svelte")
	},
}

func init() {
	rootCmd.AddCommand(tailwindCmd)

	tailwindCmd.Flags().StringP("manager", "m", "pnpm", "Set node package manager")
	tailwindCmd.Flags().Bool("daisy", false, "Add in daisy ui")
}
