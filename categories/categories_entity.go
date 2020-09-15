package categories

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
}

type ById struct {
	ID uint `uri:"id" binding:"required"`
}