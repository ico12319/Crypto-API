package account

import (
	"context"
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

func (i *InMemoryAccountRepoImpl) GetBalance(ctx context.Context) (float64, error) {
	select {
	case <-ctx.Done():
		return 0.0, ctx.Err()
	default:
	}
	i.mu.Lock()
	defer instance.mu.Unlock()

	return i.accounts[0].Balance, nil
}

func (i *InMemoryAccountRepoImpl) UpdateBalance(ctx context.Context, amount float64) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	i.mu.Lock()
	defer i.mu.Unlock()

	i.accounts[0].Balance += amount
	return nil
}
