package messages

import (
	"github.com/ElegantSoft/shabahy/common"
)

type Repository struct {
	crud *common.CrudRepository
}

func (r Repository) Paginate() (error, interface{}) {
	var result []Message
	if err, _ := r.crud.Paginate(&result); err.Error != nil {
		return err, nil
	}
	return nil, result
}

func (r *Repository) Find(id uint) (error, interface{}) {
	var itemToFind Message
	return r.crud.Find(id, &itemToFind)
}

func (r *Repository) Create(item *Message) (error, interface{}) {
	return r.crud.Create(item)
}

func (r *Repository) Update(item *Message, id uint) error {
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
