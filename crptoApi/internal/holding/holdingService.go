package holding

import (
	"context"
	"crptoApi/pkg/models"
)

type HoldingRepository interface {
	GetHolding(ctx context.Context, id string) (models.Holding, error)
	GetHoldings(ctx context.Context) (map[string]models.Holding, error)
}

type Service struct {
	hRepo HoldingRepository
}

func NewService(hRepo HoldingRepository) *Service {
	return &Service{hRepo: hRepo}
}

func (s *Service) GetHoldingRecord(ctx context.Context, id string) (models.Holding, error) {
	return s.hRepo.GetHolding(ctx, id)
}

func (s *Service) GetHoldingsRecords(ctx context.Context) (map[string]models.Holding, error) {
	return s.hRepo.GetHoldings(ctx)
}
