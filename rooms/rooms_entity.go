package rooms

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID       uint      `json:"id" gorm:"primary_key"`
	Hash     string    `json:"hash" binding:"required" gorm:"not null:true"`
	Users    []User    `json:"users" gorm:"many2many:room_users"`
	Messages []Message `json:"messages"`
}

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" binding:"required" gorm:"not null:true"`
	Phone    string `json:"phone"`
	Email    string `json:"email" binding:"required,email" gorm:"not null:true"`
	Password string `json:"password" binding:"required,min=8" gorm:"not null:true"`
	Gender   string `json:"gender" binding:"Enum=male_female" gorm:"type:gender;not null:true;default:male"`
}

type Message struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primary_key"`
	Text   string `json:"text" binding:"required" gorm:"not null:true"`
	UserID uint   `json:"user_id" gorm:"not null:true"`
	RoomID uint   `json:"room_id" gorm:"not null:true"`
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
