package models

import "crptoApi/pkg/constants"

type Transaction struct {
	Type     constants.TransactionType `json:"type"`
	Crypto   string                    `json:"crypto_id"`
	Quantity float64                   `json:"quantity"`
}
