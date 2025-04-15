package account

import (
	"context"
	"crptoApi/internal/transaction"
)

type Service struct {
	aRepo transaction.AccountRepository
}

func NewService(aRepo transaction.AccountRepository) *Service {
	return &Service{aRepo: aRepo}
}

func (s *Service) GetAccountBalance(ctx context.Context) (float64, error) {
	return s.aRepo.GetBalance(ctx)
}

func (s *Service) UpdateAccountBalance(ctx context.Context, amount float64) error {
	return s.aRepo.UpdateBalance(ctx, amount)
}
