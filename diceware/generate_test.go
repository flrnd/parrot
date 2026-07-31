// Package diceware_test contains tests for the diceware package.
package diceware_test

import (
	"testing"

	"github.com/flrnd/parrot/diceware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRollDice(t *testing.T) {
	tests := []struct {
		name    string
		sides   int
		wantErr bool
	}{
		{"six-sided dice", 6, false},
		{"twenty-sided dice", 20, false},
		{"one-sided dice", 1, true},
		{"zero-sided dice", 0, true},
		{"negative-sided dice", -6, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := diceware.RollDice(tt.sides)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.GreaterOrEqual(t, got, 1)
			assert.LessOrEqual(t, got, tt.sides)
		})
	}
}

func TestRollDicewareIndex(t *testing.T) {
	got, err := diceware.RollDicewareIndex()

	require.NoError(t, err)

	assert.GreaterOrEqual(t, got, 11111)
	assert.LessOrEqual(t, got, 66666)
}

func TestGeneratePassphrase(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		delim string
		want  string
	}{
		{
			"five words",
			[]string{"one", "two", "three", "four", "five"},
			" ",
			"One Two Three Four Five",
		},
		{
			"custom delimiter",
			[]string{"one", "two"},
			"#",
			"One#Two",
		},
		{
			"empty list",
			nil,
			" ",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, diceware.GeneratePassphrase(tt.words, tt.delim))
		})
	}
}
