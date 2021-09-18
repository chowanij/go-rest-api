package database

import (
	"fmt"

	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
)

// NewDatabaseConnection - return new database connection object
func NewDatabaseConnection() (*gorm.DB, error) {
	fmt.Println("Setting new database connection")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	connectString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword,
	)

	db, err := gorm.Open("postgres", connectString)

	if err != nil {
		return db, err
	}

	if err = db.DB().Ping(); err != nil {
		return db, err
	}

	return db, nil
}
