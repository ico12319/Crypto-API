package main

import (
	"crptoApi/internal/account"
	"crptoApi/internal/cache"
	"crptoApi/internal/coin"
	"crptoApi/internal/holding"
	"crptoApi/internal/server"
	"crptoApi/internal/transaction"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
)

// this function will be executed before the main so there is no need to call it in the main
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	driverName := os.Getenv("DRIVER_NAME")
	transactionDb, err := sqlx.Connect(driverName, "transactions.db")
	if err != nil {
		panic("failed to connect")
	}
	defer transactionDb.Close()

	holdingDb, err := sqlx.Connect(driverName, "holdings.db")
	if err != nil {
		panic("failed to connect")
	}
	defer holdingDb.Close()

	accountDb, err := sqlx.Connect(driverName, "users.db")
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
