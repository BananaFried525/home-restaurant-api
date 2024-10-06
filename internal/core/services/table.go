package services

import (
	"github.com/BananaFried525/home-restaurant-api/internal/core/domain"
	"github.com/BananaFried525/home-restaurant-api/internal/core/ports"
)

type TableService struct {
	repo ports.TableRepository
}

func NewTableService(repo ports.TableRepository) ports.TableService {
	return &TableService{
		repo: repo,
	}
}

func (t *TableService) AddTable(Number int) error {
	return t.repo.CreateTable(domain.Table{Number: Number})
}
