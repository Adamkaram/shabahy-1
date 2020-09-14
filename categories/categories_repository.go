package categories

import (
	"github.com/ElegantSoft/shabahy/db"
)

type Repository struct {
}

type model Category

func (r *Repository) Create(item *model) (error, *model) {
	createUser := db.DB.Create(&item)
	if createUser.Error != nil {
		return createUser.Error, nil
	}

	return nil, item
}

func (r Repository) Update(item *model) error {
	if err := db.DB.Model(&model{}).UpdateColumns(&item); err != nil {
		return err.Error
	}
	return nil
}

func (r Repository) Delete(item *model) error {
	if err := db.DB.Delete(model{}, &item); err != nil {
		return err.Error
	}
	return nil
}


func NewRepository() *Repository {
	return &Repository{}
}
