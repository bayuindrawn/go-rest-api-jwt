package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
)

// SetupDatabaseConnection is creating a new connection to our database
func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to oad env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp($s:3306)/%s?charset=utf&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	return db

}

// Close database connection
func CoseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close database connection")
	}
	dbSQL.close()
}
