package models

type Holding struct {
	Crypto   string  `json:"crypto_id"`
	Quantity float64 `json:"quantity"`
	Price    float64 `json:"price"`
}
