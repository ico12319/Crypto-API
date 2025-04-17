package holding

import (
	"crptoApi/pkg/models"
)

type HoldingRepository interface {
	GetHolding(id string) (models.Holding, error)
	GetHoldings() ([]models.Holding, error)
}

type Service struct {
	hRepo HoldingRepository
}

func NewService(hRepo HoldingRepository) *Service {
	return &Service{hRepo: hRepo}
}

func (s *Service) GetHoldingRecord(id string) (models.Holding, error) {
	return s.hRepo.GetHolding(id)
}

func (s *Service) GetHoldingsRecords() ([]models.Holding, error) {
	return s.hRepo.GetHoldings()
}
