package common

import (
	"github.com/ElegantSoft/shabahy/db"
)

type CrudRepository struct {
	Model interface{}
}


func (r *CrudRepository) Find(id uint, dest interface{}) (error, interface{})  {
	if err := db.DB.First(dest, id); err.Error != nil {
		return err.Error, nil
	}
	return nil, dest
}

func (r *CrudRepository) Create(item interface{}) (error, interface{}) {
	if createUser := db.DB.Create(item); createUser.Error != nil {
		return createUser.Error, nil
	}
	return nil, item
}

func (r *CrudRepository) Update(dest interface{},item interface{}) error {
	if err := db.DB.Model(dest).UpdateColumns(item); err != nil {
		return err.Error
	}
	return nil
}

func (r *CrudRepository) Delete(id uint) error {
	if err := db.DB.Delete(&r.Model, id); err != nil {
		return err.Error
	}
	return nil
}


func NewCrudRepository(model interface{}) *CrudRepository {
	return &CrudRepository{
		Model: model,
	}
}
