package holdings

import (
	"CryptoToken/pkg/models"
	"fmt"
	"sync"
)

type HoldingRepository interface {
	CreateHolding(holding models.Holding) error
	DeleteHolding(cryptoId string) error
	GetHolding(cryptoId string) (models.Holding, error)
	GetHoldings() map[string]models.Holding
	UpdateHolding(cryptoId string, quantity float64) error
}

type InMemoryHoldingDatabase struct {
	mu       sync.Mutex
	holdings map[string]models.Holding
}

var once sync.Once
var instance *InMemoryHoldingDatabase

func GetInstance() *InMemoryHoldingDatabase {
	once.Do(func() {
		instance = &InMemoryHoldingDatabase{holdings: make(map[string]models.Holding)}
	})
	return instance
}

func (i *InMemoryHoldingDatabase) CreateHolding(holding models.Holding) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	_, isAlreadyContained := i.holdings[holding.Crypto]
	if isAlreadyContained {
		return fmt.Errorf("cryto %s already exists", holding.Crypto)
	}
	i.holdings[holding.Crypto] = holding
	return nil
}

func (i *InMemoryHoldingDatabase) DeleteHolding(cryptoId string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	hd, isContained := i.holdings[cryptoId]
	if !isContained {
		return fmt.Errorf("cryto %s is not present", cryptoId)
	}
	delete(i.holdings, hd.Crypto)
	return nil
}

func (i *InMemoryHoldingDatabase) GetHolding(cryptoId string) (models.Holding, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	hd, isContained := i.holdings[cryptoId]
	if !isContained {
		return models.Holding{}, fmt.Errorf("crypto %s is not present", cryptoId)
	}
	return hd, nil
}

func (i *InMemoryHoldingDatabase) GetHoldings() map[string]models.Holding {
	i.mu.Lock()
	defer i.mu.Unlock()
	return i.holdings
}

func (i *InMemoryHoldingDatabase) UpdateHolding(cryptoId string, quantity float64) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	holding, isContained := i.holdings[cryptoId]
	if !isContained {
		return fmt.Errorf("invalid crypto id %s\n", holding.Crypto)
	}
	holding.Quantity += quantity
	i.holdings[cryptoId] = holding
	return nil
}
