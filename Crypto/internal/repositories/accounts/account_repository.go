package accounts

import (
	"CryptoToken/pkg/models"
	"sync"
)

type AccountRepository interface {
	GetBalance() float64
	SetBalance(staringAmount float64)
	UpdateBalance(amount float64)
}

type InMemoryAccountRepository struct {
	mu       sync.Mutex
	accounts []models.Account
}

var once sync.Once
var instance *InMemoryAccountRepository

func GetInstance() *InMemoryAccountRepository {
	once.Do(func() {
		instance = &InMemoryAccountRepository{accounts: make([]models.Account, 1)}
	})
	return instance
}

func (i *InMemoryAccountRepository) GetBalance() float64 {
	i.mu.Lock()
	defer i.mu.Unlock()

	return i.accounts[0].Balance
}

func (i *InMemoryAccountRepository) SetBalance(amount float64) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.accounts[0].Balance = amount
}

func (i *InMemoryAccountRepository) UpdateBalance(amount float64) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.accounts[0].Balance += amount
}
