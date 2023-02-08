package database

import (
	"fmt"
	"os"

	"github.com/Flexusaurs/israchan-server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// db connection singleton
	DBConn *gorm.DB
)

func InitDb() *gorm.DB {
	dsn, envExists := os.LookupEnv("DB_DSN")
	if !envExists {
		panic("Couldn't connect to DB")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("couldn't init DB")
	}

	migrate(db)
	return db
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Post{}, &models.Reply{})
	if err != nil {
		fmt.Print("couldn't init DB")
		return
	}
	fmt.Print("DB MIGRATED")
}
