package main

import (
	"github.com/spf13/cobra"
)

var cmdExec = &cobra.Command{
	Use:   "exec",
	Short: "Run a command within one or more environments, avoiding the need to `omni enter|exit`",
	Run: func(cmd *cobra.Command, args []string) {
		println("exec: ")
	},
}
