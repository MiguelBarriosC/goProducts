package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

var Datab string = "./products.db"

func GetConnection() *sql.DB {

	if db != nil {
		return db
	}

	var err error
	db, err = sql.Open("sqlite3", Datab)
	if err != nil {
		panic(err)
	}
	return db
}
