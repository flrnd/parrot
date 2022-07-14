package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/flrnd/parrot/db"
	"github.com/flrnd/parrot/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate [length] [delimiter]",
	Short: "Generate a secure passphrase",
	Long:  "Generate a secure 8 words long passphrase (~200bits entropy).\nExample: parrot 5 '-'\n",
	Args:  cobra.RangeArgs(0, 2),
	Run: func(cmd *cobra.Command, args []string) {
		length, delimiter := parseArgs(args)
		dbPath := db.Path()
		database, err := sql.Open("sqlite3", dbPath)
		util.Check(err)

		words := make([]string, length)

		for i := 0; i < length; i++ {
			dices := util.RollDices(5, 6)
			id := util.DicesToNumber(dices)

			word := db.Find(id, database)
			words[i] = word
		}
		defer database.Close()

		passphrase := util.GeneratePassphrase(words, delimiter)
		fmt.Println(passphrase)
	},
}

func parseArgs(args []string) (l int, d string) {
	switch len(args) {
	case 1:
		length, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return length, " "
	case 2:
		length, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return length, args[1]
	}
	return 8, " "
}
