package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdUpdate = &cobra.Command{
	Use:   "update",
	Short: "Update cache of available and installed platform versions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(platformBase)
	},
}
