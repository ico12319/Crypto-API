package cache

import (
	"sync"
	"time"
)

type cachePair struct {
	priceBought float64
	timeBought  time.Time
}

type Cache struct {
	mu          sync.Mutex
	cachedPrice map[string]cachePair
	duration    time.Duration
}

var once sync.Once
var instance *Cache

func GetInstance() *Cache {
	once.Do(func() {
		instance = &Cache{cachedPrice: make(map[string]cachePair), duration: 2 * time.Minute}
	})
	return instance
}

func (c *Cache) AddToCache(cryptoId string, price float64) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, exist := c.cachedPrice[cryptoId]
	if exist {
		return false
	}
	c.cachedPrice[cryptoId] = cachePair{
		priceBought: price,
		timeBought:  time.Now(),
	}
	return true
}

func (c *Cache) GetPrice(cryptoId string) (float64, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cPair, exist := c.cachedPrice[cryptoId]
	if !exist {
		return 0.0, false
	}
	if time.Since(cPair.timeBought) > c.duration {
		delete(c.cachedPrice, cryptoId)
		return 0.0, false
	}
	return cPair.priceBought, true
}
