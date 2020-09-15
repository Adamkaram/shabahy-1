package categories

import (
	"github.com/ElegantSoft/shabahy/db"
	"gorm.io/gorm"
)

type Repository struct {
}


func (r *Repository) Find(id uint) (error, *Category)  {
	var itemToFind Category
	if err := db.DB.First(&itemToFind, id); err.Error != nil {
		return err.Error, nil
	}
	return nil, &itemToFind
}

func (r *Repository) Create(item *Category) (error, *Category) {
	if createUser := db.DB.Create(&item); createUser.Error != nil {
		return createUser.Error, nil
	}
	return nil, item
}

func (r *Repository) Update(item *Category, id uint) error {
	var itemToUpdate = Category{
		Model: gorm.Model{
			ID:        id,
		},
	}
	if err := db.DB.Model(&itemToUpdate).UpdateColumns(&item); err != nil {
		return err.Error
	}
	return nil
}

func (r *Repository) Delete(id uint) error {
	if err := db.DB.Delete(&Category{}, id); err != nil {
		return err.Error
	}
	return nil
}


func NewRepository() *Repository {
	return &Repository{}
}
