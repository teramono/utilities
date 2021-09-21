package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB ...
type DB struct {
	*gorm.DB
}

// Connect ...
func Connect(connectionURI string) (DB, error) {
	db, err := gorm.Open(sqlite.Open(connectionURI), &gorm.Config{})
	if err != nil {
		return DB{}, err
	}

	return DB{
		db,
	}, nil
}
