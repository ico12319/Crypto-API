package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ITransactionHandler interface {
	CreateTransactionHandler(w http.ResponseWriter, r *http.Request)
	GetTransactionRecordHandler(w http.ResponseWriter, r *http.Request)
	GetTransactionsHandler(w http.ResponseWriter, r *http.Request)
}

type IAccountHandler interface {
	GetBalanceHandler(w http.ResponseWriter, r *http.Request)
	UpdateBalanceHandler(w http.ResponseWriter, r *http.Request)
}

type IHoldingHandler interface {
	GetHoldingHandler(w http.ResponseWriter, r *http.Request)
	GetHoldingsHandler(w http.ResponseWriter, r *http.Request)
}

type Server struct {
	aHandler IAccountHandler
	hHandler IHoldingHandler
	tHandler ITransactionHandler
}

func NewServer(aHandler IAccountHandler, hHandler IHoldingHandler, tHandler ITransactionHandler) *Server {
	return &Server{aHandler: aHandler, hHandler: hHandler, tHandler: tHandler}
}

func (s *Server) registerRoutes(router *mux.Router) {
	router.HandleFunc("/transaction/", s.tHandler.GetTransactionsHandler).Methods(http.MethodGet)
	router.HandleFunc("/transaction/{id}", s.tHandler.GetTransactionRecordHandler).Methods(http.MethodGet)
	router.HandleFunc("/transaction/", s.tHandler.CreateTransactionHandler).Methods(http.MethodPost)

	router.HandleFunc("/account/", s.aHandler.GetBalanceHandler).Methods(http.MethodGet)
	router.HandleFunc("/account/{quantity}", s.aHandler.UpdateBalanceHandler).Methods(http.MethodPut)

	router.HandleFunc("/holding/", s.hHandler.GetHoldingsHandler).Methods(http.MethodGet)
	router.HandleFunc("/holding/{crypto_id}", s.hHandler.GetHoldingHandler).Methods(http.MethodGet)
}

func (s *Server) Start() {
	router := mux.NewRouter()
	router.Use()
	s.registerRoutes(router)
	log.Fatal(http.ListenAndServe(":5050", router))
}
