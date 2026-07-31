// Package database manages Parrot's local Diceware SQLite database.
package database

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// Database provides access to Parrot's local Diceware database.
type Database struct {
	db *sql.DB
}

// databasePath returns the path to Parrot's local Diceware database.
var databasePath = func() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".config", "parrot", "diceware_list.db"), nil
}

// Open opens Parrot's local Diceware database.
func Open() (*Database, error) {
	path, err := databasePath()
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

// Init initializes Parrot's local Diceware database from a Diceware wordlist.
func Init(wordList string) error {
	d, err := Open()
	if err != nil {
		return err
	}
	defer d.Close()

	if err := d.createSchema(); err != nil {
		return err
	}

	f, err := os.Open(wordList)
	if err != nil {
		return err
	}
	defer f.Close()

	return d.importWordList(f)
}

// Close closes the database connection.
func (d *Database) Close() error {
	return d.db.Close()
}

// Word returns the word associated with the given Diceware index.
func (d *Database) Word(id int) (string, error) {
	var word string

	err := d.db.QueryRow(
		"SELECT word FROM words WHERE id = ?",
		id,
	).Scan(&word)
	if err != nil {
		return "", err
	}

	return word, nil
}

// createSchema creates the database schema.
func (d *Database) createSchema() error {
	_, err := d.db.Exec(createWordsTable)
	return err
}

// importWordList imports a Diceware wordlist into the database.
func (d *Database) importWordList(f *os.File) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO words(id, word) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	fmt.Printf("Initializing database from %s\n", f.Name())

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "\t", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid wordlist entry: %q", scanner.Text())
		}

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			return err
		}

		if _, err := stmt.Exec(id, parts[1]); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return tx.Commit()
}
