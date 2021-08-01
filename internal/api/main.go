package api

import (
	"genesis/internal/api/handlers"
	"genesis/internal/server"
	"github.com/go-chi/chi/v5"
)

func NewRouter() chi.Router {
	r:=chi.NewRouter()
	r.Get("/btc", handlers.GetBTC)
	r.Get("/user/create", server.CreateUser)
	r.Get("/user/auth",server.AuthenticateUser)
	return r
}
