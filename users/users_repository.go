package users

import (
	"github.com/ElegantSoft/shabahy/db"
	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
}

func (r *Repository) Create(user *User) (error, *User) {
	createUser := db.DB.Create(user)
	if createUser.Error != nil {
		return createUser.Error, nil
	}
	return nil, user
}

func (r *Repository) FindUserByIdAndPassword(data *LoginUserDTO) (error, *User) {
	var user User
	findCondition := &FindByEmail{
		Email: data.Email,
	}
	result := db.DB.First(&user, findCondition)
	if result.Error == nil {
		errorValidate := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
		if errorValidate == nil {
			return nil, &user
		}
		return errorValidate, nil
	}
	return result.Error, nil
}

func (r *Repository) GetUserData(id uint) (error, *User) {
	var user User
	findCondition := &User{
		ID: id,
	}
	if err := db.DB.Find(&user, findCondition); err.Error != nil {
		return err.Error, nil
	}
	return nil, &user

}

func NewRepository() *Repository {
	return &Repository{}
}
