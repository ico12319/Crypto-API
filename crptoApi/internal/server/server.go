package server

import (
	"crptoApi/internal/middlewares"
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

func (s *Server) registerTransactionRoutes(router *mux.Router) {
	router.HandleFunc("/", s.tHandler.GetTransactionsHandler).Methods(http.MethodGet)
	router.HandleFunc("/{id}", s.tHandler.GetTransactionRecordHandler).Methods(http.MethodGet)
	router.HandleFunc("/", s.tHandler.CreateTransactionHandler).Methods(http.MethodPost)
}

func (s *Server) registerAccountRoutes(router *mux.Router) {
	router.HandleFunc("/", s.aHandler.GetBalanceHandler).Methods(http.MethodGet)
	router.HandleFunc("/{quantity}", s.aHandler.UpdateBalanceHandler).Methods(http.MethodPut)
}

func (s *Server) registerHoldingRoutes(router *mux.Router) {
	router.HandleFunc("/", s.hHandler.GetHoldingsHandler).Methods(http.MethodGet)
	router.HandleFunc("/{crypto_id}", s.hHandler.GetHoldingHandler).Methods(http.MethodGet)
}

func (s *Server) Start() {
	router := mux.NewRouter()
	router.Use(middlewares.ValidationMiddlewareFunc)
	router.Use(middlewares.ContentTypeMiddlewareFunc)

	transactionRouter := router.PathPrefix("/transaction").Subrouter()
	transactionRouter.Use(middlewares.LoggingMiddlewareFunc)
	s.registerTransactionRoutes(transactionRouter)

	accountRouter := router.PathPrefix("/account").Subrouter()
	s.registerAccountRoutes(accountRouter)

	holdingRouter := router.PathPrefix("/holding").Subrouter()
	s.registerHoldingRoutes(holdingRouter)

	log.Fatal(http.ListenAndServe(":5050", router))
}
