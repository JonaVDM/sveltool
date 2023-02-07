package utils

import "os/exec"

func InstallPackages(manager string, packages ...string) error {
	installCmd := "install"
	if manager == "yarn" {
		installCmd = "add"
	}

	args := []string{installCmd, "-D"}

	for _, pack := range packages {
		args = append(args, pack)
	}

	return exec.Command(manager, args...).Run()
}
