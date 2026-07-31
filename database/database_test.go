// Package database contains tests for the database package.
package database

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const wordList = `11111	alpha
11112	bravo
11113	charlie
`

func newTestDatabase(t *testing.T) *Database {
	t.Helper()

	dir := t.TempDir()

	old := databasePath
	databasePath = func() (string, error) {
		return filepath.Join(dir, "diceware.db"), nil
	}
	t.Cleanup(func() {
		databasePath = old
	})

	db, err := Open()
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, db.Close())
	})

	require.NoError(t, db.createSchema())

	file := filepath.Join(dir, "wordlist.txt")
	require.NoError(t, os.WriteFile(file, []byte(wordList), 0o644))

	require.NoError(t, db.importWordList(mustOpen(t, file)))

	return db
}

func mustOpen(t *testing.T, path string) *os.File {
	t.Helper()

	f, err := os.Open(path)
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, f.Close())
	})

	return f
}

func TestWord(t *testing.T) {
	db := newTestDatabase(t)

	word, err := db.Word(11112)

	require.NoError(t, err)
	assert.Equal(t, "bravo", word)
}

func TestWordNotFound(t *testing.T) {
	db := newTestDatabase(t)

	_, err := db.Word(66666)

	require.Error(t, err)
}
