package service

import (
	"errors"

	"github.com/Nur-Hidayat-FTI22E/Go_UnitTest/entity"
	"github.com/Nur-Hidayat-FTI22E/Go_UnitTest/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}
}
