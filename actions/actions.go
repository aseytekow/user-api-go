package actions

import (
	"fmt"

	"github.com/aseytekow/user-api-go/db"
	"github.com/aseytekow/user-api-go/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateUser(c *gin.Context) {
	var user models.User

	c.ShouldBindJSON(&user)

	// // res := db.ConnectDB().Create(&user)

	// if res.Error != nil {
	// 	logrus.Fatal("Failed to create user!")
	// }

	query := fmt.Sprintf("INSERT INTO ACCOUNT(NAME, EMAIL, PASSWORD) VALUES('%v', '%v', '%v')", user.Name, user.Email, user.Password)

	res := db.DBConn().QueryRow(query)

	if res.Err() != nil {
		logrus.Fatal("Failed to create user!")
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func ListAllUsers(c *gin.Context) {
	var users []models.User

	res, err := db.DBConn().Query("SELECT * FROM ACCOUNT")

	if err != nil {
		logrus.Fatal("Failed to execute query!")
	}

	for res.Next() {
		var u models.User
		res.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
		users = append(users, u)
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetUser(c *gin.Context) {
	var id = c.Param("id")
	var user models.User

	query := fmt.Sprintf("SELECT * FROM ACCOUNT WHERE ID = '%v'", id)

	res := db.DBConn().QueryRow(query)

	if res.Err() != nil {
		logrus.Fatal("User not found!")
	}

	res.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func DeleteUser(c *gin.Context) {
	var id = c.Param("id")

	query := fmt.Sprintf("DELETE FROM ACCOUNT WHERE ID = '%v'", id)

	res := db.DBConn().QueryRow(query)

	if res.Err() != nil {
		logrus.Fatal("User not found!")
	}
}

func UpdateUser(c *gin.Context) {
	var id = c.Param("id")
	var user models.User

	c.ShouldBindJSON(&user)

	query := fmt.Sprintf("UPDATE ACCOUNT SET NAME = '%v', EMAIL = '%v', PASSWORD = '%v' WHERE ID = '%v'", user.Name, user.Email, user.Password, id)

	res := db.DBConn().QueryRow(query)

	if res.Err() != nil {
		logrus.Fatal("User not found!")
	}

	res.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	c.JSON(200, gin.H{
		"user": user,
	})
}
