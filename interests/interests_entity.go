package interests

import (
	"gorm.io/gorm"
)

type Interest struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	CategoryID uint
}

type ById struct {
	ID uint `uri:"id" binding:"required"`
}
type InterestAPI struct {
	ID uint `uri:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	CategoryID uint
}

func (InterestAPI) TableName() string {
	return "interests"
}