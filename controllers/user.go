package controllers

import (
	"errors"
	"gin-training/models"
	"gin-training/serializers/requests"
	"gin-training/serializers/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetUsersController(c *gin.Context) {
	var userModel models.User
	users, err := userModel.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, &users)
}

func DetailUserController(c *gin.Context) {
	userId := c.Param("id")
	var userModel models.User
	user, err := userModel.FindById(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusNotFound),
			"message": err.Error(),
		})
		return
	}

	result := user.Response()
	c.JSON(200, &result)
}

func CreateUserController(c *gin.Context) {
	req := requests.UserRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]responses.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = responses.ErrorMsg{Field: fe.Field(), Message: fe.Error()}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	var userModel models.User
	userObj := userModel.Request(req)
	user, err := userModel.Create(userObj)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}

	result := user.Response()
	c.JSON(201, &result)
}

func UpdateUserController(c *gin.Context) {
	req := requests.UserRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]responses.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = responses.ErrorMsg{Field: fe.Field(), Message: fe.Error()}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	userId := c.Param("id")
	var userModel models.User
	userObj, errObj := userModel.FindById(userId)
	if errObj != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusNotFound),
			"message": errObj.Error(),
		})
		return
	}

	userReq := userObj.Request(req)
	user, err := userModel.Update(userReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, &user)
}

func DeleteUserController(c *gin.Context) {
	userId := c.Param("id")
	var userModel models.User

	user, errVal := userModel.FindById(userId)
	if errVal != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusNotFound),
			"message": errVal.Error(),
		})
		return
	}

	_, errDel := userModel.Delete(user)
	if errDel != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": errDel.Error(),
		})
		return
	}

	c.Status(204)
}
