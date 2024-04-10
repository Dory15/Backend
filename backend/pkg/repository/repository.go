package repository

import "backend/pkg/models"

type IRepository interface {
	SaveUser(models.SignUpStruct) error
	GetUser(models.SignUpStruct) (models.SignUpStruct, error)
}

func NewRepository() IRepository {
	return NewMSSQLRepository()
}
