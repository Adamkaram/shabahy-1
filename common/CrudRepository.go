package common

import (
	"github.com/ElegantSoft/shabahy/db"
	"time"
)

type CrudRepository struct {
	table string
}

func (r *CrudRepository) Paginate(dest interface{}) (error, interface{}) {
	if err := db.DB.Table(r.table + " as t").Find(dest); err.Error != nil {
		return err.Error, nil
	}
	return nil, dest
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

func (r *CrudRepository) Update(id uint,item interface{}) error {
	if err := db.DB.Table(r.table + " as t").Where("id", id).UpdateColumns(item); err != nil {
		return err.Error
	}
	return nil
}

func (r *CrudRepository) Delete(id uint) error {
	if err := db.DB.Table(r.table + " as t").Where("id", id).Update("deleted_at", time.Now()); err != nil {
		return err.Error
	}
	return nil
}


func NewCrudRepository(table string) *CrudRepository {
	return &CrudRepository{
		table: table,
	}
}
