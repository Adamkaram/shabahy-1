package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

