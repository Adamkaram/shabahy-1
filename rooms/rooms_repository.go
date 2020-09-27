package rooms

import (
	"github.com/ElegantSoft/shabahy/common"
	"github.com/ElegantSoft/shabahy/db"
	"log"
)

type Repository struct {
	crud *common.CrudRepository
}

func (r Repository) paginate() (error, interface{}) {
	var result []Room
	if err, _ := r.crud.Paginate(&result); err.Error != nil {
		return err, nil
	}
	return nil, result
}

func (r *Repository) find(id uint) (error, interface{}) {
	var itemToFind Room
	return r.crud.Find(id, &itemToFind)
}

func (r *Repository) create(hash string, users *[]User) (error, *Room) {
	room := &Room{
		Hash: hash,
	}
	if err := db.DB.Create(&room); err.Error != nil {
		return err.Error, nil
	}
	err := db.DB.Model(&room).Association(RoomSchema.Users).Append(users)
	if err != nil {
		return err, nil
	}
	return nil, room
}

func (r *Repository) appendMessage(room *Room,message *Message) error {
	return db.DB.Model(&room).Association(RoomSchema.Messages).Append(message)
}


func (r *Repository) findRoomWithUsersIds(ids []uint) (error, bool) {

	type RoomUsers struct {
		RoomID string `json:"room_id"`
	}
	var result []RoomUsers

	err := db.DB.Raw(`SELECT room_id
							FROM room_users rm1
							WHERE user_id IN ?
							GROUP BY room_id
							HAVING COUNT(user_id) = ?
							AND NOT EXISTS
							  (SELECT *
							   FROM room_users rm2
							   WHERE rm2.room_id = rm1.room_id
								 AND user_id NOT IN ?)`, ids, len(ids), ids).Scan(&result)
	if err.Error != nil {
		log.Println("err2", err.Error)
		return err.Error, false
	}
	log.Println(result)
	return nil, len(result) > 0
}

func (r *Repository) update(item *Room, id uint) error {
	return r.crud.Update(id, &item)
}

func (r *Repository) delete(id uint) error {
	return r.crud.Delete(id)
}

func NewRepository(crud *common.CrudRepository) *Repository {
	return &Repository{
		crud: crud,
	}
}
