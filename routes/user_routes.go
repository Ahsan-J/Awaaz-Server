package routes

import (
	"awaaz_go_server/helpers"
	"awaaz_go_server/controller/user"
	"github.com/gorilla/mux"
)

// GenerateUserAPIRoutes generates the required user API to serve
func generateUserAPIRoutes(router *mux.Router) {
	users := router.PathPrefix("/users").Subrouter()
	users.Use(helpers.LoggingMiddleware)
	users.HandleFunc("", user.GetAll).Methods("GET") // GET Request to handle all data present in the Database

	sub := router.PathPrefix("/user").Subrouter()
	sub.Use(helpers.LoggingMiddleware)
	
	sub.HandleFunc("", user.GetUser).Methods("GET")
}
