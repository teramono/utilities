package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// DB ...
type DB struct {
	*sql.DB
}

// Connect ...
func Connect(connectionURI string) (DB, error) {
	db, err := sql.Open("sqlite3", connectionURI)
	if err != nil {
		return DB{}, err
	}

	return DB{
		db,
	}, nil
}
