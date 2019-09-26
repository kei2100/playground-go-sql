package repository

import (
	"database/sql"
	"os"
	"testing"
)

var db *sql.DB

func TestMain(t *testing.M) {
	status := run(t)
	os.Exit(status)
}

func run(t *testing.M) int {
	disposeDB := setupDB()
	defer disposeDB()

	cleanUp := setupCleaner()
	defer cleanUp()

	return t.Run()
}

func setupDB() func() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=25432 user=postgres password=postgres dbname=develop sslmode=disable")
	if err != nil {
		panic(err)
	}
	return func() {
		db.Close()
	}
}
