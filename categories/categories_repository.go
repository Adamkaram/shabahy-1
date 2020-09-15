package categories

import (
	"github.com/ElegantSoft/shabahy/common"
)

type Repository struct {
	crud *common.CrudRepository
}

func (r *Repository) Find(id uint) (error, interface{}) {
	var itemToFind Category
	return r.crud.Find(id, &itemToFind)
}

func (r *Repository) Create(item *Category) (error, interface{}) {
	return r.crud.Create(item)
}

func (r *Repository) Update(item *Category, id uint) error {
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
