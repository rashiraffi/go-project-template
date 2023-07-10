package database

import (
	"log"
	"template/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	// Connect to the PostgreSQL database
	db, err := gorm.Open(postgres.Open("host=localhost port=5433 user=<username> password=<password> dbname=<dbName> sslmode=disable"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&entities.User{})
	return db
}
