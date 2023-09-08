package packages

import "github.com/spf13/cobra"

var PackageCmd = &cobra.Command{
	Use:     "package",
	Short:   "A collection of usefull packages to download",
	Aliases: []string{"packages", "install"},
}

func init() {}
