package users

import "gorm.io/gorm"

type Gender string

const (
	MALE Gender = "male"
	FEMALE Gender = "female"
)


type User struct {
	gorm.Model
	Name string `json:"name" binding:"required" gorm:"not null:true"`
	Phone string `json:"phone" binding:"required"`
	Email string `json:"email" binding:"required,email" gorm:"not null:true"`
	Password string `json:"password" binding:"required,min=8" gorm:"not null:true"`
	Gender string `json:"gender" binding:"Enum=male_female" gorm:"type:gender;not null:true;default:male"`
}


type LoginUserDTO struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type FindByEmail struct {
	Email string
}