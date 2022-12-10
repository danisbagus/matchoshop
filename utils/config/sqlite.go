package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitSqlite() {
	var err error
	file := GetDBFile()

	db, err = sql.Open("sqlite3", file)
	if err != nil {
		panic(err)
	}
}

func GetSqliteDB() *sql.DB {
	if db == nil {
		log.Fatal("please init sqlite db")
	}

	return db
}
