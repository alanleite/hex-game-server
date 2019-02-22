package routes

import (
	"github.com/alanleite/hex-game-server/internal/pkg/controllers/echo"
	"github.com/gorilla/mux"
)

//Socket routers
func Socket() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/echo", echo.Index).Methods("GET")

	return router

}
