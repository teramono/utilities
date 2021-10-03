package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func ConnectDB(connectionURI string) (DB, error) {
	db, err := gorm.Open(sqlite.Open(connectionURI), &gorm.Config{})
	if err != nil {
		return DB{}, err
	}

	return DB{
		db,
	}, nil
}
