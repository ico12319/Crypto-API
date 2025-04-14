package account

import (
	"crptoApi/pkg/models"
	"sync"
)

type InMemoryAccountRepoImpl struct {
	mu       sync.Mutex
	accounts []models.Account
}

var once sync.Once
var instance *InMemoryAccountRepoImpl

func GetInstance() *InMemoryAccountRepoImpl {
	once.Do(func() {
		instance = &InMemoryAccountRepoImpl{accounts: make([]models.Account, 1, 8)}
	})
	return instance
}

func (i *InMemoryAccountRepoImpl) GetBalance() float64 {
	i.mu.Lock()
	defer instance.mu.Unlock()

	return i.accounts[0].Balance
}

func (i *InMemoryAccountRepoImpl) UpdateBalance(amount float64) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.accounts[0].Balance += amount
	return nil
}
