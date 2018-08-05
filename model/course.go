package model

import (
	"colorme.vn/core/model"
	)

// Table name default is the pluralized version of struct
type Course struct {
	model.Model
	ID     uint   `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`

	Price int `json:"money"`
}

func (b Course) TableName() string {
	return "courses"
}
