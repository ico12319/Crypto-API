package entities

type Transaction struct {
	ID         int64   `db:"id"`
	Type       string  `db:"type"`
	CryptoName string  `db:"crypto_name"`
	Quantity   float64 `db:"quantity"`
}
