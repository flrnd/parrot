package main

import (
	"database/sql"
	"fmt"

	"github.com/flrnd/parrot/db"
	"github.com/flrnd/parrot/util"
)

func main() {
	dbPath := db.Path()
	database, err := sql.Open("sqlite3", dbPath)
	util.Check(err)

	length := 8
	words := make([]string, length)

	for i := 0; i < length; i++ {
		dices := util.RollDices(5, 6)
		id := util.DicesToNumber(dices)

		word := db.Find(id, database)
		words[i] = word
	}
	defer database.Close()

	passphrase := util.GeneratePassphrase(words, "-")
	fmt.Println(passphrase)
}
