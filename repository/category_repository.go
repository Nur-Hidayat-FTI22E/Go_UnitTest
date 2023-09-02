package repository

import "github.com/Nur-Hidayat-FTI22E/Go_UnitTest/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
