package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mounis-bhat/go-bank/pkg/db"
)

type APIServer struct {
	listenAddress string
	dbConnection  *db.PostgresStorage
}

func NewAPIServer(listenAddress string, dbConnection *db.PostgresStorage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		dbConnection:  dbConnection,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", s.handleGetAccount).Methods(http.MethodGet)
	router.HandleFunc("/account", s.handleCreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/login", s.handleLogin).Methods(http.MethodPost)

	fmt.Println("Server is running 🚀")
	http.ListenAndServe(s.listenAddress, router)
}
