package api

import (
	"database/sql"
	"net/http"

	"github.com/ynadtochii/ecom/cmd/route/health"
	"github.com/ynadtochii/ecom/cmd/route/user"
	"github.com/ynadtochii/ecom/db"
	"github.com/ynadtochii/ecom/db/repositories"
	"github.com/ynadtochii/ecom/services"
	"gorm.io/gorm"
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

	userRepo := &repositories.UserRepository{DB: (*gorm.DB)(db.DB)}
	userService := services.NewUserService(userRepo)

	user := user.NewHandler(userService)
	user.GetUsers(router)
	user.GetUserById(router)
	user.CreateUser(router)
	user.UpdateUser(router)
	user.DeleteUser(router)

	return http.ListenAndServe(":8080", router)
}
