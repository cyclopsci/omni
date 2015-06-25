package main

import (
	"fmt"
	"path"

	"github.com/cyclopsci/omni"
	"github.com/spf13/cobra"
)

//globals platformBase

var cmdEnter = &cobra.Command{
	Use:   "enter [platform] [version]",
	Short: "Print commands to enter an execution environment",
	Run: func(cmd *cobra.Command, args []string) {
		err := argsEnter(args)
		if err != nil {
			fmt.Println(err)
			cmd.Usage()
			return
		}
		activateVersion(args[0], args[1])
	},
}

func argsEnter(args []string) error {
	if len(args) < 2 {
		return ErrMissingRequiredArgs
	}
	p := args[0]
	v := args[1]

	platforms, _ := omni.GetPlatforms(platformBase)
	for _, platform := range platforms {
		if platform.Label == p {
			for _, version := range platform.Versions {
				if version.Label == v {
					return nil
				}
			}
			return ErrInvalidVersion
		}
	}

	return ErrInvalidPlatform
}

func activateVersion(platform string, version string) {
	absPath := path.Join(platformBase, platform, version)
	switch platform {
	case "puppet":
		omni.EnterRuby(absPath)
	case "ansible":
		omni.EnterPython(absPath)
	}
}
