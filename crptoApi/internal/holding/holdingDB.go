package holding

import (
	"crptoApi/internal/converters"
	"crptoApi/internal/entities"
	"crptoApi/pkg/models"
	"database/sql"
	"fmt"
)

type IDatabase interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

type SQLHoldingDB struct {
	db IDatabase
}

func NewSQLHoldingDB(db IDatabase) *SQLHoldingDB {
	return &SQLHoldingDB{db: db}
}

func (s *SQLHoldingDB) CreateHolding(holding models.Holding) error {
	convertedEntity := converters.ConvertFromModelToEntityHolding(holding)
	_, err := s.db.Exec("INSERT INTO holdings(crypto_id,quantity,price_bought) VALUES(?,?,?)", convertedEntity.CryptoId, convertedEntity.Quantity, convertedEntity.PriceBought)
	if err != nil {
		return nil
	}
	return nil
}

func (s *SQLHoldingDB) GetHoldings() ([]models.Holding, error) {
	var entityHoldings []entities.Holding
	if err := s.db.Select(&entityHoldings, "SELECT * FROM holdings"); err != nil {
		return nil, err
	}
	modelsHoldings := make([]models.Holding, len(entityHoldings))
	for index, entity := range entityHoldings {
		convertedToModel := converters.ConvertFromEntityToModelHolding(entity)
		modelsHoldings[index] = convertedToModel
	}
	return modelsHoldings, nil
}

func (s *SQLHoldingDB) GetHolding(cryptoId string) (models.Holding, error) {
	var entity entities.Holding
	if err := s.db.Get(&entity, "SELECT * FROM holdings WHERE crypto_id=?", cryptoId); err != nil {
		return models.Holding{}, err
	}
	return converters.ConvertFromEntityToModelHolding(entity), nil
}

func (s *SQLHoldingDB) DeleteHolding(cryptoId string) error {
	res, err := s.db.Exec("DELETE FROM holdings WHERE crypto_id=?", cryptoId)
	if err != nil {
		return err
	}
	rowsAffectedCount, _ := res.RowsAffected()
	if rowsAffectedCount == 0 {
		return fmt.Errorf("invalid id provided")
	}
	return nil
}

func (s *SQLHoldingDB) UpdateHolding(cryptoId string, quantity float64) error {
	var entity entities.Holding
	if err := s.db.Get(&entity, "SELECT * FROM holdings WHERE crypto_id=?", cryptoId); err != nil {
		return fmt.Errorf("invalid crypto id provided %s", cryptoId)
	}
	desiredQuantity := entity.Quantity + quantity
	res, err := s.db.Exec("UPDATE holdings SET quantity=? WHERE crypto_id =?", desiredQuantity, cryptoId)
	if err != nil {
		return err
	}
	rowsAffectedCount, _ := res.RowsAffected()
	if rowsAffectedCount == 0 {
		return fmt.Errorf("error when trying to update token %s", cryptoId)
	}
	return nil
}
