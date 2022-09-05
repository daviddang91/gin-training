package services

import (
	"gin-training/database"
	"gin-training/models"
)

func IsDuplicateEmail(email string) bool {
	var count int64
	database.Instance.Model(&models.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}
