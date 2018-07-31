package service

import (
	"colorme.vn/core/database"
)

type Service struct {
	DB *database.DatabaseFacade
}

func NewService() *Service {
	return &Service{
		DB:      database.NewDatabase(),
	}
}