package accountService

import (
	"CryptoToken/internal/repositories/accounts"
	"fmt"
)

type AccountService interface {
	GetAccountBalance() float64
	SetAccountBalance(amount float64) error
	UpdateAccountBalance(amount float64) error
}

type Service struct {
	accountRepo accounts.AccountRepository
}

func NewService(accountRepo accounts.AccountRepository) *Service {
	return &Service{accountRepo: accountRepo}
}

func (s *Service) GetAccountBalance() float64 {
	return s.accountRepo.GetBalance()
}

func (s *Service) SetAccountBalance(amount float64) error {
	if err := validateAmount(amount); err != nil {
		return err
	}
	s.accountRepo.SetBalance(amount)
	return nil
}

func (s *Service) UpdateAccountBalance(amount float64) error {
	if err := validateAmount(amount); err != nil {
		return err
	}
	s.accountRepo.UpdateBalance(amount)
	return nil
}

func validateAmount(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("negative ammount provided %s", amount)
	}
	return nil
}
