package main

import (
	"crptoApi/internal/account"
	"crptoApi/internal/cache"
	"crptoApi/internal/coin"
	"crptoApi/internal/holding"
	"crptoApi/internal/server"
	"crptoApi/internal/transaction"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

func main() {
	transactionDb, err := sqlx.Connect("sqlite3", "transactions.db")
	if err != nil {
		panic("failed to connect")
	}
	defer transactionDb.Close()

	holdingDb, err := sqlx.Connect("sqlite3", "holdings.db")
	if err != nil {
		panic("failed to connect")
	}
	defer holdingDb.Close()

	accountDb, err := sqlx.Connect("sqlite3", "users.db")
	if err != nil {
		panic("failed to connect")
	}
	defer accountDb.Close()

	hRepo := holding.NewSQLHoldingDB(holdingDb)
	hService := holding.NewService(hRepo)
	hHandler := holding.NewHoldingHandler(hService)

	aRepo := account.NewSQLAccountDB(accountDb)
	aService := account.NewService(aRepo)
	aHandler := account.NewAccountHandler(aService)

	client := &http.Client{}

	cService := coin.NewHttpCoinService(client)

	tRepo := transaction.NewSQLTransactionDB(transactionDb)
	c := cache.GetInstance()
	tService := transaction.NewService(tRepo, aRepo, hRepo, cService, c)
	tHandler := transaction.NewTransactionHandler(tService)

	serv := server.NewServer(aHandler, hHandler, tHandler)
	serv.Start()

}
