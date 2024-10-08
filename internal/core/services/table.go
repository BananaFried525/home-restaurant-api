package services

import (
	"github.com/BananaFried525/home-restaurant-api/internal/core/domain"
	"github.com/BananaFried525/home-restaurant-api/internal/core/entities"
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

func (t *TableService) GetListTable(limit int, offset int) (*[]entities.TableInfo, error) {
	return t.repo.GetTable(limit, offset)
}

func (t *TableService) GetTableDetail(ID uint) (*entities.TableInfo, error) {
	return t.repo.GetTableByID(ID)
}

func (t *TableService) UpdateTable(ID uint, table domain.Table) error {
	return t.repo.UpdateTable(ID, table)
}

func (t *TableService) DeleteTable(ID uint) error {
	return t.repo.DeltetTable(ID)
}
