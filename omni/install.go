package main

import (
	"fmt"

	"github.com/cyclopsci/omni"
	"github.com/spf13/cobra"
)

var cmdInstall = &cobra.Command{
	Use:   "install",
	Short: "Install a new platform version",
	Run: func(cmd *cobra.Command, args []string) {
		err := argsInstall(args)
		if err != nil {
			fmt.Println(err)
			cmd.Usage()
			return
		}
		//TODO: validate p, v
		omni.InstallPlatform(platformBase, args[0], args[1])
	},
}

func argsInstall(args []string) error {
	if len(args) < 2 {
		return ErrMissingRequiredArgs
	}
	return nil
}
