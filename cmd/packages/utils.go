package packages

import (
	"fmt"

	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
)

var superFormCmd = &cobra.Command{
	Use:     "superforms",
	Short:   "Adds superforms",
	Aliases: []string{"superform", "forms", "form"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing Superforms, Zod")
		manager, _ := cmd.Flags().GetString("manager")
		if err := utils.InstallPackages(manager, "sveltekit-superforms", "zod"); err != nil {
			fmt.Println("Something went wrong whilst installing packges:", err)
			return
		}

		fmt.Println("Installed: Superforms, Zod")
	},
}

var graphCmd = &cobra.Command{
	Use:     "chartjs",
	Short:   "Adds chartjs",
	Aliases: []string{"chart", "graph"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing chart.js")
		manager, _ := cmd.Flags().GetString("manager")
		if err := utils.InstallPackages(manager, "chart.js"); err != nil {
			fmt.Println("Something went wrong whilst installing packges:", err)
			return
		}

		fmt.Println("Installed: chart.js")
	},
}

var iconifyCmd = &cobra.Command{
	Use:     "iconfiy",
	Short:   "Adds Iconify",
	Aliases: []string{"icon", "icons"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing @iconify/svelte")
		manager, _ := cmd.Flags().GetString("manager")
		if err := utils.InstallPackages(manager, "@iconify/svelte"); err != nil {
			fmt.Println("Something went wrong whilst installing packges:", err)
			return
		}

		fmt.Println("Installed: @iconify/svelte")
	},
}

func init() {
	PackageCmd.AddCommand(superFormCmd)
	PackageCmd.AddCommand(graphCmd)
	PackageCmd.AddCommand(iconifyCmd)
}
