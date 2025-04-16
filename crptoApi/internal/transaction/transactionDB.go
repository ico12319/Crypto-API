package transaction

import (
	"context"
	"crptoApi/pkg/models"
	"errors"
	"fmt"
	"sync"
)

type InMemoryTransactionRepoImpl struct {
	mu           sync.Mutex
	transactions map[string]models.Transaction
}

var once sync.Once
var instance *InMemoryTransactionRepoImpl

func GetInstance() *InMemoryTransactionRepoImpl {
	once.Do(func() {
		instance = &InMemoryTransactionRepoImpl{transactions: make(map[string]models.Transaction)}
	})
	return instance
}

func (i *InMemoryTransactionRepoImpl) CreateTransaction(ctx context.Context, transaction models.Transaction) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	_, isAlreadyContained := i.transactions[transaction.ID]
	if isAlreadyContained {
		return fmt.Errorf("transaction with this %s id already exist", transaction.ID)
	}
	i.transactions[transaction.ID] = transaction
	return nil
}

func (i *InMemoryTransactionRepoImpl) GetTransaction(ctx context.Context, id string) (models.Transaction, error) {
	resChan := make(chan models.Transaction, 1)
	errChan := make(chan error, 1)
	go func() {
		i.mu.Lock()
		defer i.mu.Unlock()

		t, isContained := i.transactions[id]
		if !isContained {
			errChan <- errors.New("invalid transaction id" + id)
		} else {
			resChan <- t
		}
	}()
	select {
	case <-ctx.Done():
		return models.Transaction{}, ctx.Err()
	case err := <-errChan:
		return models.Transaction{}, err
	case res := <-resChan:
		return res, nil
	}
}

func (i *InMemoryTransactionRepoImpl) GetTransactions(ctx context.Context) (map[string]models.Transaction, error) {
	resChan := make(chan map[string]models.Transaction, 1)
	go func() {
		i.mu.Lock()
		defer i.mu.Unlock()
		resChan <- i.transactions
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case result := <-resChan:
		return result, nil
	}
}
