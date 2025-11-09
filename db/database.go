package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func DBConn() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))

	if err != nil {
		logrus.Fatalf("Failed! %v", err)
	}
	return db
}
