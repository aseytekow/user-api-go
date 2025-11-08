package actions

import (
	"github.com/aseytekow/user-api-go/db"
	"github.com/aseytekow/user-api-go/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateUser(c *gin.Context) {
	var user models.User

	c.ShouldBindJSON(&user)

	res := db.ConnectDB().Create(&user)

	if res.Error != nil {
		logrus.Fatal("Failed to create user!")
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func ListAllUsers(c *gin.Context) {
	var users []models.User
	db.ConnectDB().Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetUser(c *gin.Context) {
	var id = c.Param("id")
	var user models.User

	db.ConnectDB().Find(&user, id)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func DeleteUser(c *gin.Context) {
	var id = c.Param("id")
	var user models.User

	db.ConnectDB().Delete(&user, id)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UpdateUser(c *gin.Context) {
	var id = c.Param("id")
	var user models.User

	c.ShouldBindJSON(&user)
	res := db.ConnectDB().Where("id = ?", id).Updates(user)

	if res.Error != nil {
		logrus.Fatal("Failed to update user!")
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}
