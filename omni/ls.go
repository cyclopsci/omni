package main

import (
	"fmt"

	"github.com/cyclopsci/omni"
	"github.com/spf13/cobra"
)

var cmdLs = &cobra.Command{
	Use:   "ls",
	Short: "List installed and available platform versions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(platformBase)

		platforms := omni.GetPlatforms()
		for _, p := range platforms {
			println(p.Label)
		}
	},
}
