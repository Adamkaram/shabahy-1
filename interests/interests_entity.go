package interests

import (
	"gorm.io/gorm"
)

type Interest struct {
	gorm.Model
	Name string `json:"name" binding:"required" gorm:"not null:true"`
	CategoryID uint `json:"category_id" gorm:"not null:true"`
}

type ById struct {
	ID uint `uri:"id" binding:"required"`
}
type InterestAPI struct {
	ID uint `uri:"id" binding:"required"`
	Name string `json:"name" binding:"required" gorm:"not null:true"`
	CategoryID uint `json:"category_id" gorm:"not null:true"`
}

func (InterestAPI) TableName() string {
	return "interests"
}