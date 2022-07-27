package db

import (
	"errors"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PrepareDB() (*gorm.DB, error) {
	dsn := os.Getenv("ECHO_POSTGRES_DSN")
	if dsn == "" {
		return nil, errors.New("ECHO_POSTGRES_DSN is not set")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&Message{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Close(db *gorm.DB) error {
	// TODO: implement me.
	return nil
}
