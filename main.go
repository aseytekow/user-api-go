package main

import (
	"os"

	"github.com/aseytekow/user-api-go/actions"
	"github.com/aseytekow/user-api-go/db"

	// "github.com/aseytekow/user-api-go/db"

	// "github.com/aseytekow/user-api-go/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Failed to load .env file!")
	}
}

func main() {
	port := os.Getenv("PORT")

	server := gin.Default()

	db := db.DBConn()

	res := db.QueryRow("CREATE TABLE IF NOT EXISTS ACCOUNT(ID SERIAL PRIMARY KEY, NAME VARCHAR(50), EMAIL VARCHAR(50), PASSWORD VARCHAR(64))")

	if res.Err() != nil {
		logrus.Fatal("Failed to create ACCOUNT table!")
	}

	server.POST("/api/accounts", actions.CreateUser)
	server.GET("/api/accounts", actions.ListAllUsers)
	server.GET("/api/accounts/:id", actions.GetUser)
	server.DELETE("/api/accounts/:id", actions.DeleteUser)
	server.PUT("/api/accounts/:id", actions.UpdateUser)

	server.Run(port)
}
