package messages

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Text string `json:"text" binding:"required" gorm:"not null:true"`
	UserID uint `json:"user_id" binding:"required" gorm:"not null:true"`
	RoomID uint `json:"room_id" binding:"required" gorm:"not null:true"`
}

type ById struct {
	ID uint `uri:"id" binding:"required"`
}
