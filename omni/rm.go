package main

import (
	"github.com/spf13/cobra"
)

var cmdRm = &cobra.Command{
	Use:   "rm",
	Short: "Remove a platform version",
	Run: func(cmd *cobra.Command, args []string) {
		println("rm: ")
	},
}
