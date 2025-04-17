package entities

type Account struct {
	Id      int64   `db:"id"`
	Balance float64 `db:"balance"`
}
