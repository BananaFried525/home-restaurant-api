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

func (t *TableService) GetListTable(limit int, offset int) ([]domain.Table, error) {
	var result []domain.Table

	tables, err := t.repo.GetTable(limit, offset)
	if err != nil {
		return result, err
	}

	result = make([]domain.Table, 0)
	for _, table := range *tables {
		tmp := domain.Table{
			ID:     table.ID,
			Number: table.Number,
			Status: string(table.Status),
		}

		result = append(result, tmp)
	}

	return result, nil
}

func (t *TableService) GetTableDetail(ID uint) (*entities.TableInfo, error) {
	return t.repo.GetTableByID(ID)
}
