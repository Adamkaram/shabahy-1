package messages

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Name      string `json:"name" binding:"required"`
}

type ById struct {
	ID uint `uri:"id" binding:"required"`
}

