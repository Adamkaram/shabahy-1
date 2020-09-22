package users

import (
	"github.com/ElegantSoft/shabahy/db"
	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
}

func (ur *Repository) Create(user *User) (error, *User) {
	createUser := db.DB.Create(user)
	if createUser.Error != nil {
		return createUser.Error, nil
	}
	return nil, user
}

func (ur Repository) FindUserByIdAndPassword(data *LoginUserDTO) (error, *User) {
	var user User
	findCondition := &FindByEmail{
		Email: data.Email,
	}
	result := db.DB.Select("email", "id", "password").First(&user, findCondition)
	if result.Error == nil {
		errorValidate := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
		if errorValidate == nil {
			return nil, &user
		}
		return errorValidate, nil
	}
	return result.Error, nil

}

func NewRepository() *Repository {
	return &Repository{}
}
