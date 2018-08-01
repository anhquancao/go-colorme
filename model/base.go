package model

import "colorme.vn/core/model"

// Table name default is the pluralized version of struct
type Base struct {
	model.Model
	ID      uint   `json:"id" gorm:"primary_int"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (b Base) TableName() string {
	return "bases"
}
