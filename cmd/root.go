// Package cmd implements some basic utilies to ease building cells.
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is the parent of all commands defined in this package.
var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "Various tools to build Cells",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
