package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var appName = "parrot"
var appVersion = "v0.0.2"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  "Print the version numner",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s %s\n", appName, appVersion)
	},
}
