package users

import (
	"github.com/ElegantSoft/shabahy/db"
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

func (ur Repository) FindUserById(id uint) (error, *User) {
	var user User
	if result := db.DB.First(&User{}, id); result.Error == nil {
		return nil, &user
	} else {

	return result.Error, &user
	}
}

func NewRepository() *Repository {
	return &Repository{}
}