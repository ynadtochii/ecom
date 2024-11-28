package api

import (
	"database/sql"
	"net/http"

	"github.com/ynadtochii/ecom/cmd/service/health"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}
func (s *APIServer) Run() error {
	router := http.NewServeMux()

	health := health.NewHandler()
	health.Health(router)

	return http.ListenAndServe(":8080", router)
}
