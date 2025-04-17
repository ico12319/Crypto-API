package account

import (
	"crptoApi/internal/transaction"
)

type Service struct {
	aRepo transaction.AccountRepository
}

func NewService(aRepo transaction.AccountRepository) *Service {
	return &Service{aRepo: aRepo}
}

func (s *Service) GetAccountBalance() (float64, error) {
	return s.aRepo.GetBalance()
}

func (s *Service) UpdateAccountBalance(amount float64) error {
	return s.aRepo.UpdateBalance(amount)
}
