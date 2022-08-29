package models

import (
	"gin-training/database"
	"gin-training/serializers/requests"
	"gin-training/serializers/responses"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (userModel *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	userModel.Password = string(bytes)
	return nil
}

func (userModel *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (userModel *User) Request(req requests.UserRequest) User {
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

func (userModel *User) Response() responses.UserResponse {
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

func (userModel *User) FindAll() ([]responses.UserResponse, error) {
	var users []responses.UserResponse
	result := database.Instance.Model(&User{}).Find(&users)
	return users, result.Error
}

func (userModel *User) FindById(userId string) (User, error) {
	var user User
	result := database.Instance.Model(&User{}).First(&user, userId)
	return user, result.Error
}

func (userModel *User) Create(userObj User) (User, error) {
	user := userObj
	result := database.Instance.Model(&user).Create(&user)
	return user, result.Error
}

func (userModel *User) Update(userObj User) (User, error) {
	user := userObj
	result := database.Instance.Model(&user).Where("id = ?", user.ID).Updates(&user)
	return user, result.Error
}

func (userModel *User) Delete(userObj User) (int64, error) {
	result := database.Instance.Delete(&userObj, userObj.ID)
	return result.RowsAffected, result.Error
}
