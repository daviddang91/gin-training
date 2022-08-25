package controllers

import (
	"gin-training/database"
	"gin-training/models"

	"github.com/gin-gonic/gin"
)

func GetUsersController(c *gin.Context) {
	users := []models.User{}
	database.DB.Find(&users)
	c.JSON(200, &users)
}

func DetailUserController(c *gin.Context) {
	var user models.User
	database.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	c.JSON(200, &user)
}

func CreateUserController(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	database.DB.Create(&user)
	c.JSON(200, &user)
}

func UpdateUserController(c *gin.Context) {
	var user models.User
	database.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	database.DB.Save(&user)
	c.JSON(200, &user)
}

func DeleteUserController(c *gin.Context) {
	var user models.User
	database.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}
