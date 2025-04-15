package holding

import (
	"context"
	"crptoApi/pkg/models"
	"fmt"
	"sync"
)

type InMemoryHoldingReoImpl struct {
	mu       sync.Mutex
	holdings map[string]models.Holding
}

var once sync.Once
var instance *InMemoryHoldingReoImpl

func GetInstance() *InMemoryHoldingReoImpl {
	once.Do(func() {
		instance = &InMemoryHoldingReoImpl{holdings: make(map[string]models.Holding)}
	})
	return instance
}

func (i *InMemoryHoldingReoImpl) CreateHolding(ctx context.Context, holding models.Holding) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	i.mu.Lock()
	defer i.mu.Unlock()

	_, isAlreadyContained := i.holdings[holding.Crypto]
	if isAlreadyContained {
		return fmt.Errorf("holding with %s crypto_id already existing", holding.Crypto)
	}
	i.holdings[holding.Crypto] = holding
	return nil
}

func (i *InMemoryHoldingReoImpl) DeleteHolding(ctx context.Context, cryptoId string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	i.mu.Lock()
	defer i.mu.Unlock()

	_, isContained := i.holdings[cryptoId]
	if !isContained {
		return fmt.Errorf("non-existing crypto_id %s", cryptoId)
	}
	delete(i.holdings, cryptoId)
	return nil
}

func (i *InMemoryHoldingReoImpl) UpdateHolding(ctx context.Context, cryptoId string, quantity float64) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	i.mu.Lock()
	defer i.mu.Unlock()

	h, isContained := i.holdings[cryptoId]
	if !isContained {
		return fmt.Errorf("non-existing crypto_id %s", cryptoId)
	}

	h.Quantity += quantity
	i.holdings[cryptoId] = h
	return nil
}

func (i *InMemoryHoldingReoImpl) GetHolding(ctx context.Context, cryptoId string) (models.Holding, error) {
	select {
	case <-ctx.Done():
		return models.Holding{}, ctx.Err()
	default:
	}
	i.mu.Lock()
	defer i.mu.Unlock()

	h, isContained := i.holdings[cryptoId]
	if !isContained {
		return models.Holding{}, fmt.Errorf("invalid crypto_id %s", cryptoId)
	}
	return h, nil
}

func (i *InMemoryHoldingReoImpl) GetHoldings(ctx context.Context) (map[string]models.Holding, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	i.mu.Lock()
	defer i.mu.Unlock()
	return i.holdings, nil
}
