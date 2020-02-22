package sqliteutils

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteAuth struct {
	DBPath string //path where .db file stored
}

//
func NewSqliteConnection(sqliteAuth SqliteAuth) *sql.DB {
	db, err := sql.Open("sqlite3", sqliteAuth.DBPath)
	if err != nil {
		log.Println(err)
		panic("Sqlite3 connection failed to open")
	}
	db.SetMaxOpenConns(1)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS hotel(name text,address text,stars int,contact text,phone text,url text)")
	if err != nil {
		log.Println(err)
		panic("Not able to create table")
	}

	return db
}
