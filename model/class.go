package model

import (
	"colorme.vn/core/model"
	"github.com/jinzhu/gorm"
	"colorme.vn/core/service"
)

// Table name default is the pluralized version of struct
type Class struct {
	model.Model
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Target   int    `json:"target"`
	BaseID   uint   `json:"base_id"`
	GenID    uint   `json:"gen_id"`
	CourseID uint   `json:"course_id"`

	//ignore field in gorm
	Money int `json:"money" gorm:"-"`

	Course Course
}

func (b Class) TableName() string {
	return "classes"
}

func ClassByGenID(genID int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("gen_id  = ? ", genID)
	}
}

func GetTargetClass(class *Class) float64 {
	db := service.GetService().DB.DB
	db.Model(class).Related(&class.Course)
	return 0.55 * float64(class.Target*class.Course.Price)
}

func GetTargetClasses(classes *[]Class) float64 {
	var targetRevenue float64 = 0
	for _, class := range *classes {
		targetRevenue += GetTargetClass(&class)
	}
	return targetRevenue
}
