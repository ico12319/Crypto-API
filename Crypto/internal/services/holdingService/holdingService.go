package holdingService

import (
	"CryptoToken/internal/repositories/holdings"
	"CryptoToken/pkg/constants"
	"CryptoToken/pkg/models"
	"fmt"
)

type HoldingService interface {
	CreateHoldingRecord(holding models.Holding) error
	DeleteHoldingRecord(crypto string) error
	GetHoldingRecord(cryptoId string) (models.Holding, error)
	GetHoldingsRecords() map[string]models.Holding
	UpdateHoldingRecord(cryptoId string, quantity float64) error
}

func validateHolding(holding models.Holding) error {
	if holding.Price < 0 {
		return fmt.Errorf("negative price %s", holding.Price)
	}
	if holding.Quantity < 0 {
		return fmt.Errorf("negative quantity %.2f", holding.Quantity)
	}
	_, isSupportedToken := constants.ACCEPTED_TOKENS[holding.Crypto]
	if !isSupportedToken {
		return fmt.Errorf("unsupported token %s", holding.Crypto)
	}
	return nil
}

type Service struct {
	repo holdings.HoldingRepository
}

func (s *Service) GetHoldingsRecords() map[string]models.Holding {
	return s.repo.GetHoldings()
}

func validateToken(crypto string) error {
	_, isSupportedToken := constants.ACCEPTED_TOKENS[crypto]
	if !isSupportedToken {
		return fmt.Errorf("unsupported crypto token %s", crypto)
	}
	return nil
}

func NewService(repo holdings.HoldingRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateHoldingRecord(holding models.Holding) error {
	if err := validateHolding(holding); err != nil {
		return err
	}
	return s.repo.CreateHolding(holding)
}

func (s *Service) DeleteHoldingRecord(crypto string) error {
	if err := validateToken(crypto); err != nil {
		return err
	}
	return s.repo.DeleteHolding(crypto)
}

func (s *Service) GetHoldingRecord(crypto string) (models.Holding, error) {
	if err := validateToken(crypto); err != nil {
		return models.Holding{}, err
	}
	return s.repo.GetHolding(crypto)
}

func (s *Service) UpdateHoldingRecord(cryptoId string, quantity float64) error {
	return s.repo.UpdateHolding(cryptoId, quantity)
}
