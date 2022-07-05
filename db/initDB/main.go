package main

import (
	"bufio"
	"database/sql"
	"os"
	"strconv"
	"strings"

	"github.com/flrnd/parrot/util"
	_ "github.com/mattn/go-sqlite3"
)

func initDB() {
	var dbPath = "./db/diceware_list.db"
	f, err := os.Open("./txt/diceware.wordlist.asc")

	util.Check(err)

	fdb, err := os.Create(dbPath)

	util.Check(err)

	defer f.Close()
	fdb.Close()

	scanner := bufio.NewScanner(f)
	wordsDatabase, err := sql.Open("sqlite3", dbPath)

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

	for scanner.Scan() {
		word := scanner.Text()
		tuple := strings.Split(word, "	")
		id, err := strconv.Atoi(tuple[0])
		util.Check(err)

		stmt.Exec(id, tuple[1])
	}
}
func main() {
	initDB()
}
