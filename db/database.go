package db

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/flrnd/parrot/util"
	_ "github.com/mattn/go-sqlite3"
)

func Path() string {
	var dbPath = "/.config/parrot/diceware_list.db"
	homeDir, _ := os.UserHomeDir()

	return homeDir + dbPath
}

func Init(f *os.File) {
	homeDir, _ := os.UserHomeDir()
	databasePath := Path()

	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		os.MkdirAll(homeDir+"/.config/parrot", 0755)
		file, err := os.Create(databasePath)
		util.Check(err)

		file.Close()
	}

	// Open the database
	parrotDB, err := sql.Open("sqlite3", databasePath)
	util.Check(err)
	defer parrotDB.Close()

	scanner := bufio.NewScanner(f)
	wordsDatabase, err := sql.Open("sqlite3", databasePath)

	util.Check(err)

	defer wordsDatabase.Close()

	createWordsTable := `create table words(id integer not null primary key, word text) without rowid;`

	if createTable, err := wordsDatabase.Prepare(createWordsTable); err != nil {
		util.Check(err)
	} else {
		createTable.Exec()
	}

	stmt, err := wordsDatabase.Prepare("INSERT INTO words(id, word) values(?,?)")
	util.Check(err)

	fmt.Printf("Initializing the database with %s\n", f.Name())

	for scanner.Scan() {
		word := scanner.Text()
		tuple := strings.Split(word, "	")
		id, err := strconv.Atoi(tuple[0])
		util.Check(err)

		stmt.Exec(id, tuple[1])
	}
}

func Find(id int, db *sql.DB) string {
	rows, err := db.Query("select id, word from words where id = ?", id)
	util.Check(err)
	var resultId int
	var resultWord string
	for rows.Next() {

		err := rows.Scan(&resultId, &resultWord)
		util.Check(err)
	}
	return resultWord
}
