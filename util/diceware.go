package util

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/flrnd/gorng"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// RollDices returns a slice with n results, where n is the number rolls and d is the dice number of sides.
func RollDices(n int, d int) []int {
	result := make([]int, n)

	for i := 0; i < n; i++ {
		num, err := gorng.GenerateRandomInt(int64(d))
		Check(err)
		result[i] = num + 1
	}

	return result
}

// DicesToNumber returns the number value from the dices slice.
func DicesToNumber(dices []int) int {
	result, err := strconv.Atoi(strings.Trim(strings.Replace(fmt.Sprint(dices), " ", "", -1), "[]"))
	Check(err)

	return result
}

func GeneratePassphrase(wordlist []string, delim string) string {
	var sb strings.Builder
	c := cases.Title(language.Und)
	for i, w := range wordlist {
		sb.WriteString(c.String(w))
		if i < len(w) {
			sb.WriteString(delim)
		}
	}
	return sb.String()
}
