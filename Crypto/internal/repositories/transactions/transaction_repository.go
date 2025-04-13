package transactions

import (
	"CryptoToken/pkg/models"
	"fmt"
	"sync"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.Transaction) error
	GetTransaction(id string) (models.Transaction, error)
	GetTransactions() map[string]models.Transaction
}

type InMemoryTransactionDatabase struct {
	mu           sync.Mutex
	transactions map[string]models.Transaction
}

var once sync.Once
var instance *InMemoryTransactionDatabase

func GetInstance() *InMemoryTransactionDatabase {
	once.Do(func() {
		instance = &InMemoryTransactionDatabase{transactions: make(map[string]models.Transaction)}
	})
	return instance
}

func (i *InMemoryTransactionDatabase) CreateTransaction(transaction models.Transaction) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	_, isAlreadyContained := i.transactions[transaction.ID]
	if isAlreadyContained {
		return fmt.Errorf("there is already a transaction with this id")
	}
	i.transactions[transaction.ID] = transaction
	return nil
}

func (i *InMemoryTransactionDatabase) GetTransaction(id string) (models.Transaction, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	tx, isAlreadyContained := i.transactions[id]
	if !isAlreadyContained {
		return models.Transaction{}, fmt.Errorf("invalid transaction id")
	}
	return tx, nil
}

func (i *InMemoryTransactionDatabase) GetTransactions() map[string]models.Transaction {
	i.mu.Lock()
	defer i.mu.Unlock()
	return i.transactions
}
