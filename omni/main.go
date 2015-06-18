package main

import (
	"github.com/spf13/cobra"
)

var (
	platformBase = "/usr/lib/omni"
)

func main() {
	root := &cobra.Command{
		Use:   "omni",
		Short: "multiple automation platforms and versions with a single interface",
		Long:  "multiple automation platforms and versions with a single interface",
	}

	root.AddCommand(
		cmdEnter,
		cmdExec,
		cmdExit,
		cmdLs,
		cmdInstall,
		cmdRm,
		cmdUpdate,
	)

	root.PersistentFlags().StringVarP(&platformBase, "platform-path", "p", "/usr/lib/omni", "Directory to store platform environments")

	root.Execute()
}
