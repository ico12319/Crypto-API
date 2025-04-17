package entities

type Holding struct {
	CryptoId    string  `db:"crypto_id"`
	Quantity    float64 `db:"quantity"`
	PriceBought float64 `db:"price_bought"`
}
