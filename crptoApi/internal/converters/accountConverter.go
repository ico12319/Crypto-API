package converters

import (
	"crptoApi/internal/entities"
	"crptoApi/pkg/models"
)

func ConvertFromModelToEntityAccount(account models.Account) entities.Account {
	return entities.Account{
		Balance: account.Balance,
	}
}

func ConvertFromEntityToModelAccount(account entities.Account) models.Account {
	return models.Account{
		Balance: account.Balance,
	}
}
