package account

import (
	"crptoApi/internal/entities"
	"database/sql"
	"errors"
)

type IDatabase interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}
type SQLAccountDB struct {
	db IDatabase
}

func NewSQLAccountDB(database IDatabase) *SQLAccountDB {
	return &SQLAccountDB{db: database}
}

func (s *SQLAccountDB) createAccount() error {
	_, err := s.db.Exec("INSERT INTO accounts(balance) VALUES(?)", 0.0)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLAccountDB) GetBalance() (float64, error) {
	var entity entities.Account
	if err := s.db.Get(&entity, "SELECT * FROM accounts WHERE id=?", 1); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = s.createAccount()
			if err != nil {
				return 0.0, err
			}
		}
	}
	return entity.Balance, nil
}

func (s *SQLAccountDB) UpdateBalance(amount float64) error {
	currBalance, err := s.GetBalance()
	if err != nil {
		return err
	}
	_, err = s.db.Exec("UPDATE accounts SET balance=? WHERE id=?", currBalance+amount, 1)
	if err != nil {
		return err
	}
	return nil
}
