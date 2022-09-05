package models

import (
	"gin-training/serializers/responses"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (user *User) BindResponse() responses.UserResponse {
	result := responses.UserResponse{
		BaseResponse: responses.BaseResponse{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
		Name:  user.Name,
		Email: user.Email,
	}
	return result
}

func (user *User) BindAuthResponse(token string) responses.AuthResponse {
	result := responses.AuthResponse{
		BaseResponse: responses.BaseResponse{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
	return result
}

// func (userModel *User) FindAll() ([]responses.UserResponse, error) {
// 	var users []responses.UserResponse
// 	result := database.Instance.Model(&User{}).Find(&users)
// 	return users, result.Error
// }

// func (userModel *User) FindById(userId string) (User, error) {
// 	var user User
// 	result := database.Instance.Model(&User{}).First(&user, userId)
// 	return user, result.Error
// }

// func (userModel *User) Create(userObj User) (User, error) {
// 	user := userObj
// 	result := database.Instance.Model(&user).Create(&user)
// 	return user, result.Error
// }

// func (userModel *User) Update(userObj User) (User, error) {
// 	user := userObj
// 	result := database.Instance.Model(&user).Where("id = ?", user.ID).Updates(&user)
// 	return user, result.Error
// }

// func (userModel *User) Delete(userObj User) (int64, error) {
// 	result := database.Instance.Delete(&userObj, userObj.ID)
// 	return result.RowsAffected, result.Error
// }
