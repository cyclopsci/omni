package main

import (
	"errors"
	"fmt"
	"path"

	"github.com/cyclopsci/omni"
	"github.com/spf13/cobra"
)

var (
	ErrMissingRequiredArgs = errors.New("Missing required arguments")
	ErrInvalidPlatform     = errors.New("Invalid Platform")
	ErrInvalidVersion      = errors.New("Invalid Version")
)

var (
	stateBase    = ".omni"
	platformBase = ".omni/platforms"
	format       = "text"
	output       = ""
)

func main() {
	root := &cobra.Command{
		Use:   "omni",
		Short: "multiple automation platforms and versions with a single interface",
		Long:  "multiple automation platforms and versions with a single interface",
	}

	cmdEnter := &cobra.Command{
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

	cmdExec := &cobra.Command{
		Use:   "exec [platform] [version] [command...]",
		Short: "Run a command within one or more environments, avoiding the need to `omni enter|exit`",
		Run: func(cmd *cobra.Command, args []string) {
			err := omni.Run(platformBase, args[0], args[1], args[2:])
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	cmdExit := &cobra.Command{
		Use:   "exit",
		Short: "Exit an execution environment",
		Run: func(cmd *cobra.Command, args []string) {
			err := omni.Exit()
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	cmdInstall := &cobra.Command{
		Use:   "install",
		Short: "Install a new platform version",
		Run: func(cmd *cobra.Command, args []string) {
			err := argsInstall(args)
			if err != nil {
				fmt.Println(err)
				cmd.Usage()
				return
			}
			err = omni.InstallPlatform(platformBase, args[0], args[1])
			if err != nil {
				fmt.Println(err)
				cmd.Usage()
				return
			}
		},
	}

	cmdLs := &cobra.Command{
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

	root.AddCommand(
		cmdEnter,
		cmdExec,
		cmdExit,
		cmdLs,
		cmdInstall,
	)

	root.PersistentFlags().StringVarP(&platformBase, "platform-path", "p", ".omni/platforms", "Directory to store platform environments")
	root.PersistentFlags().StringVarP(&format, "format", "f", "text", "Output format: [text|json]")
	root.PersistentFlags().StringVarP(&output, "out-file", "o", "", "Write output to file in addition to STDOUT")

	root.Execute()
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

func argsInstall(args []string) error {
	if len(args) < 2 {
		return ErrMissingRequiredArgs
	}
	return nil
}
