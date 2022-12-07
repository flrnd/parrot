package cmd

import (
	"os"

	"github.com/flrnd/parrot/db"
	"github.com/flrnd/parrot/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init my_diceware_word_list.asc",
	Short: "Initialize the diceware databse",
	Long:  "Initialize the diceware database. Download a file list from https://theworld.com/~reinhold/diceware.html or create your own.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open(args[0])

		util.Check(err)

		db.Init(f)
	},
}
