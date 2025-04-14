package models

type Holding struct {
	Crypto      string  `json:"crypto_id"`
	Quantity    float64 `json:"quantity"`
	PriceBought float64 `json:"price_bought"`
}
