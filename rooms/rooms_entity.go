package rooms

import (
	"github.com/ElegantSoft/shabahy/messages"
	"github.com/ElegantSoft/shabahy/users"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID       uint               `json:"id" gorm:"primary_key"`
	Hash     string             `json:"hash" binding:"required" gorm:"not null:true"`
	Users    []users.User       `json:"users" gorm:"many2many:room_users"`
	Messages []messages.Message `json:"messages"`
}

var RoomSchema = struct {
	Hash     string
	Users    string
	Messages string
}{
	Hash:     "Hash",
	Users:    "Users",
	Messages: "Messages",
}

type CreateNewRoom struct {
	Users []uint
	Hash  string
}

type ById struct {
	ID uint `uri:"id" binding:"required"`
}
