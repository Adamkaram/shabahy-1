package rooms

import (
	"github.com/ElegantSoft/shabahy/common"
	"github.com/ElegantSoft/shabahy/db"
	"github.com/ElegantSoft/shabahy/users"
)

type Repository struct {
	crud *common.CrudRepository
}

func (r Repository) Paginate() (error, interface{}) {
	var result []Room
	if err, _ := r.crud.Paginate(&result); err.Error != nil {
		return err, nil
	}
	return nil, result
}

func (r *Repository) Find(id uint) (error, interface{}) {
	var itemToFind Room
	return r.crud.Find(id, &itemToFind)
}

func (r *Repository) Create(hash string, users *[]users.User) (error, *Room) {
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

func (r *Repository) Update(item *Room, id uint) error {
	return r.crud.Update(id, &item)
}

func (r *Repository) Delete(id uint) error {
	return r.crud.Delete(id)
}

func NewRepository(crud *common.CrudRepository) *Repository {
	return &Repository{
		crud: crud,
	}
}
