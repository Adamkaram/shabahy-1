package categories

import (
	"github.com/ElegantSoft/shabahy/interests"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name      string `json:"name" binding:"required"`
	Interests []interests.Interest  `json:"interests" gorm:"foreignkey:CategoryID;reference:id"`
}

type ById struct {
	ID uint `uri:"id" binding:"required"`
}

type CategoryAPI struct {
	ID        uint                 `json:"id"`
	Name      string               `json:"name"`
	Interests []interests.InterestAPI `json:"interests" gorm:"foreignkey:CategoryID;reference:id"`
}