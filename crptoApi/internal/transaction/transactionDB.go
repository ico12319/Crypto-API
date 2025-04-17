package transaction

import (
	"context"
	"crptoApi/internal/converters"
	"crptoApi/internal/entities"
	"crptoApi/pkg/models"
	"database/sql"
)

type IDatabase interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

type SQLTransactionDB struct {
	db IDatabase
}

func NewSQLTransactionDB(DB IDatabase) *SQLTransactionDB {
	return &SQLTransactionDB{db: DB}
}

func (s *SQLTransactionDB) CreateTransaction(ctx context.Context, transaction models.Transaction) error {
	convertedDbEntity := converters.ConvertFromModelToEntity(transaction)
	res, err := s.db.Exec("INSERT INTO transactions(type,crypto_name,quantity) VALUES(?,?,?)", convertedDbEntity.Type, convertedDbEntity.CryptoName, convertedDbEntity.Quantity)
	if err != nil {
		return err
	}
	convertedDbEntity.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLTransactionDB) GetTransaction(id string) (models.Transaction, error) {
	var entity entities.Transaction
	if err := s.db.Get(&entity, "SELECT id,type,crypto_name,quantity FROM transactions WHERE id=?", id); err != nil {
		return models.Transaction{}, err
	}
	return converters.ConvertFromEntityToModel(entity), nil
}

func (s *SQLTransactionDB) GetTransactions() ([]models.Transaction, error) {
	var transactionEntities []entities.Transaction
	if err := s.db.Select(&transactionEntities, "SELECT * FROM transactions"); err != nil {
		return nil, err
	}
	result := make([]models.Transaction, len(transactionEntities))
	for index, transactionEntity := range transactionEntities {
		convertedModel := converters.ConvertFromEntityToModel(transactionEntity)
		result[index] = convertedModel
	}
	return result, nil
}
