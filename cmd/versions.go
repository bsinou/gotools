package cmd

import (
	"fmt"
	"os"

	version "github.com/hashicorp/go-version"
	"github.com/spf13/cobra"
)

var verCmd = &cobra.Command{
	Use:  "version",
	Long: `Versions commands`,
	Run: func(cm *cobra.Command, args []string) {
		cm.Usage()
	},
}

var ivCmd = &cobra.Command{
	Use:  "isvalid",
	Long: `Tries to parse passed version as string using the hashicorp/go-version library`,
	Run: func(cm *cobra.Command, args []string) {
		if len(args) != 1 {
			cm.Printf("Please provide a version to parse\n")
			os.Exit(1)
		}

		versionStr := args[0]
		fmt.Printf("Checking version %s \n", versionStr)

		_, err := version.NewVersion(versionStr)
		if err != nil {
			cm.Printf("Passed version [%s] is not a valid version\n", versionStr)
			os.Exit(1)
		}
	},
}

func init() {
	verCmd.AddCommand(ivCmd)
	RootCmd.AddCommand(verCmd)
}
