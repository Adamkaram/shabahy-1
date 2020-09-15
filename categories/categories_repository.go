package categories

import (
	"github.com/ElegantSoft/shabahy/common"
	"github.com/ElegantSoft/shabahy/db"
	"gorm.io/gorm"
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
	var itemToUpdate = &Category{
		Model: gorm.Model{
			ID: id,
		},
	}
	return r.crud.Update(&itemToUpdate, &item)
}

func (r *Repository) Delete(id uint) error {
	if err := db.DB.Delete(&Category{}, id); err != nil {
		return err.Error
	}
	return nil
}

func NewRepository(crud *common.CrudRepository) *Repository {
	return &Repository{
		crud: crud,
	}
}
