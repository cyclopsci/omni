package main

import (
	"fmt"

	"github.com/cyclopsci/omni"
	"github.com/spf13/cobra"
)

var cmdExec = &cobra.Command{
	Use:   "exec [platform] [version] [command...]",
	Short: "Run a command within one or more environments, avoiding the need to `omni enter|exit`",
	Run: func(cmd *cobra.Command, args []string) {
		err := omni.Run(platformBase, args[0], args[1], args[2:])
		if err != nil {
			fmt.Println(err)
		}
	},
}
