package utils

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func InstallPackages(cmd *cobra.Command, packages ...string) {
	manager, err := cmd.Flags().GetString("manager")
	if err != nil {
		panic("Cannot find package manager flag")
	}

	installCmd := "install"
	if manager == "yarn" {
		installCmd = "add"
	}

	fmt.Print("Installing packages " + strings.Join(packages, ", ") + " using " + manager + ": ")

	args := []string{installCmd, "-D"}
	args = append(args, packages...)

	if err := exec.Command(manager, args...).Run(); err != nil {
		fmt.Println("error! " + err.Error())
	} else {
		fmt.Println("ok")
	}
}
