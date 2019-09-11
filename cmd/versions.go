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

var igtCmd = &cobra.Command{
	Use: "isgreater",
	Long: `Tries to parse passed versions as string using the hashicorp/go-version library and returns an error if:
	- one of the 2 strings cannot be parsed
	- the first version is not strictly greater than the second`,
	Run: func(cm *cobra.Command, args []string) {
		if len(args) != 2 {
			cm.Printf("Please provide two versions to be compared\n")
			os.Exit(1)
		}

		v1Str := args[0]
		v2Str := args[1]
		fmt.Printf("Comparing versions %s & %s \n", v1Str, v2Str)

		v1, err := version.NewVersion(v1Str)
		if err != nil {
			cm.Printf("Passed version [%s] is not a valid version\n", v1Str)
			os.Exit(1)
		}
		v2, err := version.NewVersion(v2Str)
		if err != nil {
			cm.Printf("Passed version [%s] is not a valid version\n", v2Str)
			os.Exit(1)
		}
		if !v1.GreaterThan(v2) {
			cm.Printf("Passed version [%s] is *not* greater than [%s]\n", v1Str, v2Str)
			os.Exit(1)
		}
	},
}

func init() {
	verCmd.AddCommand(ivCmd)
	verCmd.AddCommand(igtCmd)
	RootCmd.AddCommand(verCmd)
}
