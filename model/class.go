package model

import (
	"colorme.vn/core/model"
	"github.com/jinzhu/gorm"
)

// Table name default is the pluralized version of struct
type Class struct {
	model.Model
	ID     uint   `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	BaseID uint   `json:"base_id"`

	//ignore field in gorm
	Money int `json:"money" gorm:"-"`
}

func (b Class) TableName() string {
	return "classes"
}

func ClassByGenID(genID int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("gen_id  = ? ", genID)
	}
}
