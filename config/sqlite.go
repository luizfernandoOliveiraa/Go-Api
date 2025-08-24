package config

import (
	"os"

	"github.com/luizfernandoOliveiraa/Go-Api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// checks if the database dfile exists
	_, err := os.Stat(dbPath)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Info("Database file not found, creating...")
			// Create the database file
			err = os.MkdirAll("./db", os.ModePerm)
			if err != nil {
				logger.Errorf("Error creating database directory: %v", err)
				return nil, err
			}
			file, err := os.Create(dbPath)
			if err != nil {
				logger.Errorf("Error creating database file: %v", err)
				return nil, err
			}
			file.Close()
		}
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Sqlite opening error: %v", err)
		return nil, err
	}
	// Migrate the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Sqlite auto migrate error: %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}
