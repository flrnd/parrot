package cmd

import (
	"github.com/flrnd/parrot/database"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init <diceware-wordlist>",
	Short: "Initialize the Diceware database",
	Long: `Initialize the local Diceware database.

The wordlist should follow the standard Diceware format, with each line
containing a five-digit dice index and its corresponding word.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return database.Init(args[0])
	},
}
