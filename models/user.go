package models

import (
	"gin-training/database"
	"gin-training/forms/requests"
	"gin-training/forms/responses"
)

type User struct {
	BaseModel
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (userModel User) Request(req requests.UserRequest) User {
	user := User{
		BaseModel: BaseModel{
			ID:        userModel.ID,
			CreatedAt: userModel.CreatedAt,
			UpdatedAt: userModel.UpdatedAt,
			DeletedAt: userModel.DeletedAt,
		},
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	return user
}

func (userModel User) Response() responses.UserResponse {
	result := responses.UserResponse{
		BaseResponse: responses.BaseResponse{
			ID:        userModel.ID,
			CreatedAt: userModel.CreatedAt,
			UpdatedAt: userModel.UpdatedAt,
			DeletedAt: userModel.DeletedAt,
		},
		Name:  userModel.Name,
		Email: userModel.Email,
	}
	return result
}

func (userModel User) FindAll() ([]responses.UserResponse, error) {
	var users []responses.UserResponse
	result := database.DB.Model(&User{}).Find(&users)
	return users, result.Error
}

func (userModel User) FindById(userId string) (User, error) {
	var user User
	result := database.DB.Model(&User{}).First(&user, userId)
	return user, result.Error
}

func (userModel User) Create(userObj User) (User, error) {
	user := userObj
	result := database.DB.Model(&user).Create(&user)
	return user, result.Error
}

func (userModel User) Update(userObj User) (User, error) {
	user := userObj
	result := database.DB.Model(&user).Where("id = ?", user.ID).Updates(&user)
	return user, result.Error
}

func (userModel User) Delete(userObj User) (int64, error) {
	result := database.DB.Delete(&userObj, userObj.ID)
	return result.RowsAffected, result.Error
}
