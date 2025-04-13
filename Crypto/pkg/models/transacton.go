package models

import "CryptoToken/pkg/constants"

type Transaction struct {
	ID       string                    `json:"id"`
	Type     constants.TransactionType `json:"type"`
	Crypto   string                    `json:"crypto_id"`
	Quantity float64                   `json:"crypto_quantity"`
}
