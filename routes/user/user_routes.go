package user

import (
	"../../helpers"
	"github.com/gorilla/mux"
)

// GenerateUserAPIRoutes generates the required user API to serve
func GenerateUserAPIRoutes(router *mux.Router) {
	r := router.PathPrefix("/users").Subrouter()
	r.Use(helpers.LoggingMiddleware)
	r.HandleFunc("", getAll).Methods("GET") // GET Request to handle all data present in the Database

	sub := router.PathPrefix("/user").Subrouter()
	// sub.Use(helpers.LoggingMiddleware)
	// All of the Registered User API routes
	sub.Use(helpers.LoggingMiddleware)
	sub.HandleFunc("", getUser).Methods("GET")
}
