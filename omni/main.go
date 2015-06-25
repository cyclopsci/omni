package main

import (
	"github.com/spf13/cobra"
)

var (
	stateBase    = ".omni"
	platformBase = ".omni/platforms"
	format       = "text"
	output       = ""
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

	root.PersistentFlags().StringVarP(&platformBase, "platform-path", "p", ".omni/platforms", "Directory to store platform environments")
	root.PersistentFlags().StringVarP(&format, "format", "f", "text", "Output format: [text|json]")
	root.PersistentFlags().StringVarP(&output, "out-file", "o", "", "Write output to file in addition to STDOUT")

	root.Execute()
}
