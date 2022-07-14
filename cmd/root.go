package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "parrot"}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		return
	}
}
