package routes

import (
	"be-dumbflix/handlers"
	"be-dumbflix/pkg/mysql"
	"be-dumbflix/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
  authRepository := repositories.RepositoryAuth(mysql.DB)
  h := handlers.HandlerAuth(authRepository)

  r.HandleFunc("/register", h.Register).Methods("POST")
  r.HandleFunc("/login", h.Login).Methods("POST")
}