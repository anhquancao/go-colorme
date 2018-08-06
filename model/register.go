package model

import (
	"colorme.vn/core/model"
	"github.com/jinzhu/gorm"
)

// Table name default is the pluralized version of struct
type Register struct {
	model.Model
	ID     uint   `json:"id" gorm:"primary_key"`
	Status string `json:"status"`
	Money  string `json:"money"`
}

func (b Register) TableName() string {
	return "registers"
}

func PaidMoney(db *gorm.DB) *gorm.DB {
	return db.Where("money > ?", 0)
}

func RegisterNew(db *gorm.DB) *gorm.DB {
	return db.Where("status = 0 OR money > 0")
}

func RegisterByClassesID(classesID []uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("class_id in (?)", classesID)
	}
}
