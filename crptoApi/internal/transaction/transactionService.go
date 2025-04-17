package transaction

import (
	"context"
	"crptoApi/pkg/constants"
	"crptoApi/pkg/models"
	"fmt"
)

type AccountRepository interface {
	GetBalance() (float64, error)
	UpdateBalance(amount float64) error
}

type HoldingRepository interface {
	CreateHolding(holding models.Holding) error
	UpdateHolding(cryptoId string, quantity float64) error
	GetHolding(cryptoId string) (models.Holding, error)
	GetHoldings() ([]models.Holding, error)
	DeleteHolding(cryptoId string) error
}

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction models.Transaction) error
	GetTransaction(id string) (models.Transaction, error)
	GetTransactions() ([]models.Transaction, error)
}

type CoinService interface {
	GetCoinPrice(ctx context.Context, cryptoId string) (float64, error)
}

type Service struct {
	tRepo    TransactionRepository
	aRepo    AccountRepository
	hRepo    HoldingRepository
	cService CoinService
}

func NewService(tRepo TransactionRepository, aRepo AccountRepository, hRepo HoldingRepository, cService CoinService) *Service {
	return &Service{tRepo: tRepo, aRepo: aRepo, hRepo: hRepo, cService: cService}
}

func (s *Service) CreateTransactionRecord(ctx context.Context, transaction models.Transaction) error {
	tokenPrice, err := s.cService.GetCoinPrice(ctx, transaction.Crypto)
	if err != nil {
		return err
	}
	if transaction.Type == constants.Buy {
		if err = handleTransactionTypeBuy(s.aRepo, transaction, s.hRepo, tokenPrice); err != nil {
			return err
		}

	} else if transaction.Type == constants.Sell {
		if err = handleTransactionTypeSell(s.aRepo, transaction, s.hRepo, tokenPrice); err != nil {
			return err
		}
	}
	if err = s.tRepo.CreateTransaction(ctx, transaction); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetTransactionRecord(id string) (models.Transaction, error) {
	return s.tRepo.GetTransaction(id)
}

func (s *Service) GetTransactionsRecords() ([]models.Transaction, error) {
	return s.tRepo.GetTransactions()
}

func getBalanceHelper(aRepo AccountRepository) (float64, error) {
	return aRepo.GetBalance()
}

func checkBalance(currBalance float64, tokenPrice float64, quantity float64) error {
	if currBalance < tokenPrice*quantity {
		return fmt.Errorf("insufficient funds")
	}
	return nil
}

func createNewHoldingHelper(hRepo HoldingRepository, cryptoId string, quantity float64, price float64) error {
	return hRepo.CreateHolding(models.Holding{
		Crypto:      cryptoId,
		Quantity:    quantity,
		PriceBought: price,
	})
}

func handleTransactionTypeBuy(aRepo AccountRepository, transaction models.Transaction, hRepo HoldingRepository, tokenPrice float64) error {
	balance, err := getBalanceHelper(aRepo)
	err = checkBalance(balance, tokenPrice, transaction.Quantity)
	if err != nil {
		return err
	}
	err = aRepo.UpdateBalance(-(tokenPrice * transaction.Quantity))
	if err != nil {
		return err
	}
	err = hRepo.UpdateHolding(transaction.Crypto, transaction.Quantity)
	if err != nil {
		if err = createNewHoldingHelper(hRepo, transaction.Crypto, transaction.Quantity, tokenPrice); err != nil {
			return err
		}
	}
	return nil
}

func handleTransactionTypeSell(aRepo AccountRepository, transaction models.Transaction, hRepo HoldingRepository, tokenPrice float64) error {
	err := hRepo.UpdateHolding(transaction.Crypto, -transaction.Quantity)
	if err != nil {
		return err
	}
	h, err := hRepo.GetHolding(transaction.Crypto)
	if err != nil {
		return err
	}
	if h.Quantity < 0 {
		err = hRepo.DeleteHolding(transaction.Crypto)
		if err != nil {
			return err
		}
		return fmt.Errorf("you don't have enough quantity")
	}
	if h.Quantity == 0 {
		err = hRepo.DeleteHolding(transaction.Crypto)
		if err != nil {
			return err
		}
	}
	err = aRepo.UpdateBalance(tokenPrice * transaction.Quantity)
	if err != nil {
		return err
	}
	return nil
}
