package main

import (
	"github.com/spf13/cobra"
)

var cmdInstall = &cobra.Command{
	Use:   "install",
	Short: "Install a new platform version",
	Run: func(cmd *cobra.Command, args []string) {
		println("install: ")
	},
}
