package dbrepo

import (
	"database/sql"
)

type DBService struct {
	Writer
	SqliteDBConnection *sql.DB
}
