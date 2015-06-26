package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cyclopsci/omni"
	"github.com/spf13/cobra"
)

var (
	ErrMissingRequiredArgs = errors.New("Missing required arguments")
)

var (
	stateBase    = "$HOME/.omni"
	platformBase = "$HOME/.omni/platforms"
	format       = "text"
	output       = ""
	execPlatform = ""
	execVersion  = "latest"
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
			homeBase := expandPath(platformBase)
			if len(args) < 2 {
				fmt.Println(ErrMissingRequiredArgs)
				cmd.Usage()
				return
			}
			if err := omni.Enter(homeBase, args[0], args[1]); err != nil {
				fmt.Println(err)
			}
		},
	}

	cmdExec := &cobra.Command{
		Use:   "exec [platform] [version] [command...]",
		Short: "Run a command within one or more environments, avoiding the need to `omni enter|exit`",
		Run: func(cmd *cobra.Command, args []string) {
			homeBase := expandPath(platformBase)
			if len(args) < 1 {
				fmt.Println(ErrMissingRequiredArgs)
				cmd.Usage()
				return
			}
			opts := omni.ExecOptions{
				Format: format,
				Output: output,
			}
			versions := strings.Split(execVersion, ",")
			if len(versions) > 1 {
				if err := omni.ExecMultiple(homeBase, execPlatform, versions, args, opts); err != nil {
					fmt.Println(err)
				}
			} else {
				if err := omni.Exec(homeBase, execPlatform, versions[0], args, opts); err != nil {
					fmt.Println(err)
				}
			}
		},
	}
	cmdExec.Flags().StringVarP(&execPlatform, "platform", "p", "", "Platform to exec against")
	cmdExec.Flags().StringVarP(&execVersion, "version", "v", "latest", "Comma separated list of versions to exec against or one of: all|latest")

	cmdExit := &cobra.Command{
		Use:   "exit",
		Short: "Exit an execution environment",
		Run: func(cmd *cobra.Command, args []string) {
			if err := omni.Exit(); err != nil {
				fmt.Println(err)
			}
		},
	}

	cmdInstall := &cobra.Command{
		Use:   "install",
		Short: "Install a new platform version",
		Run: func(cmd *cobra.Command, args []string) {
			homeBase := expandPath(platformBase)
			if len(args) < 2 {
				fmt.Println(ErrMissingRequiredArgs)
				cmd.Usage()
				return
			}
			if err := omni.InstallPlatform(homeBase, args[0], args[1]); err != nil {
				fmt.Println(err)
			}
		},
	}

	cmdLs := &cobra.Command{
		Use:   "ls",
		Short: "List installed and available platform versions",
		Run: func(cmd *cobra.Command, args []string) {
			homeBase := expandPath(platformBase)
			platforms, _ := omni.GetPlatforms(homeBase)
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

	root.PersistentFlags().StringVarP(&platformBase, "platform-dir", "d", "~/.omni/platforms", "Directory to store platform environments")
	root.PersistentFlags().StringVarP(&format, "format", "f", "text", "Output format: [text|json]")
	root.PersistentFlags().StringVarP(&output, "out", "o", "", "Write output to file in addition to STDOUT")

	root.Execute()
}

func expandPath(path string) string {
	return os.ExpandEnv(strings.Replace(path, "~", "$HOME", -1))
}
