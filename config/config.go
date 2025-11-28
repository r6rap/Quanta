package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// dsn, mean data source name. It is a string that contains the connection information for the database.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	// open a connection to postgresql
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		// logger to take notes activiy query
		Logger: logger.Default.LogMode(logger.Info),
		// prepare statement to improve performance
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}


	log.Println("Connected to database")
}