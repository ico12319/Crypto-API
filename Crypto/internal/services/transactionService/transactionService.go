package transactionService

import (
	"CryptoToken/internal/repositories/accounts"
	"CryptoToken/internal/repositories/holdings"
	"CryptoToken/internal/repositories/transactions"
	"CryptoToken/internal/services/coinService"
	"CryptoToken/pkg/constants"
	"CryptoToken/pkg/models"
	"fmt"
)

type TransactionService interface {
	CreateTransactionRecord(transaction models.Transaction) error
	GetTransactionRecord(id string) (models.Transaction, error)
	GetTransactions() map[string]models.Transaction
}

type Service struct {
	transRepo   transactions.TransactionRepository
	accountRepo accounts.AccountRepository
	holdingRepo holdings.HoldingRepository
	coinService coinService.CoinService
}

func NewService(transRepo transactions.TransactionRepository, accountRepo accounts.AccountRepository, holdingRepo holdings.HoldingRepository, coinService coinService.CoinService) *Service {
	return &Service{transRepo: transRepo, accountRepo: accountRepo, holdingRepo: holdingRepo, coinService: coinService}
}

func (s *Service) CreateTransactionRecord(transaction models.Transaction) error {
	if err := validateTransaction(transaction); err != nil {
		return err
	}
	if err := s.transRepo.CreateTransaction(transaction); err != nil {
		return err
	}
	price, err := callApiHelper(s.coinService, transaction.Crypto)
	if err != nil {
		return err
	}
	if transaction.Type == constants.Buy {
		err := handleTransactionBuyHelper(s.accountRepo, s.holdingRepo, transaction, price)
		if err != nil {
			return err
		}
	} else if transaction.Type == constants.Sell {
		err := handleTransactionSellHelper(s.accountRepo, s.holdingRepo, transaction, price)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) GetTransactionRecord(id string) (models.Transaction, error) {
	trans, err := s.transRepo.GetTransaction(id)
	if err != nil {
		return models.Transaction{}, err
	}
	return trans, nil
}

func (s *Service) GetTransactions() map[string]models.Transaction {
	return s.transRepo.GetTransactions()
}

func callApiHelper(service coinService.CoinService, cryptoId string) (float64, error) {
	return service.GetCryptoPrice(cryptoId)
}

func createHoldingHelper(repo holdings.HoldingRepository, cryptoId string, quantity float64, price float64) error {
	return repo.CreateHolding(models.Holding{
		Crypto:   cryptoId,
		Quantity: quantity,
		Price:    price,
	})
}

func handleTransactionBuyHelper(aRepo accounts.AccountRepository, hRepo holdings.HoldingRepository, transaction models.Transaction, price float64) error {
	aRepo.UpdateBalance(-(price * transaction.Quantity))
	err := hRepo.UpdateHolding(transaction.Crypto, transaction.Quantity)
	if err != nil {
		err := createHoldingHelper(hRepo, transaction.Crypto, transaction.Quantity, price)
		if err != nil {
			return err
		}
	}
	return nil
}

func handleTransactionSellHelper(aRepo accounts.AccountRepository, hRepo holdings.HoldingRepository, transaction models.Transaction, price float64) error {
	aRepo.UpdateBalance(price * transaction.Quantity)
	err := hRepo.UpdateHolding(transaction.Crypto, transaction.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func validateTransaction(transaction models.Transaction) error {
	if transaction.Type != constants.Buy && transaction.Type != constants.Sell {
		return fmt.Errorf("unsupported operation %s", transaction.Type)
	}
	if transaction.Quantity < 0 {
		return fmt.Errorf("negative quantity %.2f", transaction.Quantity)
	}
	_, isSupportedToken := constants.ACCEPTED_TOKENS[transaction.Crypto]
	if !isSupportedToken {
		return fmt.Errorf("unsupported crypto token %s", transaction.Crypto)
	}
	return nil
}
