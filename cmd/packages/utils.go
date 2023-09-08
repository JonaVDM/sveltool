package packages

import (
	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

var superFormCmd = &cobra.Command{
	Use:     "superforms",
	Short:   "Adds superforms",
	Aliases: []string{"superform", "forms", "form"},
	Run: func(cmd *cobra.Command, args []string) {
		utils.InstallPackages(cmd, "sveltekit-superforms", "zod")
	},
}

var graphCmd = &cobra.Command{
	Use:     "chartjs",
	Short:   "Adds chartjs",
	Aliases: []string{"chart", "graph"},
	Run: func(cmd *cobra.Command, args []string) {
		utils.InstallPackages(cmd, "chart.js")
	},
}

var iconifyCmd = &cobra.Command{
	Use:     "iconfiy",
	Short:   "Adds Iconify",
	Aliases: []string{"icon", "icons"},
	Run: func(cmd *cobra.Command, args []string) {
		utils.InstallPackages(cmd, "@iconify/svelte")
	},
}

func init() {
	PackageCmd.AddCommand(superFormCmd)
	PackageCmd.AddCommand(graphCmd)
	PackageCmd.AddCommand(iconifyCmd)
}
