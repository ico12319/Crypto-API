package server

import (
	"CryptoToken/internal/repositories/accounts"
	"CryptoToken/internal/repositories/holdings"
	"CryptoToken/internal/repositories/transactions"
	"CryptoToken/internal/services/accountService"
	"CryptoToken/internal/services/coinService"
	"CryptoToken/internal/services/holdingService"
	"CryptoToken/internal/services/transactionService"
	"CryptoToken/pkg/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	aController *controllers.AccountController
	tController *controllers.TransactionController
	hController *controllers.HoldingController
}

func NewServer() *Server {
	aRepo := accounts.GetInstance()
	aServ := accountService.NewService(aRepo)
	aCont := controllers.NewAccountController(aServ)

	hRepo := holdings.GetInstance()
	hServ := holdingService.NewService(hRepo)
	hCont := controllers.NewHoldingController(hServ)

	client := &http.Client{}
	cServ := coinService.NewHttpCoinService(client)

	tRepo := transactions.GetInstance()
	tServ := transactionService.NewService(tRepo, aRepo, hRepo, cServ)
	tCont := controllers.NewTransactionController(tServ)

	return &Server{tController: tCont, aController: aCont, hController: hCont}
}

func (s *Server) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/transactions/{id}", s.tController.GetTransactions).Methods(http.MethodGet)
	router.HandleFunc("/transactions/", s.tController.CreateTransaction).Methods(http.MethodPost)
	router.HandleFunc("/transactions/", s.tController.GetTransactions).Methods(http.MethodGet)

	router.HandleFunc("/holdings/{id}", s.hController.GetHoldingRecord).Methods(http.MethodGet)
	router.HandleFunc("/holdings/", s.hController.GetHoldings).Methods(http.MethodGet)

	router.HandleFunc("/accounts/{quantity}", s.aController.SetBalance).Methods(http.MethodPost)
	router.HandleFunc("/accounts/{quantity}", s.aController.UpdateBalance).Methods(http.MethodPut)
	router.HandleFunc("/accounts/", s.aController.GetBalance).Methods(http.MethodGet)

}

func (s *Server) Start() {
	router := mux.NewRouter()
	s.RegisterRoutes(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
