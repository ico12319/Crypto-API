package main

import (
	"crptoApi/internal/account"
	"crptoApi/internal/coin"
	"crptoApi/internal/holding"
	"crptoApi/internal/server"
	"crptoApi/internal/transaction"
	"net/http"
)

func main() {
	hRepo := holding.GetInstance()
	hService := holding.NewService(hRepo)
	hHandler := holding.NewHoldingHandler(hService)

	aRepo := account.GetInstance()
	aService := account.NewService(aRepo)
	aHandler := account.NewAccountHandler(aService)

	client := &http.Client{}

	cService := coin.NewHttpCoinService(client)

	tRepo := transaction.GetInstance()
	tService := transaction.NewService(tRepo, aRepo, hRepo, cService)
	tHandler := transaction.NewTransactionHandler(tService)

	serv := server.NewServer(aHandler, hHandler, tHandler)
	serv.Start()
}
