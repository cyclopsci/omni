package main

import (
	"github.com/cyclopsci/omni"
	"github.com/spf13/cobra"
)

var cmdExit = &cobra.Command{
	Use:   "exit",
	Short: "Exit an execution environment",
	Run: func(cmd *cobra.Command, args []string) {
		omni.Deactivate()
	},
}
