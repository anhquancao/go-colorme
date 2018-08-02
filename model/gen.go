package model

import (
	"colorme.vn/core/model"
	"colorme.vn/core/service"
	"time"
)

// Table name default is the pluralized version of struct
type Gen struct {
	model.Model
	ID     uint   `gorm:"primary_key"`
	Name   string `json:"name"`
	Status string `json:"status"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
}

func (b Gen) TableName() string {
	return "gens"
}

func GetCurrentGen() Gen {
	db := service.GetService().DB.DB

	var gen Gen

	db.Where("status = ?", 1).First(&gen)

	return gen
}
