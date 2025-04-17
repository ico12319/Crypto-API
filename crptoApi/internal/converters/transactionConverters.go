package converters

import (
	"crptoApi/internal/entities"
	"crptoApi/pkg/constants"
	"crptoApi/pkg/models"
)

func ConvertFromEntityToModel(transaction entities.Transaction) models.Transaction {
	return models.Transaction{
		Type:     constants.TransactionType(transaction.Type),
		Crypto:   transaction.CryptoName,
		Quantity: transaction.Quantity,
	}
}

func ConvertFromModelToEntity(transaction models.Transaction) entities.Transaction {
	return entities.Transaction{
		Type:       string(transaction.Type),
		CryptoName: transaction.Crypto,
		Quantity:   transaction.Quantity,
	}
}
