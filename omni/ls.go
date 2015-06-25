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
		platforms, _ := omni.GetPlatforms(platformBase)
		for _, p := range platforms {
			line := p.Label + ":"
			for _, v := range p.Versions {
				line += " " + v.Label
			}
			fmt.Println(line)
		}
	},
}
