package database

import (
	"fmt"
	"mygram/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     string = "localhost"
	port     int    = 5432
	username string = "postgres"
	password string = "postgres"
	dbName   string = "db_mygram"
	db       *gorm.DB
	err      error
)

func StartDB() error {
	conn := fmt.Sprintf("host=%s  user=%s password=%s dbname=%s port=%d sslmode=disable", host, username, password, dbName, port)
	db, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("Successfully Connected to Database: ", dbName)
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	return nil
}

func GetDB() *gorm.DB {
	return db
}
