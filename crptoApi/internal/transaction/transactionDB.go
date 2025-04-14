package transaction

import (
	"crptoApi/pkg/models"
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

func (i *InMemoryTransactionRepoImpl) CreateTransaction(transaction models.Transaction) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	_, isAlreadyContained := i.transactions[transaction.ID]
	if isAlreadyContained {
		return fmt.Errorf("transaction with this %s id already exist", transaction.ID)
	}
	i.transactions[transaction.ID] = transaction
	return nil
}

func (i *InMemoryTransactionRepoImpl) GetTransaction(id string) (models.Transaction, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	t, isContained := i.transactions[id]
	if !isContained {
		return models.Transaction{}, fmt.Errorf("invalid transaction id %s", id)
	}
	return t, nil
}

func (i *InMemoryTransactionRepoImpl) GetTransactions() map[string]models.Transaction {
	i.mu.Lock()
	defer i.mu.Unlock()

	return i.transactions
}
