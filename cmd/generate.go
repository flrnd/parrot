package cmd

import (
	"fmt"
	"strconv"

	"github.com/flrnd/parrot/database"
	"github.com/flrnd/parrot/diceware"
	"github.com/spf13/cobra"
)

const (
	defaultLength    = 8
	defaultDelimiter = " "
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate [length] [delimiter]",
	Short: "Generate a secure passphrase",
	Long:  `Generate a secure Diceware passphrase.\nExample: parrot generate 5 "-"`,
	Args:  cobra.RangeArgs(0, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		length, delimiter, err := parseArgs(args)
		if err != nil {
			return err
		}

		db, err := database.Open()
		if err != nil {
			return err
		}
		defer db.Close()

		words := make([]string, length)

		for i := range words {
			id, err := diceware.RollDicewareIndex()
			if err != nil {
				return err
			}

			words[i], err = db.Word(id)
			if err != nil {
				return err
			}
		}

		fmt.Println(diceware.GeneratePassphrase(words, delimiter))
		return nil
	},
}

func parseArgs(args []string) (int, string, error) {
	length := defaultLength
	delimiter := defaultDelimiter

	if len(args) > 0 {
		n, err := strconv.Atoi(args[0])
		if err != nil {
			return 0, "", err
		}
		length = n
	}

	if len(args) > 1 {
		delimiter = args[1]
	}

	return length, delimiter, nil
}
