package db

import (
	"database/sql"
	"os"

	"github.com/flrnd/parrot/util"
	_ "github.com/mattn/go-sqlite3"
)

func Path() string {
	var dbPath = "/.config/parrot/diceware_list.db"
	homeDir, _ := os.UserHomeDir()

	return homeDir + dbPath
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
