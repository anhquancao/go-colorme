package model

import "colorme.vn/core/model"

// Table name default is the pluralized version of struct
type Base struct {
	model.Model
	ID      uint   `json:"id" gorm:"primary_key"`
	ClassID uint   `json:"class_id"`
	Name    string `json:"name"`
	Address string `json:"address"`

	//ignore field in gorm
	Money int `json:"money" gorm:"-"`

	//related

	Classes []Class `gorm:"foreignkey:BaseID"`
}

func (b Base) TableName() string {
	return "bases"
}
