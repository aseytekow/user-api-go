package db

import (
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	url := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		logrus.Fatal("Failed to connect DB!\n\nError: ", err)
	}

	return db
}
