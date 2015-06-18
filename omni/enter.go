package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

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

		println("enter: ", args[0], args[1])
	},
}

func argsEnter(args []string) error {
	if len(args) < 2 {
		return ErrMissingRequiredArgs
	}

	return nil
}
