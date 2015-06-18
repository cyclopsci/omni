package main

import (
	"github.com/spf13/cobra"
)

var cmdExit = &cobra.Command{
	Use:   "exit",
	Short: "Exit an execution environment",
	Run: func(cmd *cobra.Command, args []string) {
		println("Exit")
	},
}
