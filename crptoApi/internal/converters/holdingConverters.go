package converters

import (
	"crptoApi/internal/entities"
	"crptoApi/pkg/models"
)

func ConvertFromModelToEntityHolding(holding models.Holding) entities.Holding {
	return entities.Holding{
		CryptoId:    holding.Crypto,
		Quantity:    holding.Quantity,
		PriceBought: holding.PriceBought,
	}
}

func ConvertFromEntityToModelHolding(holding entities.Holding) models.Holding {
	return models.Holding{
		Crypto:      holding.CryptoId,
		Quantity:    holding.Quantity,
		PriceBought: holding.PriceBought,
	}
}
