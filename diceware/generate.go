// Package diceware provides functions for generating Diceware passphrases.
package diceware

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	rolls = 5 // Diceware uses five dice rolls per word.
	sides = 6 // Diceware uses six-sided dice.
)

// RollDice returns a cryptographically secure dice roll in the range [1, sides].
func RollDice(sides int) (int, error) {
	if sides < 2 {
		return 0, fmt.Errorf("dice must have at least 2 sides")
	}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(sides)))
	if err != nil {
		return 0, err
	}

	return int(n.Int64()) + 1, nil
}

// RollDicewareIndex returns a Diceware index generated from five six-sided dice rolls.
func RollDicewareIndex() (int, error) {
	index := 0

	for range rolls {
		roll, err := RollDice(sides)
		if err != nil {
			return 0, err
		}

		index = index*10 + roll
	}

	return index, nil
}

// GeneratePassphrase joins the given words into a Diceware passphrase.
func GeneratePassphrase(words []string, delim string) string {
	var sb strings.Builder
	c := cases.Title(language.Und)

	for i, word := range words {
		sb.WriteString(c.String(word))
		if i < len(words)-1 {
			sb.WriteString(delim)
		}
	}

	return sb.String()
}
