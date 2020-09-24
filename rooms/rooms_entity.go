package rooms

import (
	"github.com/ElegantSoft/shabahy/messages"
	"github.com/ElegantSoft/shabahy/users"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Hash string `json:"hash" binding:"required" gorm:"not null:true"`
	Users []users.User `json:"users" gorm:"many2many:room_users"`
	Messages []messages.Message `json:"messages"`
}

type ById struct {
	ID uint `uri:"id" binding:"required"`
}

