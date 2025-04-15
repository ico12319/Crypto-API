package transaction

import (
	"context"
	"crptoApi/pkg/constants"
	"crptoApi/pkg/models"
	"fmt"
)

type AccountRepository interface {
	GetBalance(ctx context.Context) (float64, error)
	UpdateBalance(ctx context.Context, amount float64) error
}

type HoldingRepository interface {
	CreateHolding(ctx context.Context, holding models.Holding) error
	UpdateHolding(ctx context.Context, cryptoId string, quantity float64) error
	GetHolding(ctx context.Context, cryptoId string) (models.Holding, error)
	GetHoldings(ctx context.Context) (map[string]models.Holding, error)
	DeleteHolding(ctx context.Context, cryptoId string) error
}

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction models.Transaction) error
	GetTransaction(ctx context.Context, id string) (models.Transaction, error)
	GetTransactions(ctx context.Context) (map[string]models.Transaction, error)
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
		if err = handleTransactionTypeBuy(ctx, s.aRepo, transaction, s.hRepo, tokenPrice); err != nil {
			return err
		}

	} else if transaction.Type == constants.Sell {
		if err = handleTransactionTypeSell(ctx, s.aRepo, transaction, s.hRepo, tokenPrice); err != nil {
			return err
		}
	}
	if err = s.tRepo.CreateTransaction(ctx, transaction); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetTransactionRecord(ctx context.Context, id string) (models.Transaction, error) {
	return s.tRepo.GetTransaction(ctx, id)
}

func (s *Service) GetTransactionsRecords(ctx context.Context) (map[string]models.Transaction, error) {
	return s.tRepo.GetTransactions(ctx)
}

func getBalanceHelper(ctx context.Context, aRepo AccountRepository) (float64, error) {
	return aRepo.GetBalance(ctx)
}

func checkBalance(currBalance float64, tokenPrice float64, quantity float64) error {
	if currBalance < tokenPrice*quantity {
		return fmt.Errorf("insufficient funds")
	}
	return nil
}

func createNewHoldingHelper(ctx context.Context, hRepo HoldingRepository, cryptoId string, quantity float64, price float64) error {
	return hRepo.CreateHolding(ctx, models.Holding{
		Crypto:      cryptoId,
		Quantity:    quantity,
		PriceBought: price,
	})
}

func handleTransactionTypeBuy(ctx context.Context, aRepo AccountRepository, transaction models.Transaction, hRepo HoldingRepository, tokenPrice float64) error {
	balance, err := getBalanceHelper(ctx, aRepo)
	err = checkBalance(balance, tokenPrice, transaction.Quantity)
	if err != nil {
		return err
	}
	err = aRepo.UpdateBalance(ctx, -(tokenPrice * transaction.Quantity))
	if err != nil {
		return err
	}
	err = hRepo.UpdateHolding(ctx, transaction.Crypto, transaction.Quantity)
	if err != nil {
		if err = createNewHoldingHelper(ctx, hRepo, transaction.Crypto, transaction.Quantity, tokenPrice); err != nil {
			return err
		}
	}
	return nil
}

func handleTransactionTypeSell(ctx context.Context, aRepo AccountRepository, transaction models.Transaction, hRepo HoldingRepository, tokenPrice float64) error {
	err := hRepo.UpdateHolding(ctx, transaction.Crypto, -transaction.Quantity)
	if err != nil {
		return err
	}
	h, err := hRepo.GetHolding(ctx, transaction.Crypto)
	if err != nil {
		return err
	}
	if h.Quantity <= 0 {
		err = hRepo.DeleteHolding(ctx, transaction.Crypto)
		if err != nil {
			return err
		}
	}
	err = aRepo.UpdateBalance(ctx, tokenPrice*transaction.Quantity)
	if err != nil {
		return err
	}
	return nil
}
